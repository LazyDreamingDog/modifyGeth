// Copyright 2016 The go-ethereum Authors
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

package state

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
)

// journalEntry is a modification entry in the state change journal that can be
// reverted on demand.
type journalEntry interface {
	// revert undoes the changes introduced by this journal entry.
	revert(*StateDB)

	// dirtied returns the Ethereum address modified by this journal entry.
	dirtied() *common.Address
}

// journal contains the list of state modifications applied since the last state
// commit. These are tracked to be able to be reverted in the case of an execution
// exception or request for reversal.
type journal struct {
	entries []journalEntry         // Current changes tracked by the journal
	dirties map[common.Address]int // Dirty accounts and the number of changes
}

// newJournal creates a new initialized journal.
func newJournal() *journal {
	return &journal{
		dirties: make(map[common.Address]int),
	}
}

// append inserts a new modification entry to the end of the change journal.
func (j *journal) append(entry journalEntry) {
	j.entries = append(j.entries, entry)
	if addr := entry.dirtied(); addr != nil {
		j.dirties[*addr]++
	}
}

// revert undoes a batch of journalled modifications along with any reverted
// dirty handling too.
func (j *journal) revert(statedb *StateDB, snapshot int) {
	for i := len(j.entries) - 1; i >= snapshot; i-- {
		// Undo the changes made by the operation
		j.entries[i].revert(statedb)

		// Drop any dirty tracking induced by the change
		if addr := j.entries[i].dirtied(); addr != nil {
			if j.dirties[*addr]--; j.dirties[*addr] == 0 {
				delete(j.dirties, *addr)
			}
		}
	}
	j.entries = j.entries[:snapshot]
}

// dirty explicitly sets an address to dirty, even if the change entries would
// otherwise suggest it as clean. This method is an ugly hack to handle the RIPEMD
// precompile consensus exception.
func (j *journal) dirty(addr common.Address) {
	j.dirties[addr]++
}

// length returns the current number of entries in the journal.
func (j *journal) length() int {
	return len(j.entries)
}

type (
	// Changes to the account trie.
	createObjectChange struct {
		account *common.Address
	}
	resetObjectChange struct {
		account      *common.Address
		prev         *stateObject
		prevdestruct bool
		prevAccount  []byte
		prevStorage  map[common.Hash][]byte

		prevAccountOriginExist bool
		prevAccountOrigin      []byte
		prevStorageOrigin      map[common.Hash][]byte
	}
	selfDestructChange struct {
		account     *common.Address
		prev        bool // whether account had already self-destructed
		prevbalance *uint256.Int
	}

	// Changes to individual accounts.
	balanceChange struct {
		account *common.Address
		prev    *uint256.Int
	}
	interestChange struct {
		account *common.Address
		prev    *uint256.Int
	}
	blockNumberChange struct {
		account *common.Address
		prev    *big.Int
	}
	nonceChange struct {
		account *common.Address
		prev    uint64
	}
	securityLevelChange struct {
		account *common.Address
		prev    uint64
	}
	storageChange struct {
		account       *common.Address
		key, prevalue common.Hash
	}
	codeChange struct {
		account            *common.Address
		prevcode, prevhash []byte
	}

	// Changes to other state values.
	refundChange struct {
		prev uint64
	}
	addLogChange struct {
		txhash common.Hash
	}
	addPreimageChange struct {
		hash common.Hash
	}
	touchChange struct {
		account *common.Address
	}
	// Changes to the access list
	accessListAddAccountChange struct {
		address *common.Address
	}
	accessListAddSlotChange struct {
		address *common.Address
		slot    *common.Hash
	}

	transientStorageChange struct {
		account       *common.Address
		key, prevalue common.Hash
	}

	PostQuanPubChange struct {
		account *common.Address
		prev    []byte
	}

	// Changes to the contract storage trie.
	contractCallCountChange struct {
		account  *common.Address
		prevalue *big.Int
	}

	totalNumberOfGasChange struct {
		account  *common.Address
		prevalue *uint256.Int
	}

	totalValueTxChange struct {
		account  *common.Address
		prevalue *uint256.Int
	}

	// changes to pledge info
	pledgeAmountChange struct {
		account *common.Address
		prev    uint64
	}
	pledgeYearChange struct {
		account *common.Address
		prev    int
	}
	startTimeChange struct {
		account *common.Address
		prev    uint64
	}
	interestRateChange struct {
		account *common.Address
		prev    int
	}
	currentInterestChange struct {
		account *common.Address
		prev    uint64
	}
	earnInterestChange struct {
		account *common.Address
		prev    uint64
	}
	annualFeeChange struct {
		account *common.Address
		prev    uint64
	}
	lastAnnualFeeTimeChange struct {
		account *common.Address
		prev    uint64
	}
	contractAddressChange struct {
		account *common.Address
		prev    common.Address
	}
	deployedAddressChange struct {
		account *common.Address
		prev    common.Address
	}
	investorAddressChange struct {
		account *common.Address
		prev    common.Address
	}
	beneficiaryAddressChange struct {
		account *common.Address
		prev    common.Address
	}

	stakeFlagChange struct {
		account *common.Address
		prev    bool
	}
)

