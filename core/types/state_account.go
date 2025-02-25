// Copyright 2021 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package types

import (
	"bytes"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/holiman/uint256"
)

//go:generate go run ../../rlp/rlpgen -type StateAccount -out gen_account_rlp.go

// StateAccount is the Ethereum consensus representation of accounts.
// These objects are stored in the main account trie.
type StateAccount struct {
	Nonce             uint64
	Balance           *uint256.Int
	Root              common.Hash // merkle root of the storage trie
	CodeHash          []byte
	SecurityLevel     uint64 // security level range 1-5. 0 means the account is locked
	Interest          *uint256.Int
	LastBlockNumber   *big.Int // The block number where the interest was last calculated
	LastPostQuanPub   []byte
	TotalNumberOfGas  *uint256.Int
	ContractCallCount *big.Int
	TotalValueTx      *uint256.Int

	// 质押信息
	PledgeAmount uint64 //质押金额
	PledgeYear   uint64 //质押年限
	StartTime    uint64 //开始时间(区块高度)
	InterestRate uint64 //利率

	CurrentInterest uint64 //当前利息
	EarnInterest    uint64 //收益利息（利息差）

	AnnualFee         uint64 //合约部署年费
	LastAnnualFeeTime uint64 //上一次收取年费的时间

	ContractAddress    common.Address //合约地址
	DeployedAddress    common.Address //部署人地址
	InvestorAddress    common.Address //投资人地址
	BeneficiaryAddress common.Address //受益人地址

	StakeFlag bool //质押标志
}

// NewEmptyStateAccount constructs an empty state account.
func NewEmptyStateAccount() *StateAccount {
	return &StateAccount{
		Balance:           new(uint256.Int),
		Root:              EmptyRootHash,
		CodeHash:          EmptyCodeHash.Bytes(),
		SecurityLevel:     1,
		Interest:          uint256.NewInt(0),
		LastBlockNumber:   big.NewInt(0),
		LastPostQuanPub:   EmptyCodeHash.Bytes(),
		ContractCallCount: big.NewInt(0),
		TotalNumberOfGas:  uint256.NewInt(0),
		TotalValueTx:      uint256.NewInt(0),

		PledgeAmount:       0,
		PledgeYear:         0,
		StartTime:          0,
		InterestRate:       0,
		CurrentInterest:    0,
		EarnInterest:       0,
		AnnualFee:          0,
		LastAnnualFeeTime:  0,
		ContractAddress:    common.Address{},
		DeployedAddress:    common.Address{},
		InvestorAddress:    common.Address{},
		BeneficiaryAddress: common.Address{},
		StakeFlag:          false,
	}
}

// Copy returns a deep-copied state account object.
func (acct *StateAccount) Copy() *StateAccount {
	var balance *uint256.Int
	if acct.Balance != nil {
		balance = new(uint256.Int).Set(acct.Balance)
	}
	return &StateAccount{
		Nonce:             acct.Nonce,
		Balance:           balance,
		Root:              acct.Root,
		CodeHash:          common.CopyBytes(acct.CodeHash),
		SecurityLevel:     acct.SecurityLevel,
		Interest:          acct.Interest,
		LastBlockNumber:   acct.LastBlockNumber,
		LastPostQuanPub:   common.CopyBytes(acct.LastPostQuanPub),
		ContractCallCount: acct.ContractCallCount,
		TotalNumberOfGas:  acct.TotalNumberOfGas,
		TotalValueTx:      acct.TotalValueTx,

		PledgeAmount:       acct.PledgeAmount,
		PledgeYear:         acct.PledgeYear,
		StartTime:          acct.StartTime,
		InterestRate:       acct.InterestRate,
		CurrentInterest:    acct.CurrentInterest,
		EarnInterest:       acct.EarnInterest,
		AnnualFee:          acct.AnnualFee,
		LastAnnualFeeTime:  acct.LastAnnualFeeTime,
		ContractAddress:    acct.ContractAddress,
		DeployedAddress:    acct.DeployedAddress,
		InvestorAddress:    acct.InvestorAddress,
		BeneficiaryAddress: acct.BeneficiaryAddress,
		StakeFlag:          acct.StakeFlag,
	}
}

