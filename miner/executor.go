package miner

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"net"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/consensus/misc/eip1559"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/txpool"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/proto/pb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

const txMaxSize = 4 * 32 * 1024 // 128KB
const ContributionContractAddr = "0x67fbF000Fc60CBE25D9658D83C5C2506ca323Fdd"

// environment is the worker's current environment and holds all
// information of the sealing block generation.
type executor_env struct {
	// 打包区块前的一些参数
	signer   types.Signer
	state    *state.StateDB // apply state changes here
	gasPool  *core.GasPool  // available gas used to pack transactions
	coinbase common.Address
	header   *types.Header

	// 最后执行的结束后的结果，有多少tx被包括，他们的收据是什么
	// 打包区块使用
	initTxcount int // 初始收到块时的交易数量
	tcount      int // 成功执行的交易数量
	txs         types.Transactions
	receipts    []*types.Receipt
}

// copy creates a deep copy of environment.
func (env *executor_env) copy() *executor_env {
	cpy := &executor_env{
		signer:   env.signer,
		state:    env.state.Copy(),
		tcount:   env.tcount,
		coinbase: env.coinbase,
		header:   types.CopyHeader(env.header),
		receipts: copyReceipts(env.receipts),
	}
	if env.gasPool != nil {
		gasPool := *env.gasPool
		cpy.gasPool = &gasPool
	}
	cpy.txs = make([]*types.Transaction, len(env.txs))
	copy(cpy.txs, env.txs)
	return cpy
}

type execReq struct {
	timestamp    int64
	txs          types.Transactions
	leader       common.Address
	randomNumber *big.Int
	incentive    []byte
}

type executorServer struct {
	executorPtr                    *executor
	pb.UnimplementedExecutorServer // indicated executor can be a grpc server
}

// func (es *executorServer) isRunning() bool {
// 	return es.running.Load()
// }

// Receive txs from consensus layer
func (es *executorServer) CommitBlock(ctx context.Context, pbBlock *pb.ExecBlock) (*pb.Empty, error) {
	// set leader
	randomNuber := new(big.Int).SetUint64(pbBlock.GetRandomNumber())

	incentiveBytes := pbBlock.GetIncentive()
	IncentiveData, err := DecodeIncentive(incentiveBytes)
	if err != nil {
		log.Warn("decode incentive failed", "err", err)
		return &pb.Empty{}, nil
	}
	// es.executorPtr.currenLeader = common.HexToAddress(IncentiveData.Leader)
	// log.Info("get leader", "leader", es.executorPtr.currenLeader)

	// // loop to get incentive.NodeIncentives
	// log.Info("get node incentives", "number", len(IncentiveData.NodeIncentives))
	// for _, nodeIncentive := range IncentiveData.NodeIncentives {
	// 	log.Info("get node incentive", "address", nodeIncentive.NodePublicAddress, "type", nodeIncentive.Type, "number", nodeIncentive.Number)
	// }

	// sharding check
	sharding, err := hexutil.DecodeUint64(string(pbBlock.ShardingName))
	if err != nil {
		log.Warn("get sharding failed", "sharding", pbBlock.ShardingName)
		return &pb.Empty{}, nil
	}

	if sharding != es.executorPtr.networkId {
		// if sharding check failed, return an error
		log.Warn("get another sharding transactions", "sharding", sharding, "networkId", es.executorPtr.networkId)
		return &pb.Empty{}, nil
	}

	pbtxs := pbBlock.GetTxs()
	if len(pbtxs) == 0 {
		return &pb.Empty{}, nil
	}
	var errs []error = make([]error, 0)
	var txs types.Transactions = make(types.Transactions, 0)

	for _, byte := range pbtxs {
		pbTx := new(pb.Transaction)
		err := proto.Unmarshal(byte, pbTx)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		tx := new(types.Transaction)
		err = tx.UnmarshalBinary(pbTx.Payload)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		txs = append(txs, tx)
	}

	log.Info("get commited tx from consensus", "txs len:", txs.Len())

	// Receive txs from consensus layer
	if txs.Len() != 0 {
		// construct a execReq
		execreq := &execReq{
			timestamp:    time.Now().Unix(),
			txs:          txs,
			randomNumber: randomNuber,
			leader:       common.HexToAddress(IncentiveData.Leader),
			incentive:    incentiveBytes,
		}
		es.executorPtr.execCh <- execreq
	}

	// Check if there are protobuf errors in the consensus block
	if len(errs) != 0 {
		errStr := fmt.Sprintf("There are %d errors in the block", len(errs))
		return &pb.Empty{}, fmt.Errorf(errStr)
	}
	return &pb.Empty{}, nil
}

