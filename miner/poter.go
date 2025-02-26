package miner

import (
	"context"
	"fmt"
	"net"
	"sync/atomic"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/txpool"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/proto/pb"
	"google.golang.org/grpc"
)

type poter struct {
	eth    Backend // blockchain and txpool
	chain  *core.BlockChain
	txpool *txpool.TxPool

	// running atomic.Bool // a functional judge
	serving   atomic.Bool
	networkId uint64

	poterClient pb.PoTExecutorClient

	server *grpc.Server // server pointer to the running server
	pb.UnimplementedPoTExecutorServer
}

func newPoter(eth Backend, cli pb.PoTExecutorClient) *poter {
	poter := &poter{
		eth:    eth,
		chain:  eth.BlockChain(),
		txpool: eth.TxPool(),
		// running: atomic.Bool{},
		serving:   atomic.Bool{},
		networkId: eth.NetworkId(),

		server: grpc.NewServer(),
	}
	poter.poterClient = cli
	// Register the grpc server
	s := grpc.NewServer()
	pb.RegisterPoTExecutorServer(s, poter)
	poter.server = s // then we can handle the server

	poter.serving.Store(false)

	return poter
}

// start sets the running status as 1 and triggers new work submitting.
func (p *poter) start() {

	if !p.serving.Load() {
		// !!! 这一段应该进入配置文件
		listen, err := net.Listen("tcp", "127.0.0.1:9877") // will be included in config
		if err != nil {
			fmt.Println(err)
			panic("cannot listen!")
		}
		p.serving.Store(true)
		go p.server.Serve(listen)
	}

}

// func (p *poter) isRunning() bool {
// 	return p.running.Load()
// }

func (p *poter) close() {
	p.server.Stop()
	p.serving.Store(false)
}

func (p *poter) GetTxs(ctx context.Context, getTxRq *pb.GetTxRequest) (*pb.GetTxResponse, error) {
	// get the transaction from the txpool
	res := &pb.GetTxResponse{
		Start:  getTxRq.GetStartHeight(),
		End:    p.eth.BlockChain().CurrentHeader().Number.Uint64(),
		Blocks: make([]*pb.ExecuteBlock, 0),
	}

	var coinbaseValue int64

	for i := getTxRq.GetStartHeight(); i <= p.eth.BlockChain().CurrentHeader().Number.Uint64(); i++ {
		block := p.eth.BlockChain().GetBlockByNumber(i)
		header := &pb.ExecuteHeader{
			Height:        block.Header().Number.Uint64(),
			BlockHash:     block.Header().Hash().Bytes(),
			ChainID:       int64(p.networkId),
			TxsHash:       block.Header().Root[:],
			CommitedTxNum: block.Header().CommitTxLength,
			ExecutedTxNum: uint64(len(block.Transactions())),
			GasIncentive:  block.Header().Incentive.ToBig().Uint64(),
			PoSLeader:     block.Header().PoSLeader[:],
			PoSVoteInfo:   block.Header().PoSVoting,
		}

		coinbaseValue += block.Header().Incentive.ToBig().Int64()

		txs := make([]*pb.ExecutedTx, 0)

		for _, tx := range block.Transactions() {
			txs = append(txs, &pb.ExecutedTx{
				TxHash: tx.Hash().Bytes(),
				Height: block.Header().Number.Uint64(),
				Data:   tx.Data(),
			})
		}

		res.Blocks = append(res.Blocks, &pb.ExecuteBlock{
			Header: header,
			Txs:    txs,
		})
	}

	return res, nil
}

func (p *poter) VerifyTxs(ctx context.Context, veriRq *pb.VerifyTxRequest) (*pb.VerifyTxResponse, error) {
	res := &pb.VerifyTxResponse{
		Txs:  veriRq.GetTxs(),
		Flag: make([]bool, len(veriRq.GetTxs())),
	}

	for _, txData := range veriRq.GetTxs() {
		// verify the transaction
		lookup, _, err := p.eth.BlockChain().GetTransactionLookup(common.Hash(txData.TxHash))
		if err != nil {
			return nil, err
		}
		if lookup == nil {
			res.Flag = append(res.Flag, false)
			continue
		}
		if lookup.BlockIndex == txData.ExecutedHeight {
			res.Flag = append(res.Flag, true)
		}
	}

	return res, nil

}

// IncensentiveVerify is a function to verify the incensentive of Each partition
func (p *poter) IncensentiveVerify(ctx context.Context, veriRq *pb.IncensentiveVerifyRequest) (*pb.IncensentiveVerifyResponse, error) {
	// 传入block Hash 和 txindex ，返回验证结果
	res := make([]bool, len(veriRq.TxHash))
	for index, txHash := range veriRq.TxHash {
		receipts := p.chain.GetReceiptsByHash(common.Hash(txHash))
		// TODO : Supplement the event verification logic of each district
		if receipts[0].Status == 1 {
			res[index] = true
		} else {
			res[index] = false
		}
	}
	return &pb.IncensentiveVerifyResponse{VerifyRes: res}, nil
}

func (p *poter) ExecuteTxs(ctx context.Context, exeRq *pb.ExecuteTxRequest) (*pb.ExecuteTxResponse, error) {
	txByte := exeRq.Tx
	tx := new(types.Transaction)
	// unmarshal the tx
	err := tx.UnmarshalBinary(txByte)
	if err != nil {
		return &pb.ExecuteTxResponse{Tx: exeRq.Tx, Flag: false, TxID: nil}, err
	}
	// add the tx to the txpool
	txs := []*types.Transaction{tx}
	errs := p.txpool.Add(txs, true, false)
	if errs[0] != nil {
		return &pb.ExecuteTxResponse{Tx: exeRq.Tx, Flag: false, TxID: tx.Hash().Bytes()}, errs[0]
	}

	return &pb.ExecuteTxResponse{Tx: exeRq.Tx, Flag: true, TxID: tx.Hash().Bytes()}, nil
}

func (p *poter) GetIncentive(ctx context.Context, getIncentiveRq *pb.GetIncentiveRequest) (*pb.GetIncentiveResponse, error) {
	incentiveArray, newEnd, err := p.chain.GetCoinBaseIncentive(getIncentiveRq.Begin, getIncentiveRq.End)
	if err != nil {
		return &pb.GetIncentiveResponse{BciReward: nil}, err
	}
	return &pb.GetIncentiveResponse{
		BciReward: incentiveArray,
		End:       newEnd,
	}, nil
}
