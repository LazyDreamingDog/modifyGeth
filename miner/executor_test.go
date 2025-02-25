package miner

import (
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/consensus/clique"
	"github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/txpool"
	"github.com/ethereum/go-ethereum/core/txpool/legacypool"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/proto/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	// testCode is the testing contract binary code which will initialises some
	// variables in constructor
	testCode = "0x60806040527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0060005534801561003457600080fd5b5060fc806100436000396000f3fe6080604052348015600f57600080fd5b506004361060325760003560e01c80630c4dae8814603757806398a213cf146053575b600080fd5b603d607e565b6040518082815260200191505060405180910390f35b607c60048036036020811015606757600080fd5b81019080803590602001909291905050506084565b005b60005481565b806000819055507fe9e44f9f7da8c559de847a3232b57364adc0354f15a2cd8dc636d54396f9587a6000546040518082815260200191505060405180910390a15056fea265627a7a723058208ae31d9424f2d0bc2a3da1a5dd659db2d71ec322a17db8f87e19e209e3a1ff4a64736f6c634300050a0032"

	// testGas is the gas required for contract deployment.
	testGas = 144109
)

var (
	// Test chain configurations
	testTxPoolConfig  legacypool.Config
	ethashChainConfig *params.ChainConfig
	cliqueChainConfig *params.ChainConfig

	// Test accounts
	testBankKey, _  = crypto.GenerateKey()
	testBankAddress = crypto.PubkeyToAddress(testBankKey.PublicKey)
	testBankFunds   = big.NewInt(1000000000000000000)

	testUserKey, _  = crypto.GenerateKey()
	testUserAddress = crypto.PubkeyToAddress(testUserKey.PublicKey)

	// Test transactions
	pendingTxs []*types.Transaction
	newTxs     []*types.Transaction

	testConfig = &Config{
		Recommit: time.Second,
		GasCeil:  params.GenesisGasLimit,
	}
)

func init() {
	testTxPoolConfig = legacypool.DefaultConfig
	testTxPoolConfig.Journal = ""
	ethashChainConfig = new(params.ChainConfig)
	*ethashChainConfig = *params.TestChainConfig
	cliqueChainConfig = new(params.ChainConfig)
	*cliqueChainConfig = *params.TestChainConfig
	cliqueChainConfig.Clique = &params.CliqueConfig{
		Period: 10,
		Epoch:  30000,
	}

	signer := types.LatestSigner(params.TestChainConfig)
	tx1 := types.MustSignNewTx(testBankKey, signer, &types.AccessListTx{
		ChainID:  params.TestChainConfig.ChainID,
		Nonce:    0,
		To:       &testUserAddress,
		Value:    big.NewInt(1000),
		Gas:      params.TxGas,
		GasPrice: big.NewInt(params.InitialBaseFee),
	})
	// pendingTxs = make([]*types.Transaction, 1)
	pendingTxs = append(pendingTxs, tx1)

	tx2 := types.MustSignNewTx(testBankKey, signer, &types.LegacyTx{
		Nonce:    1,
		To:       &testUserAddress,
		Value:    big.NewInt(1000),
		Gas:      params.TxGas,
		GasPrice: big.NewInt(params.InitialBaseFee),
	})
	newTxs = append(newTxs, tx2)
}

// testWorkerBackend implements worker.Backend interfaces and wraps all information needed during the testing.
type testWorkerBackend struct {
	db      ethdb.Database
	txPool  *txpool.TxPool
	chain   *core.BlockChain
	genesis *core.Genesis
}

func (b *testWorkerBackend) BlockChain() *core.BlockChain      { return b.chain }
func (b *testWorkerBackend) TxPool() *txpool.TxPool            { return b.txPool }
func (b *testWorkerBackend) NetworkId() uint64                 { return 1 }
func (b *testWorkerBackend) AccountManager() *accounts.Manager { return nil }

func (b *testWorkerBackend) newTx(nonce uint64) *types.Transaction {
	signer := types.LatestSigner(b.chain.Config())
	tx := types.MustSignNewTx(testBankKey, signer, &types.DynamicFeeTx{
		ChainID:   b.chain.Config().ChainID,
		Nonce:     nonce,
		To:        &testUserAddress,
		Value:     big.NewInt(1000),
		Gas:       params.TxGas,
		GasFeeCap: big.NewInt(10 * params.InitialBaseFee),
		GasTipCap: big.NewInt(0),
	})

	return tx
}

func (b *testWorkerBackend) newContractTx(nonce uint64) *types.Transaction {
	signer := types.LatestSigner(b.chain.Config())
	tx := types.MustSignNewTx(testBankKey, signer, &types.DynamicFeeTx{
		ChainID:   b.chain.Config().ChainID,
		Nonce:     nonce,
		Value:     big.NewInt(0),
		Data:      common.FromHex(testCode),
		Gas:       testGas,
		GasFeeCap: big.NewInt(10 * params.InitialBaseFee),
		GasTipCap: big.NewInt(0),
	})

	return tx
}