func (es *executorServer) VerifyTx(ctx context.Context, pTx *pb.Transaction) (*pb.Result, error) {
	if pTx.Type != pb.TransactionType_NORMAL && pTx.Type != pb.TransactionType_UPGRADE {
		return &pb.Result{Success: false}, nil
	}
	tx := new(types.Transaction)
	err := tx.UnmarshalBinary(pTx.Payload)
	if err != nil {
		return &pb.Result{Success: false}, nil
	}
	// default all txs here are remote
	env := es.executorPtr.env
	err = txpool.ValidateTransaction(tx, env.header, env.signer, es.executorPtr.opts)
	if err != nil {
		return &pb.Result{Success: false}, nil
	}
	return &pb.Result{Success: true}, nil
}

//----------------------------------------------------------------------------------------------

type executorClient struct {
	consensusClient pb.P2PClient // to send txs to consensus layer

	transferClient pb.TransferGRPCClient

	dciClient pb.DciExectorClient
}

// need add a loop routine to sendTx to consensus layer, when execCh has new txs
func (ec *executorClient) sendTx(tx *types.Transaction, nid uint64) (*pb.Empty, error) {
	log.Info("begin send tx to consensus")
	data, err := tx.MarshalBinary()
	if err != nil {
		return nil, err
	}
	ptx := &pb.Transaction{
		Type:    pb.TransactionType_NORMAL,
		Payload: data,
	}
	btx, err := proto.Marshal(ptx)
	if err != nil {
		return nil, err
	}

	// add Indentifer
	// request := &pb.Request{Tx: btx}

	// TODO：在发送交易时加上执行节点的分区标识networkid
	sharding := []byte(hexutil.EncodeUint64(nid))
	log.Info("send tx to consensus", "sharding", sharding)
	request := &pb.Request{
		Tx:       btx,
		Sharding: sharding,
	}

	rawRequest, err := proto.Marshal(request)
	if err != nil {
		return nil, err
	}
	packet := &pb.Packet{
		Msg:         rawRequest,
		ConsensusID: -1,
		Epoch:       -1,
		Type:        pb.PacketType_CLIENTPACKET,
	}
	_, err = ec.consensusClient.Send(context.Background(), packet)
	if err != nil {
		return nil, err
	}
	log.Info("finish send tx to consensus")
	return &pb.Empty{}, nil

}

func (ec *executorClient) verifyTokenTransitionTx(tx *types.Transaction) (bool, error) {
	request := &pb.VerifyUTXORequest{
		From:  nil,
		To:    tx.To().Bytes(),
		Value: tx.Value().Int64(),
		Proof: tx.Data(),
	}
	res, err := ec.dciClient.VerifyUTXO(context.Background(), request)
	if err != nil {
		return false, err
	}
	return res.Flag, nil
}

//----------------------------------------------------------------------------------------------

type executor struct {
	config      *Config                   // other config
	chainConfig *params.ChainConfig       // chain config
	engine      consensus.Engine          // assemble block
	eth         Backend                   // blockchain and txpool
	opts        *txpool.ValidationOptions // to do basic validation of Tx
	chain       *core.BlockChain

	pendingLogsFeed event.Feed

	env *executor_env  // may be useful in some local situations
	wg  sync.WaitGroup // for go-routine

	running atomic.Bool // a functional judge
	syncing atomic.Bool // The indicator whether the node is still syncing.
	serving atomic.Bool

	startCh chan struct{} // ...
	exitCh  chan struct{} // ...

	resubmitIntervalCh chan time.Duration

	// Subscriptions
	mux *event.TypeMux

	newWorkCh  chan *newWorkReq // to launch a new batch to consensus
	execCh     chan *execReq    // received from consensus, and go to execute
	offChainCh chan bool        //  communicate with WASM

	mu           sync.RWMutex   // The lock used to protect the coinbase
	coinbase     common.Address // yeah, baby
	currenLeader common.Address // consensus leader
	extra        []byte
	networkId    uint64

	// recommit is the time interval to re-create sealing work or to re-build
	// payload in proof-of-stake stage.
	recommit time.Duration

	// client to consensus layer
	execClient *executorClient

	// server to consensus layer
	server *grpc.Server // server pointer to the running server

	// difficultyAdaptor *DifficultyAdaptor
	// priceAdaptor      *PriceAdaptor
	powAdaptor *PoWAdaptor
	gasAdaptor *GasAdaptor

	planPool *core.PlanPool
}

