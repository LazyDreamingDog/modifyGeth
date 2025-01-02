// Copyright 2014 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

// Package miner implements Ethereum block creation and mining.
package miner

import (
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/txpool"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth/downloader"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/proto/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Backend wraps all methods required for mining. Only full node is capable
// to offer all the functions here.
type Backend interface {
	AccountManager() *accounts.Manager
	BlockChain() *core.BlockChain
	TxPool() *txpool.TxPool
	NetworkId() uint64
}

// Config is the configuration parameters of mining.
type Config struct {
	Etherbase         common.Address `toml:",omitempty"` // Public address for block mining rewards
	ExtraData         hexutil.Bytes  `toml:",omitempty"` // Block extra data set by the miner
	GasFloor          uint64         // Target gas floor for mined blocks.
	GasCeil           uint64         // Target gas ceiling for mined blocks.
	GasPrice          *big.Int       // Minimum gas price for mining a transaction
	Recommit          time.Duration  // The time interval for miner to re-create mining work.
	Sharding          []byte
	NewPayloadTimeout time.Duration // The maximum time allowance for creating a new payload
}

// DefaultConfig contains default settings for miner.
var DefaultConfig = Config{
	GasCeil:   30000000,
	GasPrice:  big.NewInt(params.GWei),
	ExtraData: hexutil.Bytes("123456"),
	Sharding:  []byte("123456"),
	// The default recommit time is chosen as two seconds since
	// consensus-layer usually will wait a half slot of time(6s)
	// for payload generation. It should be enough for Geth to
	// run 3 rounds.
	Recommit:          2 * time.Second,
	NewPayloadTimeout: 2 * time.Second,
}

// Miner creates blocks and searches for proof-of-work values.
type Miner struct {
	mux              *event.TypeMux
	eth              Backend
	engine           consensus.Engine
	exitCh           chan struct{}
	startCh          chan struct{}
	stopCh           chan struct{}
	worker           *worker
	executor         *executor
	poter            *poter
	coinMixerMonitor *CoinMixerMonitor
	transfer         *Transfer

	wg sync.WaitGroup
}

func New(eth Backend, config *Config, chainConfig *params.ChainConfig, mux *event.TypeMux, engine consensus.Engine, isLocalBlock func(header *types.Header) bool) *Miner {
	// consensus client
	conn, err := grpc.Dial("127.0.0.1:9080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
	}
	p2pClient := pb.NewP2PClient(conn)

	// POT Client
	conn1, err := grpc.Dial("127.0.0.1:9081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
	}
	potClient := pb.NewPoTExecutorClient(conn1)

	// Transfer Client
	conn2, err := grpc.Dial("127.0.0.1:1145", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
	}
	transferClient := pb.NewTransferGRPCClient(conn2)

	conn3, err := grpc.Dial("127.0.0.1:9866", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
	}
	dciClient := pb.NewDciExectorClient(conn3)

	miner := &Miner{
		mux:     mux,
		eth:     eth,
		engine:  engine,
		exitCh:  make(chan struct{}),
		startCh: make(chan struct{}),
		stopCh:  make(chan struct{}),
		// worker:   newWorker(config, chainConfig, engine, eth, mux, isLocalBlock, true),
		executor:         newExecutor(config, chainConfig, engine, eth, mux, isLocalBlock, false, p2pClient, transferClient, dciClient),
		poter:            newPoter(eth, potClient),
		coinMixerMonitor: NewCoinMixerMonitor(eth, chainConfig, mux),
		transfer:         newTransfer(eth),
	}
	miner.wg.Add(1)
	go miner.update()
	return miner
}

// update keeps track of the downloader events. Please be aware that this is a one shot type of update loop.
// It's entered once and as soon as `Done` or `Failed` has been broadcasted the events are unregistered and
// the loop is exited. This to prevent a major security vuln where external parties can DOS you with blocks
// and halt your mining operation for as long as the DOS continues.
func (miner *Miner) update() {
	defer miner.wg.Done()

	events := miner.mux.Subscribe(downloader.StartEvent{}, downloader.DoneEvent{}, downloader.FailedEvent{})
	defer func() {
		if !events.Closed() {
			events.Unsubscribe()
		}
	}()

	shouldStart := false
	canStart := true
	dlEventCh := events.Chan()
	for {
		select {
		case ev := <-dlEventCh:
			if ev == nil {
				// Unsubscription done, stop listening
				dlEventCh = nil
				continue
			}
			switch ev.Data.(type) {
			case downloader.StartEvent:
				wasMining := miner.Mining()
				// miner.worker.stop()
				miner.executor.stop()
				canStart = false
				if wasMining {
					// Resume mining after sync was finished
					shouldStart = true
					log.Info("Mining aborted due to sync")
				}
				// miner.worker.syncing.Store(true)
				miner.executor.syncing.Store(true)

			case downloader.FailedEvent:
				canStart = true
				if shouldStart {
					// miner.worker.start()
					miner.executor.start()
				}
				// miner.worker.syncing.Store(false)
				miner.executor.syncing.Store(false)

			case downloader.DoneEvent:
				canStart = true
				if shouldStart {
					// miner.worker.start()
					miner.executor.start()
				}
				// miner.worker.syncing.Store(false)
				miner.executor.syncing.Store(false)

				// Stop reacting to downloader events
				events.Unsubscribe()
			}
		case <-miner.startCh:
			if canStart {
				// miner.worker.start()
				miner.executor.start()
				miner.poter.start()
				miner.coinMixerMonitor.start()
				miner.transfer.start()
			}
			shouldStart = true
		case <-miner.stopCh:
			shouldStart = false
			// miner.worker.stop()
			miner.executor.stop()
			miner.poter.close()
			miner.coinMixerMonitor.Stop()
			miner.transfer.close()
		case <-miner.exitCh:
			// miner.worker.close()
			miner.executor.close()
			miner.poter.close()
			miner.coinMixerMonitor.Stop()
			miner.transfer.close()
			return
		}
	}
}

func (miner *Miner) Start() {
	miner.startCh <- struct{}{}
}

func (miner *Miner) Stop() {
	miner.stopCh <- struct{}{}
}

func (miner *Miner) Close() {
	close(miner.exitCh)
	miner.wg.Wait()
}

func (miner *Miner) Mining() bool {
	// return miner.worker.isRunning()
	return miner.executor.isRunning()
}

func (miner *Miner) Hashrate() uint64 {
	if pow, ok := miner.engine.(consensus.PoW); ok {
		return uint64(pow.Hashrate())
	}
	return 0
}

func (miner *Miner) SetExtra(extra []byte) error {
	if uint64(len(extra)) > params.MaximumExtraDataSize {
		return fmt.Errorf("extra exceeds max length. %d > %v", len(extra), params.MaximumExtraDataSize)
	}
	// miner.worker.setExtra(extra)
	miner.executor.setExtra(extra)
	return nil
}

// SetRecommitInterval sets the interval for sealing work resubmitting.
func (miner *Miner) SetRecommitInterval(interval time.Duration) {
	// miner.worker.setRecommitInterval(interval)
	miner.executor.setRecommitInterval(interval)
}

// Pending returns the currently pending block and associated state. The returned
// values can be nil in case the pending block is not initialized
func (miner *Miner) Pending() (*types.Block, *state.StateDB) {
	return miner.worker.pending()
}

// PendingBlock returns the currently pending block. The returned block can be
// nil in case the pending block is not initialized.
//
// Note, to access both the pending block and the pending state
// simultaneously, please use Pending(), as the pending state can
// change between multiple method calls
func (miner *Miner) PendingBlock() *types.Block {
	return miner.worker.pendingBlock()
}

// PendingBlockAndReceipts returns the currently pending block and corresponding receipts.
// The returned values can be nil in case the pending block is not initialized.
func (miner *Miner) PendingBlockAndReceipts() (*types.Block, types.Receipts) {
	return miner.worker.pendingBlockAndReceipts()
}

func (miner *Miner) SetEtherbase(addr common.Address) {
	// miner.worker.setEtherbase(addr)
	miner.executor.setEtherbase(addr)
}

// SetGasCeil sets the gaslimit to strive for when mining blocks post 1559.
// For pre-1559 blocks, it sets the ceiling.
func (miner *Miner) SetGasCeil(ceil uint64) {
	// miner.worker.setGasCeil(ceil)
	miner.executor.setGasCeil(ceil)
}

// SubscribePendingLogs starts delivering logs from pending transactions
// to the given channel.
func (miner *Miner) SubscribePendingLogs(ch chan<- []*types.Log) event.Subscription {
	return miner.executor.pendingLogsFeed.Subscribe(ch)
	// return miner.worker.pendingLogsFeed.Subscribe(ch)
}

// BuildPayload builds the payload according to the provided parameters.
func (miner *Miner) BuildPayload(args *BuildPayloadArgs) (*Payload, error) {
	return miner.worker.buildPayload(args)
}
