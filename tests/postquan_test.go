package tests

import (
	"fmt"
	"math/big"
	"teddycode/pqcgo"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
)

func newPostQuanTx(bc *core.BlockChain, lastPub *[]byte, signer types.Signer, Nonce int, to *common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	// Prepare
	cryptoType := "aigis_sig"
	scheme := pqcgo.PQCSignType[cryptoType]

	var pk, sk []byte
	var err error
	if *lastPub == nil || len(*lastPub) == 0 {
		pk, sk, err = pqcgo.KeyGen(scheme)
	} else {
		fmt.Println(*lastPub)
		pk, sk, err = pqcgo.KeyGenWithSeed(scheme, *lastPub)
	}

	// LastPub update

	if err != nil {
		return nil, err
	}
	txData := &types.DynamicCryptoTx{
		ChainID:    bc.Config().ChainID,
		Nonce:      uint64(Nonce),
		To:         to,
		Value:      value,
		Gas:        50000,
		Data:       data,
		GasFeeCap:  big.NewInt(8750000000),
		GasTipCap:  big.NewInt(444444),
		AccessList: nil,
		// Post-quantum attributes, these four fields are ont computed in hash at func @HashExcludePostQumSign
		// But are computed in hash in @siger.Hash()
		CryptoType:    []byte(cryptoType),
		SignatureData: nil,
		PostAddress:   &common.Address{},
		PublicKey:     pk,
	}

	// Post-quantum signature
	tx_temp := types.NewTx(txData)
	txHash := tx_temp.HashExcludePostQumSign()

	// Post-quantum sign
	sig, err := pqcgo.Sign(scheme, txHash.Bytes(), sk)
	if err != nil {
		return nil, err
	}
	txData.SignatureData = sig

	// ETH signature
	tx := types.MustSignNewTx(bankKey, signer, txData)
	return tx, nil
}

// Construct TX to test voucher
func TestPostQuan(t *testing.T) {
	backend := newTestBackend()
	miner := backend.CreateMiner()
	miner.Start()
	defer miner.Stop()

	lastPub := make([]byte, 0)
	testDes := common.BytesToAddress([]byte{67})

	tx1, err := newPostQuanTx(backend.bc, &lastPub, backend.newSigner(), 0, &testDes, big.NewInt(0), nil)
	if err != nil {
		t.Fatalf("Create tx error:%v", err)
	}
	backend.AddTx(tx1)

	tx2, err := newPostQuanTx(backend.bc, &lastPub, backend.newSigner(), 1, &testDes, big.NewInt(0), nil)
	if err != nil {
		t.Fatalf("Create tx error:%v", err)
	}
	backend.AddTx(tx2)
}