// newExecutor creates a new executor.
func newExecutor(config *Config, chainConfig *params.ChainConfig, engine consensus.Engine, eth Backend, mux *event.TypeMux, isLocalBlock func(header *types.Header) bool, init bool, consensusCli pb.P2PClient, transferCli pb.TransferGRPCClient, dciClient pb.DciExectorClient) *executor {
	executor := &executor{
		config:      config,
		chainConfig: chainConfig,
		engine:      engine,
		eth:         eth,
		chain:       eth.BlockChain(),
		mux:         mux,

		opts: &txpool.ValidationOptions{
			Config: chainConfig,
			Accept: 0 |
				1<<types.LegacyTxType |
				1<<types.AccessListTxType |
				1<<types.DynamicFeeTxType,
			MaxSize: txMaxSize,
			MinTip:  config.GasPrice,
		},

		coinbase:  config.Etherbase,
		extra:     config.ExtraData,
		networkId: eth.NetworkId(),
		// pendingTasks: make(map[common.Hash]*task),

		// chainHeadCh: make(chan core.ChainHeadEvent, chainHeadChanSize),

		startCh: make(chan struct{}, 1),
		exitCh:  make(chan struct{}),

		resubmitIntervalCh: make(chan time.Duration),

		newWorkCh:  make(chan *newWorkReq),
		execCh:     make(chan *execReq),
		offChainCh: make(chan bool),

		// TODO: 有问题 Params需要被plan改动，不然重启之后就没有用了
		gasAdaptor: NewGasAdaptor(
			params.MinPowGas,  // minGas: 最小 gas 限制
			params.MaxPowGas,  // maxGas: 最大 gas 限制，与 DefaultConfig.GasCeil 保持一致
			params.InitialGas, // initialGas: 初始 gas 限制
			params.Alpha,      // alpha: EMA 平滑因子为 0.2
		),
		powAdaptor: NewPoWAdaptor(
			params.TargetPowRatio,    // targetPowRatio: 目标 PoW 交易比例为 30%
			params.Alpha,             // alpha: EMA 平滑因子为 0.2
			params.Fmin,              // fMin: 最小调整因子为 0.8
			params.Fmax,              // fMax: 最大调整因子为 1.2
			params.InitialDifficulty, // initialDifficulty: 初始难度
			params.Kp,                // kp: 比例调节系数为 0.1
			params.Ki,                // ki: 积分调节系数为 0.01
			params.MinPrice,          // minPrice: 最小价格
			params.MaxPrice,          // maxPrice: 最大价格
		),
		planPool: core.NewPlanPool(),
	}
	// TODO: 草率开始聆听
	StartPlanClient(executor.planPool, 22324)
	// Subscribe events for blockchain
	// executor.chainHeadSub = eth.BlockChain().SubscribeChainHeadEvent(executor.chainHeadCh)

	// Sanitize recommit interval if the user-specified one is too short.
	// recommit := executor.config.Recommit
	// if recommit < minRecommitInterval {
	// 	log.Warn("Sanitizing miner recommit interval", "provided", recommit, "updated", minRecommitInterval)
	// 	recommit = minRecommitInterval
	// }
	// TODO : 暂定recommit为minRecommitInterval
	recommit := minRecommitInterval
	executor.recommit = recommit

	// Register the grpc client
	executor.execClient = &executorClient{
		consensusClient: consensusCli,
		transferClient:  transferCli,
		dciClient:       dciClient,
	}

	// Register the grpc server
	executorServer := executorServer{executorPtr: executor}
	s := grpc.NewServer()
	pb.RegisterExecutorServer(s, &executorServer)
	executor.server = s // then we can handle the server

	executor.serving.Store(false)

	// start loop
	executor.wg.Add(3)
	go executor.sendLoop()
	go executor.executionLoop()
	go executor.newExecLoop(recommit)

	// Submit first work to initialize pending state.
	if init {
		executor.startCh <- struct{}{}
	}
	return executor
}

// isRunning returns an indicator whether worker is running or not.
func (e *executor) isRunning() bool {
	return e.running.Load()
}

// start sets the running status as 1 and triggers new work submitting.
func (e *executor) start() {
	e.running.Store(true)
	if !e.serving.Load() {
		// !!! 这一段应该进入配置文件
		listen, err := net.Listen("tcp", "127.0.0.1:9876") // will be included in config
		if err != nil {
			fmt.Println(err)
			panic("cannot listen!")
		}
		e.serving.Store(true)
		go e.server.Serve(listen)
	}

	e.startCh <- struct{}{}
}

// stop sets the running status as 0.
func (e *executor) stop() {
	e.running.Store(false)
}

// close terminates all background threads maintained by the worker.
// Note the worker does not support being closed multiple times.
func (e *executor) close() {
	e.running.Store(false)
	e.server.Stop()
	close(e.exitCh)
	e.wg.Wait()
}

// setEtherbase sets the etherbase used to initialize the block coinbase field.
func (e *executor) setEtherbase(addr common.Address) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.coinbase = addr
}

// etherbase retrieves the configured etherbase address.
func (e *executor) etherbase() common.Address {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return e.coinbase
}