// SlimAccount is a modified version of an Account, where the root is replaced
// with a byte slice. This format can be used to represent full-consensus format
// or slim format which replaces the empty root and code hash as nil byte slice.
type SlimAccount struct {
	Nonce             uint64
	Balance           *uint256.Int
	Root              []byte // Nil if root equals to types.EmptyRootHash
	CodeHash          []byte // Nil if hash equals to types.EmptyCodeHash
	SecurityLevel     uint64
	Interest          *uint256.Int
	LastBlockNumber   *big.Int
	LastPostQuanPub   []byte
	ContractCallCount *big.Int
	TotalNumberOfGas  *uint256.Int
	TotalValueTx      *uint256.Int

	// 质押信息
	PledgeAmount uint64 //质押金额
	PledgeYear   uint64 //质押年限
	StartTime    uint64 //开始时间(区块高度)
	InterestRate uint64 //利率

	CurrentInterest uint64 //当前利息
	EarnInterest    uint64 //收益利息（利息差）

	AnnualFee         uint64 //合约部署年费
	LastAnnualFeeTime uint64 //上一次收取年费的时间

	ContractAddress    common.Address //合约地址
	DeployedAddress    common.Address //部署人地址
	InvestorAddress    common.Address //投资人地址
	BeneficiaryAddress common.Address //受益人地址

	StakeFlag bool //质押标志
}

// SlimAccountRLP encodes the state account in 'slim RLP' format.
func SlimAccountRLP(account StateAccount) []byte {
	slim := SlimAccount{
		Nonce:             account.Nonce,
		Balance:           account.Balance,
		SecurityLevel:     account.SecurityLevel,
		Interest:          account.Interest,
		LastBlockNumber:   account.LastBlockNumber,
		ContractCallCount: account.ContractCallCount,
		TotalNumberOfGas:  account.TotalNumberOfGas,
		TotalValueTx:      account.TotalValueTx,

		PledgeAmount:       account.PledgeAmount,
		PledgeYear:         account.PledgeYear,
		StartTime:          account.StartTime,
		InterestRate:       account.InterestRate,
		CurrentInterest:    account.CurrentInterest,
		EarnInterest:       account.EarnInterest,
		AnnualFee:          account.AnnualFee,
		LastAnnualFeeTime:  account.LastAnnualFeeTime,
		ContractAddress:    account.ContractAddress,
		DeployedAddress:    account.DeployedAddress,
		InvestorAddress:    account.InvestorAddress,
		BeneficiaryAddress: account.BeneficiaryAddress,
		StakeFlag:          account.StakeFlag,
	}
	if account.Root != EmptyRootHash {
		slim.Root = account.Root[:]
	}
	if !bytes.Equal(account.CodeHash, EmptyCodeHash[:]) {
		slim.CodeHash = account.CodeHash
	}
	if !bytes.Equal(account.LastPostQuanPub, EmptyCodeHash[:]) {
		slim.LastPostQuanPub = account.LastPostQuanPub
	}
	data, err := rlp.EncodeToBytes(slim)
	if err != nil {
		panic(err)
	}
	return data
}

// FullAccount decodes the data on the 'slim RLP' format and returns
// the consensus format account.
func FullAccount(data []byte) (*StateAccount, error) {
	var slim SlimAccount
	if err := rlp.DecodeBytes(data, &slim); err != nil {
		return nil, err
	}
	var account StateAccount
	account.Nonce, account.Balance = slim.Nonce, slim.Balance
	account.SecurityLevel = slim.SecurityLevel
	account.LastBlockNumber = slim.LastBlockNumber
	account.Interest = slim.Interest
	account.ContractCallCount = slim.ContractCallCount
	account.TotalNumberOfGas = slim.TotalNumberOfGas
	account.TotalValueTx = slim.TotalValueTx

	// 质押信息
	account.PledgeAmount = slim.PledgeAmount
	account.PledgeYear = slim.PledgeYear
	account.StartTime = slim.StartTime
	account.InterestRate = slim.InterestRate
	account.CurrentInterest = slim.CurrentInterest
	account.EarnInterest = slim.EarnInterest
	account.AnnualFee = slim.AnnualFee
	account.LastAnnualFeeTime = slim.LastAnnualFeeTime
	account.ContractAddress = slim.ContractAddress
	account.DeployedAddress = slim.DeployedAddress
	account.InvestorAddress = slim.InvestorAddress
	account.BeneficiaryAddress = slim.BeneficiaryAddress
	account.StakeFlag = slim.StakeFlag

	// Interpret the storage root and code hash in slim format.
	if len(slim.Root) == 0 {
		account.Root = EmptyRootHash
	} else {
		account.Root = common.BytesToHash(slim.Root)
	}
	if len(slim.CodeHash) == 0 {
		account.CodeHash = EmptyCodeHash[:]
	} else {
		account.CodeHash = slim.CodeHash
	}
	return &account, nil
}

// FullAccountRLP converts data on the 'slim RLP' format into the full RLP-format.
func FullAccountRLP(data []byte) ([]byte, error) {
	account, err := FullAccount(data)
	if err != nil {
		return nil, err
	}
	return rlp.EncodeToBytes(account)
}
