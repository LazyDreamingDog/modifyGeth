package types

import (
	"bytes"
	"encoding/json"
	"io"
	"math/big"
	"reflect"
	"sort"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

// DynamicCryptoTx represents a new transaction type with additional fields: CryptoType and SignatureData.
type DynamicCryptoTx struct {
	ChainID    *big.Int
	Nonce      uint64
	GasTipCap  *big.Int // a.k.a. maxPriorityFeePerGas
	GasFeeCap  *big.Int // a.k.a. maxFeePerGas
	Gas        uint64
	To         *common.Address `rlp:"nil"` // nil means contract creation
	Value      *big.Int
	Data       []byte
	AccessList AccessList

	// Post-quantum relevant attributes
	PostAddress   *common.Address
	CryptoType    []byte // New field to represent the type of cryptocurrency
	SignatureData []byte // New field for signature data
	PublicKey     []byte // New field for public key

	// Signature values
	V *big.Int `json:"v" gencodec:"required"`
	R *big.Int `json:"r" gencodec:"required"`
	S *big.Int `json:"s" gencodec:"required"`
}

// copy creates a deep copy of the transaction data and initializes all fields.
func (tx *DynamicCryptoTx) copy() TxData {
	cpy := &DynamicCryptoTx{
		Nonce:         tx.Nonce,
		To:            copyAddressPtr(tx.To),
		Data:          common.CopyBytes(tx.Data),
		Gas:           tx.Gas,
		CryptoType:    common.CopyBytes(tx.CryptoType),
		SignatureData: common.CopyBytes(tx.SignatureData),
		PostAddress:   tx.PostAddress,
		PublicKey:     common.CopyBytes(tx.PublicKey),
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
func (tx *DynamicCryptoTx) txType() byte                 { return DynamicCryptoTxType }
func (tx *DynamicCryptoTx) chainID() *big.Int            { return tx.ChainID }
func (tx *DynamicCryptoTx) accessList() AccessList       { return tx.AccessList }
func (tx *DynamicCryptoTx) data() []byte                 { return tx.Data }
func (tx *DynamicCryptoTx) gas() uint64                  { return tx.Gas }
func (tx *DynamicCryptoTx) gasFeeCap() *big.Int          { return tx.GasFeeCap }
func (tx *DynamicCryptoTx) gasTipCap() *big.Int          { return tx.GasTipCap }
func (tx *DynamicCryptoTx) gasPrice() *big.Int           { return tx.GasFeeCap }
func (tx *DynamicCryptoTx) value() *big.Int              { return tx.Value }
func (tx *DynamicCryptoTx) nonce() uint64                { return tx.Nonce }
func (tx *DynamicCryptoTx) to() *common.Address          { return tx.To }
func (tx *DynamicCryptoTx) cryptoType() []byte           { return tx.CryptoType }
func (tx *DynamicCryptoTx) signatureData() []byte        { return tx.SignatureData }
func (tx *DynamicCryptoTx) postAddress() *common.Address { return tx.PostAddress }
func (tx *DynamicCryptoTx) publicKey() []byte            { return tx.PublicKey }

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
	AccountNonce  uint64
	Price         *big.Int
	GasFeeCap     *big.Int
	GasLimit      uint64
	Recipient     *common.Address
	Amount        *big.Int
	Payload       []byte
	V             *big.Int
	R             *big.Int
	S             *big.Int
	ChainID       *big.Int
	CryptoType    []byte
	SignatureData []byte
	PostAddress   *common.Address
	PublicKey     []byte
}

// EncodeRLP implements rlp.Encoder
func (tx *DynamicCryptoTx) EncodeRLP(w io.Writer) error {
	enc := &dynamicCryptoTxData{
		AccountNonce:  tx.Nonce,
		Price:         tx.GasTipCap,
		GasFeeCap:     tx.GasFeeCap,
		GasLimit:      tx.Gas,
		Recipient:     tx.To,
		Amount:        tx.Value,
		Payload:       tx.Data,
		V:             tx.V,
		R:             tx.R,
		S:             tx.S,
		ChainID:       tx.ChainID,
		CryptoType:    tx.CryptoType,
		SignatureData: tx.SignatureData,
		PostAddress:   tx.PostAddress,
		PublicKey:     tx.PublicKey,
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
	tx.GasFeeCap = dec.GasFeeCap
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
	tx.PostAddress = dec.PostAddress
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

// Hash used by post-quantum signature.
func (tx *DynamicCryptoTx) jsonExcludeSign() ([]byte, error) {

	type dynamciData struct {
		ChainID   *big.Int        `json:"ChainID"`
		Nonce     uint64          `json:"Nonce"`
		GasTipCap *big.Int        `json:"GasTipCap"`
		GasFeeCap *big.Int        `json:"GasFeeCap"`
		Gas       uint64          `json:"Gas"`
		To        *common.Address `json:"To"`
		Value     *big.Int        `json:"Value"`
		Data      []byte          `json:"Data"`
	}

	txInfo := dynamciData{
		ChainID:   tx.ChainID,
		Nonce:     tx.Nonce,
		GasTipCap: tx.GasTipCap,
		GasFeeCap: tx.GasFeeCap,
		Gas:       tx.Gas,
		To:        tx.To,
		Value:     tx.Value,
		Data:      tx.Data,
	}
	// Serialize based attribute's name
	return sortedJSONMarshal(txInfo)
}

// Manually serialize the structure by JSON, sorting by field name
func sortedJSONMarshal(v interface{}) ([]byte, error) {
	value := reflect.ValueOf(v)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	if value.Kind() != reflect.Struct {
		return json.Marshal(v)
	}
	// Get struct's type
	typ := value.Type()
	// Store attributes contained in struct
	fieldNames := make([]string, 0, typ.NumField())
	fieldValues := make(map[string]interface{})
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		// Get JSON tag in struct: if not given will be attribute name
		tag := field.Tag.Get("json")
		if tag == "" {
			tag = field.Name
		}
		// Get attribute's value
		fieldValues[tag] = value.Field(i).Interface()
		fieldNames = append(fieldNames, tag)
	}
	// Sort
	sort.Strings(fieldNames)
	// Simulate the sorted structure
	result := make(map[string]interface{})
	for _, name := range fieldNames {
		result[name] = fieldValues[name]
	}
	// Serialize
	return json.Marshal(result)
}

// Helper function that finds the position of a character in a string
func idxOr(s string, c byte, or int) int {
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			return i
		}
	}
	return or
}
