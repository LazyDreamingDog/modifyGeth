package types

// import (
// 	"bytes"
// 	"io"
// 	"math/big"

// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/rlp"
// )

// // Todo:该交易后续修改为只有leader才可以产生
// type SystemTx struct {
// 	ChainID    *big.Int
// 	Nonce      uint64
// 	GasTipCap  *big.Int // a.k.a. maxPriorityFeePerGas
// 	GasFeeCap  *big.Int // a.k.a. maxFeePerGas
// 	Gas        uint64
// 	To         *common.Address `rlp:"nil"` // nil means contract creation
// 	Value      *big.Int
// 	Data       []byte
// 	AccessList AccessList

// 	SystemFlag uint64

// 	// Signature values
// 	V *big.Int `json:"v" gencodec:"required"`
// 	R *big.Int `json:"r" gencodec:"required"`
// 	S *big.Int `json:"s" gencodec:"required"`
// }

// // copy creates a deep copy of the transaction data and initializes all fields.
// func (tx *SystemTx) copy() TxData {
// 	cpy := &SystemTx{
// 		Nonce: tx.Nonce,
// 		To:    copyAddressPtr(tx.To),
// 		Data:  common.CopyBytes(tx.Data),
// 		Gas:   tx.Gas,
// 		// These are copied below.
// 		SystemFlag: tx.SystemFlag,
// 		AccessList: make(AccessList, len(tx.AccessList)),
// 		Value:      new(big.Int),
// 		ChainID:    new(big.Int),
// 		GasTipCap:  new(big.Int),
// 		GasFeeCap:  new(big.Int),
// 		V:          new(big.Int),
// 		R:          new(big.Int),
// 		S:          new(big.Int),
// 	}
// 	copy(cpy.AccessList, tx.AccessList)
// 	if tx.Value != nil {
// 		cpy.Value.Set(tx.Value)
// 	}
// 	if tx.ChainID != nil {
// 		cpy.ChainID.Set(tx.ChainID)
// 	}
// 	if tx.GasTipCap != nil {
// 		cpy.GasTipCap.Set(tx.GasTipCap)
// 	}
// 	if tx.GasFeeCap != nil {
// 		cpy.GasFeeCap.Set(tx.GasFeeCap)
// 	}
// 	if tx.V != nil {
// 		cpy.V.Set(tx.V)
// 	}
// 	if tx.R != nil {
// 		cpy.R.Set(tx.R)
// 	}
// 	if tx.S != nil {
// 		cpy.S.Set(tx.S)
// 	}
// 	return cpy
// }

// // accessors for innerTx.
// func (tx *SystemTx) txType() byte           { return SystemTxType }
// func (tx *SystemTx) chainID() *big.Int      { return tx.ChainID }
// func (tx *SystemTx) accessList() AccessList { return tx.AccessList }
// func (tx *SystemTx) data() []byte           { return tx.Data }
// func (tx *SystemTx) gas() uint64            { return tx.Gas }
// func (tx *SystemTx) gasFeeCap() *big.Int    { return tx.GasFeeCap }
// func (tx *SystemTx) gasTipCap() *big.Int    { return tx.GasTipCap }
// func (tx *SystemTx) gasPrice() *big.Int     { return tx.GasFeeCap }
// func (tx *SystemTx) value() *big.Int        { return tx.Value }
// func (tx *SystemTx) nonce() uint64          { return tx.Nonce }
// func (tx *SystemTx) to() *common.Address    { return tx.To }
// func (tx *SystemTx) systemFlag() uint64     { return tx.SystemFlag }

// func (tx *SystemTx) effectiveGasPrice(dst *big.Int, baseFee *big.Int) *big.Int {
// 	if baseFee == nil {
// 		return dst.Set(tx.GasFeeCap)
// 	}
// 	tip := dst.Sub(tx.GasFeeCap, baseFee)
// 	if tip.Cmp(tx.GasTipCap) > 0 {
// 		tip.Set(tx.GasTipCap)
// 	}
// 	return tip.Add(tip, baseFee)
// }

// func (tx *SystemTx) rawSignatureValues() (v, r, s *big.Int) {
// 	return tx.V, tx.R, tx.S
// }

// func (tx *SystemTx) setSignatureValues(chainID, v, r, s *big.Int) {
// 	tx.ChainID, tx.V, tx.R, tx.S = chainID, v, r, s
// }

// // systemTxData 用于序列化/反序列化的数据结构
// type systemTxData struct {
// 	AccountNonce uint64
// 	Price        *big.Int
// 	GasLimit     uint64
// 	Recipient    *common.Address
// 	Amount       *big.Int
// 	Payload      []byte
// 	V            *big.Int
// 	R            *big.Int
// 	S            *big.Int
// 	ChainID      *big.Int
// 	SysFlag      uint64 // 系统交易标识
// }

// // EncodeRLP implements rlp.Encoder
// func (tx *SystemTx) EncodeRLP(w io.Writer) error {
// 	enc := &systemTxData{
// 		AccountNonce: tx.Nonce,
// 		Price:        tx.GasTipCap,
// 		GasLimit:     tx.Gas,
// 		Recipient:    tx.To,
// 		Amount:       tx.Value,
// 		Payload:      tx.Data,
// 		V:            tx.V,
// 		R:            tx.R,
// 		S:            tx.S,
// 		ChainID:      tx.ChainID,
// 		SysFlag:      tx.SystemFlag,
// 	}
// 	return rlp.Encode(w, enc)
// }

// // DecodeRLP implements rlp.Decoder
// func (tx *SystemTx) DecodeRLP(s *rlp.Stream) error {
// 	var dec systemTxData
// 	if err := s.Decode(&dec); err != nil {
// 		return err
// 	}
// 	tx.Nonce = dec.AccountNonce
// 	tx.GasTipCap = dec.Price
// 	tx.Gas = dec.GasLimit
// 	tx.To = dec.Recipient
// 	tx.Value = dec.Amount
// 	tx.Data = dec.Payload
// 	tx.V = dec.V
// 	tx.R = dec.R
// 	tx.S = dec.S
// 	tx.ChainID = dec.ChainID
// 	tx.SystemFlag = dec.SysFlag
// 	return nil
// }

// // encode implements txdata
// func (tx *SystemTx) encode(b *bytes.Buffer) error {
// 	return rlp.Encode(b, tx)
// }

// // decode implements txdata
// func (tx *SystemTx) decode(input []byte) error {
// 	return rlp.DecodeBytes(input, tx)
// }

// func NewSystemTx(chainID *big.Int, nonce uint64, gasTipCap, gasFeeCap *big.Int, gas uint64, to *common.Address, value *big.Int, data []byte, systemFlag uint64) *Transaction {
// 	return NewTx(&SystemTx{
// 		ChainID:    chainID,
// 		Nonce:      nonce,
// 		GasTipCap:  gasTipCap,
// 		GasFeeCap:  gasFeeCap,
// 		Gas:        gas,
// 		To:         to,
// 		Value:      value,
// 		Data:       data,
// 		SystemFlag: systemFlag,
// 	})
// }
