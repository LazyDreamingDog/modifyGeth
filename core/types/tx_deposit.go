package types

import (
	"bytes"
	"io"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

// DynamicCryptoTx represents a new transaction type with additional fields: CryptoType and SignatureData.
type DepositTx struct {
	ChainID    *big.Int
	Nonce      uint64
	GasTipCap  *big.Int // a.k.a. maxPriorityFeePerGas
	GasFeeCap  *big.Int // a.k.a. maxFeePerGas
	Gas        uint64
	To         *common.Address `rlp:"nil"` // nil means contract creation
	Value      *big.Int
	Data       []byte
	AccessList AccessList

	// 新增字段
	DeployerAddress    *common.Address `rlp:"nil"`
	InvestorAddress    *common.Address `rlp:"nil"`
	BeneficiaryAddress *common.Address `rlp:"nil"`
	StakedAmount       *big.Int
	StakedTime         uint64 // 单位时间是半年

	// Signature values
	V *big.Int `json:"v" gencodec:"required"`
	R *big.Int `json:"r" gencodec:"required"`
	S *big.Int `json:"s" gencodec:"required"`
}

// copy creates a deep copy of the transaction data and initializes all fields.
func (tx *DepositTx) copy() TxData {
	cpy := &DepositTx{
		Nonce:              tx.Nonce,
		To:                 copyAddressPtr(tx.To),
		Data:               common.CopyBytes(tx.Data),
		Gas:                tx.Gas,
		DeployerAddress:    copyAddressPtr(tx.DeployerAddress),
		InvestorAddress:    copyAddressPtr(tx.InvestorAddress),
		BeneficiaryAddress: copyAddressPtr(tx.BeneficiaryAddress),
		StakedTime:         tx.StakedTime,
		AccessList:         make(AccessList, len(tx.AccessList)),
		Value:              new(big.Int),
		ChainID:            new(big.Int),
		GasTipCap:          new(big.Int),
		GasFeeCap:          new(big.Int),
		StakedAmount:       new(big.Int),
		V:                  new(big.Int),
		R:                  new(big.Int),
		S:                  new(big.Int),
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
	if tx.StakedAmount != nil {
		cpy.StakedAmount.Set(tx.StakedAmount)
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
func (tx *DepositTx) txType() byte           { return DynamicCryptoTxType }
func (tx *DepositTx) chainID() *big.Int      { return tx.ChainID }
func (tx *DepositTx) accessList() AccessList { return tx.AccessList }
func (tx *DepositTx) data() []byte           { return tx.Data }
func (tx *DepositTx) gas() uint64            { return tx.Gas }
func (tx *DepositTx) gasFeeCap() *big.Int    { return tx.GasFeeCap }
func (tx *DepositTx) gasTipCap() *big.Int    { return tx.GasTipCap }
func (tx *DepositTx) gasPrice() *big.Int     { return tx.GasFeeCap }
func (tx *DepositTx) value() *big.Int        { return tx.Value }
func (tx *DepositTx) nonce() uint64          { return tx.Nonce }
func (tx *DepositTx) to() *common.Address    { return tx.To }
func (tx *DepositTx) stakedAmount() *big.Int { return tx.StakedAmount }
func (tx *DepositTx) deployerAddress() *common.Address {
	return tx.DeployerAddress
}
func (tx *DepositTx) investorAddress() *common.Address {
	return tx.InvestorAddress
}
func (tx *DepositTx) beneficiaryAddress() *common.Address {
	return tx.BeneficiaryAddress
}
func (tx *DepositTx) stakedTime() uint64 { return tx.StakedTime }

func (tx *DepositTx) effectiveGasPrice(dst *big.Int, baseFee *big.Int) *big.Int {
	if baseFee == nil {
		return dst.Set(tx.GasFeeCap)
	}
	tip := dst.Sub(tx.GasFeeCap, baseFee)
	if tip.Cmp(tx.GasTipCap) > 0 {
		tip.Set(tx.GasTipCap)
	}
	return tip.Add(tip, baseFee)
}

func (tx *DepositTx) rawSignatureValues() (v, r, s *big.Int) {
	return tx.V, tx.R, tx.S
}

func (tx *DepositTx) setSignatureValues(chainID, v, r, s *big.Int) {
	tx.ChainID, tx.V, tx.R, tx.S = chainID, v, r, s
}

// dynamicCryptoTxData 用于序列化/反序列化的数据结构
type depositTxData struct {
	AccountNonce       uint64
	Price              *big.Int
	GasLimit           uint64
	Recipient          *common.Address
	Amount             *big.Int
	Payload            []byte
	V                  *big.Int
	R                  *big.Int
	S                  *big.Int
	ChainID            *big.Int
	DeployerAddress    *common.Address
	InvestorAddress    *common.Address
	BeneficiaryAddress *common.Address
	StakedAmount       *big.Int
	StakedTime         uint64
}

// EncodeRLP implements rlp.Encoder
func (tx *DepositTx) EncodeRLP(w io.Writer) error {
	enc := &depositTxData{
		AccountNonce:       tx.Nonce,
		Price:              tx.GasTipCap,
		GasLimit:           tx.Gas,
		Recipient:          tx.To,
		Amount:             tx.Value,
		Payload:            tx.Data,
		V:                  tx.V,
		R:                  tx.R,
		S:                  tx.S,
		ChainID:            tx.ChainID,
		DeployerAddress:    tx.DeployerAddress,
		InvestorAddress:    tx.InvestorAddress,
		BeneficiaryAddress: tx.BeneficiaryAddress,
		StakedAmount:       tx.StakedAmount,
		StakedTime:         tx.StakedTime,
	}
	return rlp.Encode(w, enc)
}

// DecodeRLP implements rlp.Decoder
func (tx *DepositTx) DecodeRLP(s *rlp.Stream) error {
	var dec depositTxData
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
	tx.DeployerAddress = dec.DeployerAddress
	tx.InvestorAddress = dec.InvestorAddress
	tx.BeneficiaryAddress = dec.BeneficiaryAddress
	tx.StakedAmount = dec.StakedAmount
	tx.StakedTime = dec.StakedTime
	return nil
}

// 补全RLP编码
func (tx *DepositTx) encode(b *bytes.Buffer) error {
	return rlp.Encode(b, tx)
}

func (tx *DepositTx) decode(input []byte) error {
	return rlp.DecodeBytes(input, tx)
}