// setExtra sets the content used to initialize the block extra field.
func (e *executor) setExtra(extra []byte) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.extra = extra
	// transfer extra to networkId
	// nid, err := uint256.FromHex(common.Bytes2Hex(extra))
	// if err != nil {
	// 	e.networkId = nid.Uint64()
	// }
}

func (e *executor) setGasCeil(ceil uint64) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.config.GasCeil = ceil
}

// setRecommitInterval updates the interval for miner sealing work recommitting.
func (e *executor) setRecommitInterval(interval time.Duration) {
	select {
	case e.resubmitIntervalCh <- interval:
	case <-e.exitCh:
	}
}

// // updateSnapshot updates pending snapshot block, receipts and state.
// func (e *executor) updateSnapshot(env *environment) {
// 	e.snapshotMu.Lock()
// 	defer e.snapshotMu.Unlock()

// 	e.snapshotBlock = types.NewBlock(
// 		env.header,
// 		env.txs,
// 		nil,
// 		env.receipts,
// 		trie.NewStackTrie(nil),
// 	)
// 	e.snapshotReceipts = copyReceipts(env.receipts)
// 	e.snapshotState = env.state.Copy()
// }

// 缺少启动用循环newWorkLoop
// newExecLoop
func (e *executor) newExecLoop(recommit time.Duration) {
	defer e.wg.Done()
	var (
		interrupt *atomic.Int32
		// minRecommit = recommit // minimal resubmit interval specified by user.
		timestamp int64 // timestamp for each round of sealing.
	)

	timer := time.NewTimer(0)
	defer timer.Stop()
	<-timer.C // discard the initial tick

	// commit aborts in-flight transaction execution with given signal and resubmits a new one.
	commit := func(s int32) {
		// 这里应该不用暂存中断的逻辑？
		// if interrupt != nil {
		// 	interrupt.Store(s)
		// }
		interrupt = new(atomic.Int32)
		select {
		case e.newWorkCh <- &newWorkReq{interrupt: interrupt, timestamp: timestamp}:
		case <-e.exitCh:
			return
		}
		timer.Reset(recommit)
	}

	// clearPending cleans the stale pending tasks.
	// clearPending := func(number uint64) {
	// 	e.pendingMu.Lock()
	// 	for h, t := range e.pendingTasks {
	// 		if t.block.NumberU64()+staleThreshold <= number {
	// 			delete(e.pendingTasks, h)
	// 		}
	// 	}
	// 	e.pendingMu.Unlock()
	// }

	// 逻辑大概是启动的时候发一个信号开启sendloop，然后每隔recommit时间发一个信号启动sendloop（1秒一次）
	// 比较担心的是这些interrupt的处理，不知道是不是有
	for {
		select {
		case <-e.startCh:
			// clearPending(e.chain.CurrentBlock().Number.Uint64())
			// fmt.Println("send the first start signal")
			timestamp = time.Now().Unix()
			commit(commitInterruptNewHead)

		case <-timer.C:
			// fmt.Println("send the time start signal")
			if e.isRunning() {
				commit(commitInterruptResubmit)
			}

		// case head := <-e.chainHeadCh:
		// 	// clearPending(head.Block.NumberU64())
		// 	// no use
		// 	head.Block.Number()
		// 	timestamp = time.Now().Unix()
		// 	// make a block and send a block! 666
		// 	commit(commitInterruptNewHead)

		case <-e.exitCh:
			return
		}
	}
}

