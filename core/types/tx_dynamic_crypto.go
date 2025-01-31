package types

import (
	"bytes"
	"io"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

// DynamicCryptoTx represents a new transaction type with additional fields: CryptoType and SignatureData.
type DynamicCryptoTx struct {
	ChainID        *big.Int
	Nonce          uint64
	GasTipCap      *big.Int // a.k.a. maxPriorityFeePerGas
	GasFeeCap      *big.Int // a.k.a. maxFeePerGas
	Gas            uint64
	To             *common.Address `rlp:"nil"` // nil means contract creation
	Value          *big.Int
	Data           []byte
	AccessList     AccessList
	CryptoType     []byte // New field to represent the type of cryptocurrency
	SignatureData  []byte // New field for signature data
	PublicKeyIndex uint64 // New field for public index
	PublicKey      []byte // New field for public key

	// Signature values
	V *big.Int `json:"v" gencodec:"required"`
	R *big.Int `json:"r" gencodec:"required"`
	S *big.Int `json:"s" gencodec:"required"`
}

// copy creates a deep copy of the transaction data and initializes all fields.
func (tx *DynamicCryptoTx) copy() TxData {
	cpy := &DynamicCryptoTx{
		Nonce:          tx.Nonce,
		To:             copyAddressPtr(tx.To),
		Data:           common.CopyBytes(tx.Data),
		Gas:            tx.Gas,
		CryptoType:     common.CopyBytes(tx.CryptoType),
		SignatureData:  common.CopyBytes(tx.SignatureData),
		PublicKeyIndex: tx.PublicKeyIndex,
		PublicKey:      common.CopyBytes(tx.PublicKey),
		// These are copied below.
		AccessList: make(AccessList, len(tx.AccessList)),
		Value:      new(big.Int),
		ChainID:    new(big.Int),
		GasTipCap:  new(big.Int),
		GasFeeCap:  new(big.Int),
		V:          new(big.Int),
		R:          new(big.Int),
		S:          new(big.Int),
	}
	copy(cpy.AccessList, tx.AccessList)
	if tx.Value != nil {
		cpy.Value.Set(tx.Value)
	}
	if tx.ChainID != nil {
		cpy.ChainID.Set(tx.ChainID)
	}
	if tx.GasTipCap != nil {
		cpy.GasTipCap.Set(tx.GasTipCap)
	}
	if tx.GasFeeCap != nil {
		cpy.GasFeeCap.Set(tx.GasFeeCap)
	}
	if tx.V != nil {
		cpy.V.Set(tx.V)
	}
	if tx.R != nil {
		cpy.R.Set(tx.R)
	}
	if tx.S != nil {
		cpy.S.Set(tx.S)
	}
	return cpy
}

// accessors for innerTx.
func (tx *DynamicCryptoTx) txType() byte           { return DynamicCryptoTxType }
func (tx *DynamicCryptoTx) chainID() *big.Int      { return tx.ChainID }
func (tx *DynamicCryptoTx) accessList() AccessList { return tx.AccessList }
func (tx *DynamicCryptoTx) data() []byte           { return tx.Data }
func (tx *DynamicCryptoTx) gas() uint64            { return tx.Gas }
func (tx *DynamicCryptoTx) gasFeeCap() *big.Int    { return tx.GasFeeCap }
func (tx *DynamicCryptoTx) gasTipCap() *big.Int    { return tx.GasTipCap }
func (tx *DynamicCryptoTx) gasPrice() *big.Int     { return tx.GasFeeCap }
func (tx *DynamicCryptoTx) value() *big.Int        { return tx.Value }
func (tx *DynamicCryptoTx) nonce() uint64          { return tx.Nonce }
func (tx *DynamicCryptoTx) to() *common.Address    { return tx.To }
func (tx *DynamicCryptoTx) cryptoType() []byte     { return tx.CryptoType }
func (tx *DynamicCryptoTx) signatureData() []byte  { return tx.SignatureData }
func (tx *DynamicCryptoTx) publicKeyIndex() uint64 { return tx.PublicKeyIndex }
func (tx *DynamicCryptoTx) publicKey() []byte      { return tx.PublicKey }

func (tx *DynamicCryptoTx) effectiveGasPrice(dst *big.Int, baseFee *big.Int) *big.Int {
	if baseFee == nil {
		return dst.Set(tx.GasFeeCap)
	}
	tip := dst.Sub(tx.GasFeeCap, baseFee)
	if tip.Cmp(tx.GasTipCap) > 0 {
		tip.Set(tx.GasTipCap)
	}
	return tip.Add(tip, baseFee)
}

func (tx *DynamicCryptoTx) rawSignatureValues() (v, r, s *big.Int) {
	return tx.V, tx.R, tx.S
}

func (tx *DynamicCryptoTx) setSignatureValues(chainID, v, r, s *big.Int) {
	tx.ChainID, tx.V, tx.R, tx.S = chainID, v, r, s
}

// dynamicCryptoTxData 用于序列化/反序列化的数据结构
type dynamicCryptoTxData struct {
	AccountNonce   uint64
	Price          *big.Int
	GasLimit       uint64
	Recipient      *common.Address
	Amount         *big.Int
	Payload        []byte
	V              *big.Int
	R              *big.Int
	S              *big.Int
	ChainID        *big.Int
	CryptoType     []byte
	SignatureData  []byte
	PublicKeyIndex uint64
	PublicKey      []byte
}

// EncodeRLP implements rlp.Encoder
func (tx *DynamicCryptoTx) EncodeRLP(w io.Writer) error {
	enc := &dynamicCryptoTxData{
		AccountNonce:   tx.Nonce,
		Price:          tx.GasTipCap,
		GasLimit:       tx.Gas,
		Recipient:      tx.To,
		Amount:         tx.Value,
		Payload:        tx.Data,
		V:              tx.V,
		R:              tx.R,
		S:              tx.S,
		ChainID:        tx.ChainID,
		CryptoType:     tx.CryptoType,
		SignatureData:  tx.SignatureData,
		PublicKeyIndex: tx.PublicKeyIndex,
		PublicKey:      tx.PublicKey,
	}
	return rlp.Encode(w, enc)
}

// DecodeRLP implements rlp.Decoder
func (tx *DynamicCryptoTx) DecodeRLP(s *rlp.Stream) error {
	var dec dynamicCryptoTxData
	if err := s.Decode(&dec); err != nil {
		return err
	}
	tx.Nonce = dec.AccountNonce
	tx.GasTipCap = dec.Price
	tx.Gas = dec.GasLimit
	tx.To = dec.Recipient
	tx.Value = dec.Amount
	tx.Data = dec.Payload
	tx.V = dec.V
	tx.R = dec.R
	tx.S = dec.S
	tx.ChainID = dec.ChainID
	tx.CryptoType = dec.CryptoType
	tx.SignatureData = dec.SignatureData
	tx.PublicKeyIndex = dec.PublicKeyIndex
	tx.PublicKey = dec.PublicKey
	return nil
}

// 补全RLP编码
func (tx *DynamicCryptoTx) encode(b *bytes.Buffer) error {
	return rlp.Encode(b, tx)
}

func (tx *DynamicCryptoTx) decode(input []byte) error {
	return rlp.DecodeBytes(input, tx)
}
