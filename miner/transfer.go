package miner

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"net"
	"sync"
	"sync/atomic"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/txpool"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/proto/pb"
	"google.golang.org/grpc"
)

type txReq struct {
	verifyHash []byte
	height     uint64
	tx         *types.Transaction
}

type Transfer struct {
	eth    Backend // blockchain and txpool
	chain  *core.BlockChain
	txpool *txpool.TxPool

	// running atomic.Bool // a functional judge
	serving atomic.Bool
	wg      sync.WaitGroup // for go-routine

	txChannel chan *txReq
	exitCh    chan struct{}

	server                                     *grpc.Server // server pointer to the running server
	pb.UnimplementedTransferExecutorGRPCServer              // indicated transfer can be a grpc server
}

func newTransfer(eth Backend) *Transfer {
	transfer := &Transfer{
		eth:       eth,
		chain:     eth.BlockChain(),
		txpool:    eth.TxPool(),
		txChannel: make(chan *txReq),
		exitCh:    make(chan struct{}),
	}

	// 注册server
	s := grpc.NewServer()
	pb.RegisterTransferExecutorGRPCServer(s, transfer)
	transfer.server = s // then we can handle the server

	transfer.serving.Store(false)
	transfer.wg.Add(1)
	go transfer.loop()

	return transfer
}

func (t *Transfer) start() {
	if !t.serving.Load() {
		listen, err := net.Listen("tcp", "127.0.0.1:9809") // will be included in config
		if err != nil {
			fmt.Println(err)
			panic("cannot listen!")
		}
		t.serving.Store(true)
		go t.server.Serve(listen)
	}
}

func (t *Transfer) close() {
	t.server.Stop()
	t.serving.Store(false)
	close(t.exitCh)
	t.wg.Wait()
}

// 循环接受转账过来的交易，一旦有就add到Txpool里
func (t *Transfer) loop() {
	defer t.wg.Done()
	for {
		select {
		case txreq := <-t.txChannel:
			log.Info("get tx", "tx", txreq.tx.Hash())
			// todo: 需要验证txreq.verifyHash
			// 调用转账接口进行验证
			t.txpool.Add([]*types.Transaction{txreq.tx}, true, false)
		case <-t.exitCh:
			return
		}
	}
}

// Receive txs from utxo layer
func (t *Transfer) CommitWithdrawTx(ctx context.Context, request *pb.CommitWithdrawTxRequest) (*pb.Empty, error) {
	// // 方案一：构造一个from是空的系统交易
	// toAddress := common.BytesToAddress(request.To)
	// value := new(big.Int).SetUint64(request.Value)
	// // 构造交易数据0x0D06
	// data := []byte{0x0D, 0x06}
	// // data后面附带验证的哈希
	// verifyHash := request.VerifyHash
	// data = append(data, verifyHash...)
	// systemTx := types.NewSystemTx(t.chain.Config().ChainID, 0, big.NewInt(0), big.NewInt(0), 0, &toAddress, value, data, 0)

	// 方案二：反序列化交易
	gettx := new(types.Transaction)
	err := gettx.UnmarshalBinary(request.TxData)
	if err != nil {
		return nil, err
	}
	// 将交易插入到txpool
	t.txChannel <- &txReq{
		height:     request.Height,
		verifyHash: request.VerifyHash,
		tx:         gettx,
	}

	// 返回空
	return &pb.Empty{}, nil
}

// 验证交易
func (t *Transfer) VerifyChangeTx(ctx context.Context, request *pb.VerifyChangeTxRequest) (*pb.VerifyChangeTxReply, error) {
	// 验证交易
	// 读取txhash和value
	txhash := request.VerifyHash
	value := request.Value
	// TODO: 需要验证重放攻击
	height := request.Height
	fmt.Println("height", height)
	// 读取账户的提现的最新高度，判断验证交易的区块高度是否落后
	// 从blockchain中读取交易进行比对
	blockchain := t.chain
	_, tx, err := blockchain.GetTransactionLookup(common.Hash(txhash))
	if err != nil {
		return &pb.VerifyChangeTxReply{VerifyRes: false}, errors.New("tx not found")
	}
	if tx.Value().Cmp(big.NewInt(int64(value))) != 0 {
		return &pb.VerifyChangeTxReply{VerifyRes: false}, errors.New("value not match")
	}
	// 返回验证结果
	return &pb.VerifyChangeTxReply{VerifyRes: true}, nil
}