func (ch totalValueTxChange) revert(s *StateDB) {
	s.getStateObject(*ch.account).SetTotalValueTx(ch.prevalue)
}

func (ch totalValueTxChange) dirtied() *common.Address {
	return ch.account
}

func (ch totalNumberOfGasChange) revert(s *StateDB) {
	s.getStateObject(*ch.account).SetTotalNumberOfGas(ch.prevalue)
}

func (ch totalNumberOfGasChange) dirtied() *common.Address {
	return ch.account
}

func (ch contractCallCountChange) revert(s *StateDB) {
	// 返回旧值（回滚）
	s.getStateObject(*ch.account).SetContractCallCount(ch.prevalue)
	// s.getStateObject(*ch.account).setBalance(ch.prev)
}

func (ch contractCallCountChange) dirtied() *common.Address {
	return ch.account
}

func (ch createObjectChange) revert(s *StateDB) {
	delete(s.stateObjects, *ch.account)
	delete(s.stateObjectsDirty, *ch.account)
}

func (ch createObjectChange) dirtied() *common.Address {
	return ch.account
}

func (ch resetObjectChange) revert(s *StateDB) {
	s.setStateObject(ch.prev)
	if !ch.prevdestruct {
		delete(s.stateObjectsDestruct, ch.prev.address)
	}
	if ch.prevAccount != nil {
		s.accounts[ch.prev.addrHash] = ch.prevAccount
	}
	if ch.prevStorage != nil {
		s.storages[ch.prev.addrHash] = ch.prevStorage
	}
	if ch.prevAccountOriginExist {
		s.accountsOrigin[ch.prev.address] = ch.prevAccountOrigin
	}
	if ch.prevStorageOrigin != nil {
		s.storagesOrigin[ch.prev.address] = ch.prevStorageOrigin
	}
}

func (ch resetObjectChange) dirtied() *common.Address {
	return ch.account
}

func (ch selfDestructChange) revert(s *StateDB) {
	obj := s.getStateObject(*ch.account)
	if obj != nil {
		obj.selfDestructed = ch.prev
		obj.setBalance(ch.prevbalance)
	}
}

func (ch selfDestructChange) dirtied() *common.Address {
	return ch.account
}

var ripemd = common.HexToAddress("0000000000000000000000000000000000000003")

func (ch touchChange) revert(s *StateDB) {
}

func (ch touchChange) dirtied() *common.Address {
	return ch.account
}

func (ch balanceChange) revert(s *StateDB) {
	s.getStateObject(*ch.account).setBalance(ch.prev)
}

func (ch balanceChange) dirtied() *common.Address {
	return ch.account
}

func (ch interestChange) revert(s *StateDB) {
	s.getStateObject(*ch.account).setInterest(ch.prev)
}

func (ch interestChange) dirtied() *common.Address {
	return ch.account
}

func (ch blockNumberChange) revert(s *StateDB) {
	s.getStateObject(*ch.account).SetLastNumber(ch.prev)
}

func (ch blockNumberChange) dirtied() *common.Address {
	return ch.account
}

func (ch nonceChange) revert(s *StateDB) {
	s.getStateObject(*ch.account).setNonce(ch.prev)
}

func (ch nonceChange) dirtied() *common.Address {
	return ch.account
}

func (ch securityLevelChange) revert(s *StateDB) {
	s.getStateObject(*ch.account).setSecurityLevel(ch.prev)
}

func (ch securityLevelChange) dirtied() *common.Address {
	return ch.account
}

func (ch codeChange) revert(s *StateDB) {
	s.getStateObject(*ch.account).setCode(common.BytesToHash(ch.prevhash), ch.prevcode)
}

func (ch codeChange) dirtied() *common.Address {
	return ch.account
}

func (ch storageChange) revert(s *StateDB) {
	s.getStateObject(*ch.account).setState(ch.key, ch.prevalue)
}

func (ch storageChange) dirtied() *common.Address {
	return ch.account
}

func (ch transientStorageChange) revert(s *StateDB) {
	s.setTransientState(*ch.account, ch.key, ch.prevalue)
}

func (ch transientStorageChange) dirtied() *common.Address {
	return nil
}

func (ch refundChange) revert(s *StateDB) {
	s.refund = ch.prev
}