// prepareWork constructs the sealing task according to the given parameters,
// either based on the last chain head or specified parent. In this function
// the pending transactions are not filled yet, only the empty task returned.
func (e *executor) prepareWork(genParams *generateParams) (*executor_env, error) {
	e.mu.RLock()
	defer e.mu.RUnlock()

	// Find the parent block for sealing task
	parent := e.eth.BlockChain().CurrentBlock()
	if genParams.parentHash != (common.Hash{}) {
		block := e.eth.BlockChain().GetBlockByHash(genParams.parentHash)
		if block == nil {
			return nil, fmt.Errorf("missing parent")
		}
		parent = block.Header()
	}
	// Sanity check the timestamp correctness, recap the timestamp
	// to parent+1 if the mutation is allowed.
	timestamp := genParams.timestamp
	if parent.Time >= timestamp {
		if genParams.forceTime {
			return nil, fmt.Errorf("invalid timestamp, parent %d given %d", parent.Time, timestamp)
		}
		timestamp = parent.Time + 1
	}

	// TODO: Havn't test yet.
	nextPlanHeight := e.planPool.GetMinHeight()
	if parent.Number.Uint64()+1 == nextPlanHeight {
		e.ExecuteAndRemovePlan(nextPlanHeight)
	}

	newPoWPrice := big.NewInt(0)
	newGas := uint64(0)
	newGasNumerator := uint64(0)
	newGasDenominator := uint64(0)
	newRatioNumerator := uint64(0)
	newRatioDenominator := uint64(0)

	if genParams.isExecution {
		// Adjust PoW parameters
		_, newPoWPrice, newRatioNumerator, newRatioDenominator = e.powAdaptor.AdjustParameters(
			genParams.currentRatio,
			common.NewRational(parent.AvgRatioNumerator, parent.AvgRatioDenominator),
			parent.PowPrice,
		)

		// Adjust gas parameters
		newGas, newGasNumerator, newGasDenominator = e.gasAdaptor.AdjustGas(
			genParams.currentGasRatio,
			common.NewRational(parent.AvgGasNumerator, parent.AvgGasDenominator),
		)
	}

	header := &types.Header{
		ParentHash:          parent.Hash(),
		Number:              new(big.Int).Add(parent.Number, common.Big1),
		GasLimit:            core.CalcGasLimit(parent.GasLimit, e.config.GasCeil),
		Time:                timestamp,
		Coinbase:            genParams.coinbase,
		Difficulty:          big.NewInt(1),
		PowPrice:            newPoWPrice,
		PowGas:              newGas,
		AvgGasNumerator:     newGasNumerator,
		AvgGasDenominator:   newGasDenominator,
		AvgRatioNumerator:   newRatioNumerator,
		AvgRatioDenominator: newRatioDenominator,
		// random
		RandomNumber: genParams.currentRandomNumber,
		// consensus info
		PoSLeader: genParams.currentLeader,
		PoSVoting: genParams.votingData,
	}

	if len(e.extra) != 0 {
		header.Extra = e.extra
	}

	if e.chainConfig.IsLondon(header.Number) {
		header.BaseFee = eip1559.CalcBaseFee(e.chainConfig, parent)
		if !e.chainConfig.IsLondon(parent.Number) {
			parentGasLimit := parent.GasLimit * e.chainConfig.ElasticityMultiplier()
			header.GasLimit = core.CalcGasLimit(parentGasLimit, e.config.GasCeil)
		}
	}

	env, err := e.makeEnv(parent, header, genParams.coinbase)
	if err != nil {
		log.Error("Failed to create sealing context", "err", err)
		return nil, err
	}

	return env, nil
}

// makeEnv creates a new environment for the sealing block.
func (e *executor) makeEnv(parent *types.Header, header *types.Header, coinbase common.Address) (*executor_env, error) {
	// Retrieve the parent state to execute on top and start a prefetcher for
	// the miner to speed block sealing up a bit.
	state, err := e.eth.BlockChain().StateAt(parent.Root)
	if err != nil {
		return nil, err
	}
	state.StartPrefetcher("miner")

	// Note the passed coinbase may be different with header.Coinbase.
	env := &executor_env{
		signer:   types.MakeSigner(e.chainConfig, header.Number, header.Time),
		state:    state,
		coinbase: coinbase,
		header:   header,
	}

	env.tcount = 0
	return env, nil
}

func (e *executor) sendLoop() {
	defer e.wg.Done()
	for {
		select {
		case req := <-e.newWorkCh:
			// fmt.Println("sendLoop get a newWorkCh and start send tx")
			e.sendNewTxBatch(req.interrupt, req.timestamp)
		case <-e.exitCh:
			return
		}
	}
}

func (e *executor) sendNewTxBatch(interrupt *atomic.Int32, timestamp int64) {
	// Abort committing if node is still syncing
	if e.syncing.Load() {
		return
	}

	// Set the coinbase if the worker is running or it's required
	var coinbase common.Address
	if e.isRunning() {
		coinbase = e.etherbase()
		if coinbase == (common.Address{}) {
			log.Error("Refusing to mine without etherbase")
			return
		}
	}

	work, err := e.prepareWork(&generateParams{
		timestamp:   uint64(timestamp),
		coinbase:    coinbase,
		isExecution: false,
	})
	if err != nil {
		return
	}
	e.fillTransactions(interrupt, work)
}

func (e *executor) fillTransactions(interrupt *atomic.Int32, env *executor_env) error {
	pending := e.eth.TxPool().Pending(true)

	// Split the pending transactions into locals and remotes.
	localTxs, remoteTxs := make(map[common.Address][]*txpool.LazyTransaction), pending
	for _, account := range e.eth.TxPool().Locals() {
		if txs := remoteTxs[account]; len(txs) > 0 {
			delete(remoteTxs, account)
			localTxs[account] = txs
		}
	}
	// TODO: for test
	if len(localTxs) == 0 {
		return nil
	}

	// Fill the block with all available pending transactions.
	if len(localTxs) > 0 {
		txs := newTransactionsByPriceAndNonce(env.signer, localTxs, env.header.BaseFee)
		if err := e.sendTransactions(env, txs, interrupt); err != nil {
			return err
		}
	}
	if len(remoteTxs) > 0 {
		txs := newTransactionsByPriceAndNonce(env.signer, remoteTxs, env.header.BaseFee)
		if err := e.sendTransactions(env, txs, interrupt); err != nil {
			return err
		}
	}
	return nil
}

