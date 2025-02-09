// * This file store global variable for @package tests

package tests

import (
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/clique"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/txpool"
	"github.com/ethereum/go-ethereum/core/txpool/legacypool"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/miner"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/trie"
)

// Global variable
var (
	// Create bank account
	bankKey, _  = crypto.GenerateKey()
	bankAddress = crypto.PubkeyToAddress(bankKey.PublicKey)
	// Create user account
	userKey, _  = crypto.GenerateKey()
	userAddress = crypto.PubkeyToAddress(userKey.PublicKey)
	// Contract setting, parse from json which is created by hardhat
	contractAbi, contractBytecode, deployedBytecode = abi.LoadHardhatContract("../voucher/abi/mutivoucher.json")
	contractAddress                                 = common.BytesToAddress([]byte{68})
)

// Implements the interface of miner.Backend
type TestBackend struct {
	bc      *core.BlockChain
	txpool  *txpool.TxPool
	genesis *core.Genesis
}

// ! NetworkID must be 1. POT sharding require
func (m *TestBackend) NetworkId() uint64            { return 1 }
func (m *TestBackend) BlockChain() *core.BlockChain { return m.bc }
func (m *TestBackend) TxPool() *txpool.TxPool       { return m.txpool }
func (m *TestBackend) AccountManager() *accounts.Manager {
	return accounts.NewManager(&accounts.Config{})
}

func (m *TestBackend) AddTx(tx *types.Transaction) {
	errs := m.txpool.Add([]*types.Transaction{tx}, true, false)
	if len(errs) != 0 {
		fmt.Printf("Add tx err: %v\n", errs)
	}
	// Wait for consense
	time.Sleep(30 * time.Second)
}

func (m *TestBackend) parseContractAddress() *common.Address {
	var contractAddress *common.Address
	// Output every one second
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	done := make(chan bool)

	// Wait 5s, write to done
	go func() {
		time.Sleep(5 * time.Second)
		done <- true
	}()
	for {
		select {
		case <-done:
			return nil
		case <-ticker.C:
			latestHeader := m.bc.CurrentBlock()
			receipts := m.bc.GetReceiptsByHash(latestHeader.Hash())
			if receipts == nil {
				fmt.Println("no receipts in latest block")
			} else {
				for i := range receipts {
					// Search the receipt which contractAddress is not none
					if receipts[i].ContractAddress != (common.Address{}) {
						fmt.Printf("Contract Address: %v \n", receipts[i].ContractAddress)
						contractAddress = &receipts[i].ContractAddress
						return contractAddress
					}
				}
			}
		}
	}
}

func (m *TestBackend) newSigner() types.Signer {
	header := m.bc.CurrentBlock()
	return types.MakeSigner(m.bc.Config(), header.Number, header.Time)
}

func (m *TestBackend) CreateMiner() *miner.Miner {
	// ! EtherBase is necessary
	minerConfig := miner.Config{
		Etherbase: common.HexToAddress("123456789"),
	}
	miner := miner.New(m, &minerConfig, m.bc.Config(), new(event.TypeMux), m.bc.Engine(), nil)
	return miner
}

func newTestBackend() *TestBackend {
	bc := newBlockChain()
	txpool := newTxPool(bc)
	genesis := genesisBlock()
	return &TestBackend{
		bc:      bc,
		txpool:  txpool,
		genesis: genesis,
	}
}

// Gensis block
func genesisBlock() *core.Genesis {
	config := *params.AllCliqueProtocolChanges
	config.Clique = &params.CliqueConfig{
		Period: config.Clique.Period,
		Epoch:  config.Clique.Epoch,
	}
	config.ChainID = big.NewInt(1)
	// Assemble and return the genesis with the precompiles and faucet pre-funded
	return &core.Genesis{
		Config:     &config,
		ExtraData:  append(append(make([]byte, 32), bankAddress[:]...), make([]byte, crypto.SignatureLength)...),
		GasLimit:   params.MaxGasLimit,
		BaseFee:    big.NewInt(params.InitialBaseFee),
		Difficulty: big.NewInt(1),
		Alloc: map[common.Address]core.GenesisAccount{
			bankAddress:     {Balance: new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(9))},
			userAddress:     {Balance: big.NewInt(1000000)},
			contractAddress: {Code: common.FromHex(deployedBytecode)},
		},
	}
}

// Create new blockchain to test
func newBlockChain() *core.BlockChain {
	database := rawdb.NewMemoryDatabase()
	triedb := trie.NewDatabase(database, nil)
	genesis := genesisBlock()
	chainConfig, _, err := core.SetupGenesisBlock(database, triedb, genesis)
	if err != nil {
		fmt.Printf("can't create new chain config: %v\n", err)
	}
	engine := clique.New(chainConfig.Clique, database)
	// Create blockchain
	bc, err := core.NewBlockChain(database, nil, genesis, nil, engine, vm.Config{}, nil, nil)
	if err != nil {
		fmt.Printf("can't create new chain %v\n", err)
	}
	return bc
}

// Create new txpool basing blockchain
func newTxPool(bc *core.BlockChain) *txpool.TxPool {
	testTxPoolConfig := legacypool.DefaultConfig
	pool := legacypool.New(testTxPoolConfig, bc)
	txpool, _ := txpool.New(new(big.Int).SetUint64(testTxPoolConfig.PriceLimit), bc, []txpool.SubPool{pool})
	return txpool
}

// Create new EVM instance
func newEVM(bc *core.BlockChain) *vm.EVM {
	// Create stateDB
	header := bc.CurrentBlock()
	stateDB, err := bc.StateAt(header.Root)
	if err != nil {
		fmt.Printf("failed to get stateDB: %v\n", err)
	}
	// Create EVM
	blockContext := core.NewEVMBlockContext(header, bc, nil)
	evm := vm.NewEVM(blockContext, vm.TxContext{}, stateDB, bc.Config(), vm.Config{})
	return evm
}

// Constrcut function to depoly contract
func newDeployContractTx(bc *core.BlockChain, Nonce int) *types.Transaction {
	signer := types.LatestSigner(bc.Config())
	tx0 := types.MustSignNewTx(bankKey, signer, &types.AccessListTx{
		ChainID:  bc.Config().ChainID,
		Nonce:    uint64(Nonce),
		Gas:      529992,
		GasPrice: big.NewInt(params.InitialBaseFee),
		Data:     common.FromHex(contractBytecode),
	})
	return tx0
}