func (ch refundChange) dirtied() *common.Address {
	return nil
}

func (ch addLogChange) revert(s *StateDB) {
	logs := s.logs[ch.txhash]
	if len(logs) == 1 {
		delete(s.logs, ch.txhash)
	} else {
		s.logs[ch.txhash] = logs[:len(logs)-1]
	}
	s.logSize--
}

func (ch addLogChange) dirtied() *common.Address {
	return nil
}

func (ch addPreimageChange) revert(s *StateDB) {
	delete(s.preimages, ch.hash)
}

func (ch addPreimageChange) dirtied() *common.Address {
	return nil
}

func (ch accessListAddAccountChange) revert(s *StateDB) {
	/*
		One important invariant here, is that whenever a (addr, slot) is added, if the
		addr is not already present, the add causes two journal entries:
		- one for the address,
		- one for the (address,slot)
		Therefore, when unrolling the change, we can always blindly delete the
		(addr) at this point, since no storage adds can remain when come upon
		a single (addr) change.
	*/
	s.accessList.DeleteAddress(*ch.address)
}

func (ch accessListAddAccountChange) dirtied() *common.Address {
	return nil
}

func (ch accessListAddSlotChange) revert(s *StateDB) {
	s.accessList.DeleteSlot(*ch.address, *ch.slot)
}

func (ch accessListAddSlotChange) dirtied() *common.Address {
	return nil
}

func (ch PostQuanPubChange) revert(s *StateDB) {
	s.getStateObject(*ch.account).setPostQuanPub(ch.prev)
}

func (ch PostQuanPubChange) dirtied() *common.Address {
	return ch.account
}

// 质押类
func (ch pledgeAmountChange) revert(s *StateDB) {
	s.getStateObject(*ch.account).SetPledgeAmount(ch.prev)
}

func (ch pledgeAmountChange) dirtied() *common.Address {
	return ch.account
}

func (ch pledgeYearChange) revert(s *StateDB) {
	s.getStateObject(*ch.account).SetPledgeYear(ch.prev)
}

func (ch pledgeYearChange) dirtied() *common.Address {
	return ch.account
}

func (ch startTimeChange) revert(s *StateDB) {
	s.getStateObject(*ch.account).SetStartTime(ch.prev)
}

func (ch startTimeChange) dirtied() *common.Address {
	return ch.account
}

func (ch interestRateChange) revert(s *StateDB) {
	s.getStateObject(*ch.account).SetInterestRate(ch.prev)
}

func (ch interestRateChange) dirtied() *common.Address {
	return ch.account
}

func (ch currentInterestChange) revert(s *StateDB) {
	s.getStateObject(*ch.account).SetCurrentInterest(ch.prev)
}

func (ch currentInterestChange) dirtied() *common.Address {
	return ch.account
}

func (ch earnInterestChange) revert(s *StateDB) {
	s.getStateObject(*ch.account).SetEarnInterest(ch.prev)
}

func (ch earnInterestChange) dirtied() *common.Address {
	return ch.account
}

func (ch annualFeeChange) revert(s *StateDB) {
	s.getStateObject(*ch.account).SetAnnualFee(ch.prev)
}

func (ch annualFeeChange) dirtied() *common.Address {
	return ch.account
}

func (ch lastAnnualFeeTimeChange) revert(s *StateDB) {
	s.getStateObject(*ch.account).SetLastAnnualFeeTime(ch.prev)
}

func (ch lastAnnualFeeTimeChange) dirtied() *common.Address {
	return ch.account
}

func (ch contractAddressChange) revert(s *StateDB) {
	s.getStateObject(*ch.account).SetContractAddress(ch.prev)
}

func (ch contractAddressChange) dirtied() *common.Address {
	return ch.account
}

func (ch deployedAddressChange) revert(s *StateDB) {
	s.getStateObject(*ch.account).SetDeployedAddress(ch.prev)
}

func (ch deployedAddressChange) dirtied() *common.Address {
	return ch.account
}

func (ch investorAddressChange) revert(s *StateDB) {
	s.getStateObject(*ch.account).SetInvestorAddress(ch.prev)
}

func (ch investorAddressChange) dirtied() *common.Address {
	return ch.account
}

func (ch beneficiaryAddressChange) revert(s *StateDB) {
	s.getStateObject(*ch.account).SetBeneficiaryAddress(ch.prev)
}

func (ch beneficiaryAddressChange) dirtied() *common.Address {
	return ch.account
}

func (ch stakeFlagChange) revert(s *StateDB) {
	s.getStateObject(*ch.account).SetStakeFlag(ch.prev)
}

func (ch stakeFlagChange) dirtied() *common.Address {
	return ch.account
}