func (e *executor) sendTransactions(env *executor_env, txs *transactionsByPriceAndNonce, interrupt *atomic.Int32) error {
	gasLimit := env.header.GasLimit
	if env.gasPool == nil {
		env.gasPool = new(core.GasPool).AddGas(gasLimit)
	}

	for {
		// Check interruption signal and abort building if it's fired.
		if interrupt != nil {
			if signal := interrupt.Load(); signal != commitInterruptNone {
				return signalToErr(signal)
			}
		}
		// If we don't have enough gas for any further transactions then we're done.
		if env.gasPool.Gas() < params.TxGas {
			log.Trace("Not enough gas for further transactions", "have", env.gasPool, "want", params.TxGas)
			break
		}
		// Retrieve the next transaction and abort if all done.
		ltx := txs.Peek()

		if ltx == nil {
			break
		}

		// If we don't have enough space for the next transaction, skip the account.
		if env.gasPool.Gas() < ltx.Gas {
			log.Trace("Not enough gas left for transaction", "hash", ltx.Hash, "left", env.gasPool.Gas(), "needed", ltx.Gas)
			txs.Pop()
			continue
		}
		// Transaction seems to fit, pull it up from the pool
		tx := ltx.Resolve()
		if tx == nil {
			log.Trace("Ignoring evicted transaction", "hash", ltx.Hash)
			txs.Pop()
			continue
		}
		// Check whether the tx is replay protected. If we're not in the EIP155 hf
		// phase, start ignoring the sender until we do.
		if tx.Protected() && !e.chainConfig.IsEIP155(env.header.Number) {
			log.Trace("Ignoring replay protected transaction", "hash", ltx.Hash, "eip155", e.chainConfig.EIP155Block)
			txs.Pop()
			continue
		}

		// sendTx to consensus
		_, err := e.execClient.sendTx(tx, e.networkId)

		if err != nil {
			log.Trace("Failed to send transaction", "hash", ltx.Hash, "err", err)
			txs.Pop()
			continue
		}
		// !!! 不然这里的gasPool没被更新
		env.gasPool.SubGas(tx.Gas())
		txs.Shift()
	}

	return nil
}

func (e *executor) executionLoop() {
	defer e.wg.Done()

	for {
		select {
		case req := <-e.execCh:
			e.executeNewTxBatch(req.timestamp, req.txs, req.leader, req.randomNumber, req.incentive)
		case <-e.exitCh:
			return
		}
	}
}

