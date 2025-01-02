package miner

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/consensus/clique"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/params"
)

func newTestTransfer(chainConfig *params.ChainConfig, engine consensus.Engine, db ethdb.Database) (*Transfer, *testWorkerBackend) {
	backend := newTestExecBackend(chainConfig, engine, db, 0)

	t := newTransfer(backend)

	return t, backend
}

func TestAddSystemTx(t *testing.T) {
	fmt.Println("TestAddSystemTx")
	fmt.Println(types.LegacyTxType)
	fmt.Println(types.SystemTxType)
	var (
		db     = rawdb.NewMemoryDatabase()
		config = *params.AllCliqueProtocolChanges
	)
	config.Clique = &params.CliqueConfig{Period: 1, Epoch: 30000}
	engine := clique.New(config.Clique, db)

	transfer, backend := newTestTransfer(&config, engine, db)
	defer transfer.close()

	key, _ := crypto.GenerateKey()
	addr := crypto.PubkeyToAddress(key.PublicKey)

	systemTx := types.NewSystemTx(
		big.NewInt(1),
		123,
		big.NewInt(1),
		big.NewInt(200),
		50000,
		&addr,
		big.NewInt(1000),
		[]byte{},
		1,
	)
	fmt.Println(systemTx.Hash())
	errs := backend.txPool.Add([]*types.Transaction{systemTx}, true, false)

	fmt.Println(errs)

	fmt.Println(systemTx.Hash())
}