func newTestExecBackend(chainConfig *params.ChainConfig, engine consensus.Engine, db ethdb.Database, n int) *testWorkerBackend {
	var gspec = &core.Genesis{
		Config: chainConfig,
		Alloc:  core.GenesisAlloc{testBankAddress: {Balance: testBankFunds}},
	}
	switch e := engine.(type) {
	case *clique.Clique:
		gspec.ExtraData = make([]byte, 32+common.AddressLength+crypto.SignatureLength)
		copy(gspec.ExtraData[32:32+common.AddressLength], testBankAddress.Bytes())
		e.Authorize(testBankAddress, func(account accounts.Account, s string, data []byte) ([]byte, error) {
			return crypto.Sign(crypto.Keccak256(data), testBankKey)
		})
	case *ethash.Ethash:
	default:
		panic(fmt.Sprintf("unsupported consensus engine: %T", e))
	}
	// Blockchain 中 engine 只管 VerifyHeader
	chain, err := core.NewBlockChain(db, &core.CacheConfig{TrieDirtyDisabled: true}, gspec, nil, engine, vm.Config{}, nil, nil)
	if err != nil {
		panic(err)
	}
	pool := legacypool.New(testTxPoolConfig, chain)
	txpool, _ := txpool.New(new(big.Int).SetUint64(testTxPoolConfig.PriceLimit), chain, []txpool.SubPool{pool})

	return &testWorkerBackend{
		db:      db,
		chain:   chain,
		txPool:  txpool,
		genesis: gspec,
	}
}

func newTestExecutor(chainConfig *params.ChainConfig, engine consensus.Engine, db ethdb.Database, blocks int) (*executor, *testWorkerBackend) {
	backend := newTestExecBackend(chainConfig, engine, db, blocks)
	// fmt.Println(pendingTxs)
	// errs := backend.txPool.Add(pendingTxs, true, false)
	// fmt.Println("acltx errs", errs)
	// 实例化共识客户端
	conn, err := grpc.Dial("127.0.0.1:9080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
	}
	p2pClient := pb.NewP2PClient(conn)

	conn1, err := grpc.Dial("127.0.0.1:1145", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
	}

	conn2, err := grpc.Dial("127.0.0.1:2233", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
	}

	dciClient := pb.NewBciExectorClient(conn2)
	transferClient := pb.NewTransferGRPCClient(conn1)
	e := newExecutor(testConfig, chainConfig, engine, backend, new(event.TypeMux), nil, false, p2pClient, transferClient, dciClient)
	e.coinbase = testBankAddress
	return e, backend
}

func newTestPotExecutor(chainConfig *params.ChainConfig, engine consensus.Engine, db ethdb.Database) *poter {
	backend := newTestExecBackend(chainConfig, engine, db, 0)

	conn, err := grpc.Dial("127.0.0.1:9081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
	}

	potExcutorClient := pb.NewPoTExecutorClient(conn)
	p := newPoter(backend, potExcutorClient)

	return p
}

func TestExecutor(t *testing.T) {
	var (
		db     = rawdb.NewMemoryDatabase()
		config = *params.AllCliqueProtocolChanges
	)
	config.Clique = &params.CliqueConfig{Period: 1, Epoch: 30000}
	engine := clique.New(config.Clique, db)

	e, b := newTestExecutor(&config, engine, db, 0)
	defer e.close()

	// This test chain imports the mined blocks.
	// chain, _ := core.NewBlockChain(rawdb.NewMemoryDatabase(), nil, b.genesis, nil, engine, vm.Config{}, nil, nil)
	// defer chain.Stop()

	// Start mining!
	e.start()
	fmt.Println("11111111111")
	for i := 0; i < 3; i++ {
		// fmt.Println("add tx")
		// testtx := b.newRandomTx(false)
		testtx := b.newTx(uint64(i))
		// testtx := b.newContractTx(uint64(i))
		errs := b.txPool.Add([]*types.Transaction{testtx}, true, false)
		fmt.Println("errs", errs)
		// b.txPool.Add([]*types.Transaction{b.newRandomTx(false)}, true, false)
		time.Sleep(10 * time.Second)
	}
	time.Sleep(50 * time.Second)
}

func TestPotExecutor(t *testing.T) {
	var (
		db     = rawdb.NewMemoryDatabase()
		config = *params.AllCliqueProtocolChanges
	)
	config.Clique = &params.CliqueConfig{Period: 1, Epoch: 30000}
	engine := clique.New(config.Clique, db)

	e, b := newTestExecutor(&config, engine, db, 0)
	defer e.close()

	p := newTestPotExecutor(&config, engine, db)
	defer p.close()

	// This test chain imports the mined blocks.
	// chain, _ := core.NewBlockChain(rawdb.NewMemoryDatabase(), nil, b.genesis, nil, engine, vm.Config{}, nil, nil)
	// defer chain.Stop()

	// Start mining!
	e.start()
	p.start()
	fmt.Println("111111")
	for i := 0; i < 1; i++ {
		// fmt.Println("add tx")
		// testtx := b.newRandomTx(false)
		// testtx := b.newTx(uint64(i))
		testtx := b.newContractTx(uint64(i))
		errs := b.txPool.Add([]*types.Transaction{testtx}, true, false)
		fmt.Println("errs", errs)
		// b.txPool.Add([]*types.Transaction{b.newRandomTx(false)}, true, false)
		time.Sleep(10 * time.Second)
	}
	time.Sleep(50 * time.Second)
}