func (e *executor) executeNewTxBatch(timestamp int64, txs types.Transactions, leader common.Address, randomNumber *big.Int, incentiveData []byte) {
	var coinbase common.Address
	if e.isRunning() {
		coinbase = e.etherbase()
		if coinbase == (common.Address{}) {
			log.Error("Refusing to mine without etherbase")
			return
		}
	}

	// Split transactions into normal and external txs based on destination address
	var normalTxs, externalTxs types.Transactions
	for _, tx := range txs {
		if tx.To() != nil && *tx.To() == params.ExternalExecutionAddress {
			externalTxs = append(externalTxs, tx)
		} else {
			normalTxs = append(normalTxs, tx)
		}
	}

	// Count PoW transactions and total gas
	powCount := uint64(0)
	totalGas := uint64(0)
	txCount := uint64(len(normalTxs))

	if txCount > 0 {
		for _, tx := range normalTxs {
			if tx.Type() == types.PowTxType {
				powCount++
			}
			totalGas += tx.Gas()
		}
	}

	work, err := e.prepareWork(&generateParams{
		timestamp: uint64(timestamp),
		coinbase:  coinbase,
		currentRatio: &common.Rational{
			Numerator:   powCount,
			Denominator: txCount,
		},
		currentGasRatio: &common.Rational{
			Numerator:   totalGas,
			Denominator: txCount,
		},
		isExecution: true,
	})
	if err != nil {
		return
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go e.executeTransactions(work, normalTxs, &wg)
	go e.executeTransactions(work, externalTxs, &wg)
	wg.Wait()
	e.writeToChain(work)
}

// 串行地执行交易，会返回一个Logs，或许以后会有用
func (e *executor) executeTransactions(env *executor_env, txs types.Transactions, wg *sync.WaitGroup) []*types.Log {
	defer wg.Done()
	gasLimit := env.header.GasLimit
	if env.gasPool == nil {
		env.gasPool = new(core.GasPool).AddGas(gasLimit)
	}

	var coalescedLogs []*types.Log
	// fmt.Println("start exec,txs len:", len((txs)))
	log.Info("start execute transactions", "receive init txs len:", txs.Len())
	env.initTxcount = txs.Len()
	for _, tx := range txs {
		if isCorrectFromTransferTx(tx) {
			log.Info("Correct from transfer transaction", "hash", tx.Hash())
		} else {
			// 扔掉这个交易
			log.Info("Incorrect from transfer transaction", "hash", tx.Hash())
			continue
		}

		// If we don't have enough gas for any further transactions then we're done.
		if env.gasPool.Gas() < params.TxGas {
			log.Trace("Not enough gas for further transactions", "have", env.gasPool, "want", params.TxGas)
			break
		}
		// If we don't have enough space for the next transaction, skip.
		if env.gasPool.Gas() < tx.Gas() {
			log.Trace("Not enough gas left for transaction", "hash", tx.Hash(), "left", env.gasPool.Gas(), "needed", tx.Gas())
			continue
		}
		// Transaction seems to fit, pull it up from the pooltinue
		// Check whether the tx is replay protected. If we're not in the EIP155 hf
		// phase, start ignoring the sender until we do.
		if tx.Protected() && !e.chainConfig.IsEIP155(env.header.Number) {
			log.Trace("Ignoring replay protected transaction", "hash", tx.Hash(), "eip155", e.chainConfig.EIP155Block)
			continue
		}

		from, _ := types.Sender(env.signer, tx)
		if tx.To() != nil {
			if tx.To().Hex() == ContributionContractAddr && from != e.currenLeader {
				log.Trace("Ignoring contribution transaction because it is not from current leader", "hash", tx.Hash(), "leader", e.currenLeader)
				continue
			}
		}

		// Set interest rate and transfer interest rate
		env.state.SetInterestRate(params.InterestRate)
		env.state.SetTransferInterestRate(params.TransferInterestRate)

		env.state.SetTxContext(tx.Hash(), env.tcount)
		logs, err := e.executeTransaction(env, tx)
		switch {
		case errors.Is(err, core.ErrNonceTooLow):
			// New head notification data race between the transaction pool and miner, shift
			log.Trace("Skipping transaction with low nonce", "hash", tx.Hash, "sender", from, "nonce", tx.Nonce())
			continue

		case errors.Is(err, nil):
			// Everything ok, collect the logs and shift in the next transaction from the same account
			coalescedLogs = append(coalescedLogs, logs...)
			continue

		default:
			// Transaction is regarded as invalid, drop all consecutive transactions from
			// the same sender because of `nonce-too-high` clause.
			log.Debug("Transaction failed, account skipped", "hash", tx.Hash, "err", err)
			log.Error("Transaction failed, account skipped", "hash", tx.Hash, "err", err)
			continue
		}
	}
	return coalescedLogs
}

// 看看交易行成功没有，如果成功把它收集进Env里
func (e *executor) executeTransaction(env *executor_env, tx *types.Transaction) ([]*types.Log, error) {
	// TODO : send to transfer
	if isToTransferTransaction(tx) {
		from, err := types.Sender(env.signer, tx)
		if err != nil {
			return nil, err
		}
		e.execClient.transferClient.ToTransferCommit(context.Background(), &pb.ToTransferRequest{FromAddress: from.Bytes(), BAddress: tx.To().Bytes(), Amount: int32(tx.Value().Int64())})
		// 这里依然要执行，然后产生链上记录给转账区进行查询验证。
	}

	// 检查交易是否是代币转换并调用dciClient校验函数
	log.Info("check token transition")
	if isTokenTransition(tx) {
		log.Info("verify token transition transaction")
		flag, err := e.execClient.verifyTokenTransitionTx(tx)
		if err != nil {
			log.Error("Failed to verify token transition transaction", "err", err)
			return nil, err
		}
		if !flag {
			log.Error("Failed to verify token transition transaction, flag is false")
			return nil, errors.New("failed to verify token transition transaction flag is false")
		}
	}

	// TODO: may not need this
	// offchain tx executor
	// The first three bytes, determine the transaction type, and remove the field that identifies the type
	data := tx.Data()
	if len(data) >= 3 {
		if data[0] == 0x0A && data[1] == 0x0D {
			fmt.Printf("Transaction type:%v\n", data[2])
			switch data[2] {
			case 1:
				// 1.remove identifies field   2. go routine: push result in to channel
				go e.offchainCom(data[3:])
			case 2:
				// offchain seconde request: catch the result of offchainCalc from channel
				e.offchainResultCatch(env)
			}
		} else {
			log.Info("Tx is not offchain type")
		}
	}

	receipt, err := e.applyTransaction(env, tx)
	if err != nil {
		return nil, err
	}
	env.txs = append(env.txs, tx)
	env.receipts = append(env.receipts, receipt)
	env.tcount++
	log.Info("exec transaction success", "hash", tx.Hash())
	return receipt.Logs, nil
}

// applyTransaction runs the transaction. If execution fails, state and gas pool are reverted.
// 真正在执行一笔交易
func (e *executor) applyTransaction(env *executor_env, tx *types.Transaction) (*types.Receipt, error) {
	var (
		snap = env.state.Snapshot()
		gp   = env.gasPool.Gas()
	)
	receipt, err := core.ApplyTransaction(e.chainConfig, e.eth.BlockChain(), &env.coinbase, env.gasPool, env.state, env.header, tx, &env.header.GasUsed, *e.eth.BlockChain().GetVMConfig())
	if err != nil {
		env.state.RevertToSnapshot(snap)
		env.gasPool.SetGas(gp)
	}
	return receipt, err
}

func (e *executor) writeToChain(env *executor_env) error {
	// 插入header的新数据
	env.header.CommitTxLength = uint64(env.initTxcount)

	// 组装一个区块
	block, err := e.engine.FinalizeAndAssemble(e.eth.BlockChain(), env.header, env.state, env.txs, nil, env.receipts, nil)
	if err != nil {
		return err
	}
	var (
		receipts = make([]*types.Receipt, len(env.receipts))
		logs     []*types.Log
	)
	hash := block.Hash()
	// fmt.Println(hash)
	// 拷贝一下env.receipts
	for i, env_receipt := range env.receipts {
		receipt := new(types.Receipt)
		receipts[i] = receipt
		*receipt = *env_receipt

		receipt.BlockHash = hash
		receipt.BlockNumber = block.Number()
		receipt.TransactionIndex = uint(i)

		receipt.Logs = make([]*types.Log, len(env_receipt.Logs))
		for i, env_log := range env_receipt.Logs {
			log := new(types.Log)
			receipt.Logs[i] = log
			*log = *env_log
			log.BlockHash = hash
		}
		logs = append(logs, receipt.Logs...)
	}
	// Commit block and state to database.
	_, err = e.eth.BlockChain().WriteBlockAndSetHead(block, receipts, logs, env.state, true)
	if err != nil {
		log.Error("Failed writing block to chain", "err", err)
		return err
	}

	// fmt.Println(e.eth.BlockChain().CurrentBlock().Number)
	log.Info("Successfully sealed new block", "number", block.Number(), "hash", hash)
	// 比较有信心说，这就是我的env
	e.env = env.copy()

	// emit broadcast
	e.mux.Post(core.NewMinedBlockEvent{Block: block, NetworkID: e.networkId})

	return nil
}

func isTokenTransition(tx *types.Transaction) bool {
	if tx.Data() == nil || len(tx.Data()) < 3 {
		return false
	}

	if tx.Data()[0] == 0x0D && tx.Data()[1] == 0x02 {
		return true
	}
	return false
}

func isToTransferTransaction(tx *types.Transaction) bool {
	if tx.Data() == nil || len(tx.Data()) < 3 {
		return false
	}

	if tx.Data()[0] == 0x0D && tx.Data()[1] == 0x07 {
		return true
	}
	return false
}

func isCorrectFromTransferTx(tx *types.Transaction) bool {
	if tx.Type() == types.SystemTxType && tx.Data()[0] == 0x0D && tx.Data()[1] == 0x06 {
		// 读取tx.Data()[2:]作为verifyHash
		verifyHash := tx.Data()[2:]
		fmt.Println("verifyHash", verifyHash)
		// 调用转账接口进行验证
		// e.execClient.verifyTransferTx(verifyHash)
		return true
	}
	return false
}

func (e *executor) ExecuteAndRemovePlan(height uint64) error {
	// Get merged plan from pool
	mergedPlan := e.planPool.MergePlans(height)
	if mergedPlan == nil {
		return nil
	}

	// Update parameters
	e.planPool.UpdateParams(mergedPlan)

	// Recreate adaptors with new parameters
	e.powAdaptor = NewPoWAdaptor(
		params.TargetPowRatio,
		params.Alpha,
		params.Fmin,
		params.Fmax,
		params.InitialDifficulty,
		params.Kp,
		params.Ki,
		params.MinPrice,
		params.MaxPrice,
	)

	e.gasAdaptor = NewGasAdaptor(
		params.MinPowGas,
		params.MaxPowGas,
		params.InitialGas,
		params.Alpha,
	)

	e.planPool.RemovePlan(height)

	return nil
}
