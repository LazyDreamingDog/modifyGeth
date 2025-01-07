// Copyright 2014 The go-ethereum Authors
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

package core

import (
	"fmt"

	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	cmath "github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto/kzg4844"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/voucher"
	"github.com/holiman/uint256"
)

// ExecutionResult includes all output after executing given evm
// message no matter the execution itself is successful or not.
type ExecutionResult struct {
	UsedGas     uint64       // Total used gas, not including the refunded gas
	RefundedGas uint64       // Total gas refunded after execution
	Incentive   *uint256.Int //Total incentive of miner
	Err         error        // Any error encountered during the execution(listed in core/vm/errors.go)
	ReturnData  []byte       // Returned data from evm(function result or data supplied with revert opcode)
}

// Unwrap returns the internal evm error which allows us for further
// analysis outside.
func (result *ExecutionResult) Unwrap() error {
	return result.Err
}

// Failed returns the indicator whether the execution is successful or not
func (result *ExecutionResult) Failed() bool { return result.Err != nil }

// Return is a helper function to help caller distinguish between revert reason
// and function return. Return returns the data after execution if no error occurs.
func (result *ExecutionResult) Return() []byte {
	if result.Err != nil {
		return nil
	}
	return common.CopyBytes(result.ReturnData)
}

// Revert returns the concrete revert reason if the execution is aborted by `REVERT`
// opcode. Note the reason can be nil if no data supplied with revert opcode.
func (result *ExecutionResult) Revert() []byte {
	if result.Err != vm.ErrExecutionReverted {
		return nil
	}
	return common.CopyBytes(result.ReturnData)
}

// IntrinsicGas computes the 'intrinsic gas' for a message with the given data.
func IntrinsicGas(data []byte, accessList types.AccessList, isContractCreation bool, isHomestead, isEIP2028 bool, isEIP3860 bool) (uint64, error) {
	// Set the starting gas for the raw transaction
	var gas uint64
	if isContractCreation && isHomestead {
		gas = params.TxGasContractCreation
	} else {
		gas = params.TxGas
	}
	dataLen := uint64(len(data))
	// Bump the required gas by the amount of transactional data
	if dataLen > 0 {
		// Zero and non-zero bytes are priced differently
		var nz uint64
		for _, byt := range data {
			if byt != 0 {
				nz++
			}
		}
		// Make sure we don't exceed uint64 for all data combinations
		nonZeroGas := params.TxDataNonZeroGasFrontier
		if isEIP2028 {
			nonZeroGas = params.TxDataNonZeroGasEIP2028
		}
		if (math.MaxUint64-gas)/nonZeroGas < nz {
			return 0, ErrGasUintOverflow
		}
		gas += nz * nonZeroGas

		z := dataLen - nz
		if (math.MaxUint64-gas)/params.TxDataZeroGas < z {
			return 0, ErrGasUintOverflow
		}
		gas += z * params.TxDataZeroGas

		if isContractCreation && isEIP3860 {
			lenWords := toWordSize(dataLen)
			if (math.MaxUint64-gas)/params.InitCodeWordGas < lenWords {
				return 0, ErrGasUintOverflow
			}
			gas += lenWords * params.InitCodeWordGas
		}
	}
	if accessList != nil {
		gas += uint64(len(accessList)) * params.TxAccessListAddressGas
		gas += uint64(accessList.StorageKeys()) * params.TxAccessListStorageKeyGas
	}
	return gas, nil
}

// toWordSize returns the ceiled word size required for init code payment calculation.
func toWordSize(size uint64) uint64 {
	if size > math.MaxUint64-31 {
		return math.MaxUint64/32 + 1
	}

	return (size + 31) / 32
}

// A Message contains the data derived from a single transaction that is relevant to state
// processing.
type Message struct {
	To            *common.Address
	From          common.Address
	Nonce         uint64
	Value         *big.Int
	GasLimit      uint64
	GasPrice      *big.Int
	GasFeeCap     *big.Int
	GasTipCap     *big.Int
	Data          []byte
	AccessList    types.AccessList
	BlobGasFeeCap *big.Int
	BlobHashes    []common.Hash

	// Attribute for MultiVoucher
	tokenName   string
	FeeCurrency *common.Address

	// When SkipAccountChecks is true, the message nonce is not checked against the
	// account nonce in state. It also disables checking that the sender is an EOA.
	// This field will be set to true for operations like RPC eth_call.
	SkipAccountChecks bool

	HashNonce      *big.Int
	CryptoType     []byte
	SignatureData  []byte
	PublicKey      []byte
	PublicKeyIndex uint64
	SystemFlag     uint64

	// if isPow is true, the message is a PoW transaction
	// TODO: check whether it can use pow as gas.
	IsPow bool
}

// Parse voucher info from Tx.Data, delete flag header when finished.
// Extra information of Tx.Data contains: name([10]byte)
func (m *Message) ParseVoucher() {
	// log.Info("Begin to parse voucher")
	// fmt.Printf("len(m.Data): %v. m.Data %v\n", len(m.Data), m.Data)
	// Data's legitimacy check
	if len(m.Data) >= 23 && m.Data[0] == 0x0A && m.Data[1] == 0x0D && m.Data[2] == 0x03 {
		// Initial and assign pointer
		m.FeeCurrency = new(common.Address)
		*m.FeeCurrency = voucher.VoucherAddress

		// Parse token name
		i := 3
		for ; i < 23; i++ {
			if m.Data[i] == 0 {
				break
			}
		}
		m.tokenName = string(m.Data[3:i])
		// Delete prefix
		m.Data = m.Data[23:]
		log.Info("Tx use %s to pay Gas,contract address is %s!\n", m.tokenName, m.FeeCurrency)
	} else {
		log.Info("Tx is not use voucher to pay gas")
	}
}

// TransactionToMessage converts a transaction into a Message.
func TransactionToMessage(tx *types.Transaction, s types.Signer, baseFee *big.Int) (*Message, error) {
	if tx.Type() == types.SystemTxType {
		systemMsg := &Message{
			GasLimit:          tx.Gas(),
			GasPrice:          new(big.Int).Set(tx.GasPrice()),
			GasFeeCap:         new(big.Int).Set(tx.GasFeeCap()),
			GasTipCap:         new(big.Int).Set(tx.GasTipCap()),
			From:              common.HexToAddress("0x0000000000000000000000000000000000000000"),
			To:                tx.To(),
			Value:             tx.Value(),
			Data:              tx.Data(),
			SystemFlag:        tx.SystemFlag(),
			SkipAccountChecks: true,
		}
		return systemMsg, nil
	}

	msg := &Message{
		Nonce:             tx.Nonce(),
		GasLimit:          tx.Gas(),
		GasPrice:          new(big.Int).Set(tx.GasPrice()),
		GasFeeCap:         new(big.Int).Set(tx.GasFeeCap()),
		GasTipCap:         new(big.Int).Set(tx.GasTipCap()),
		To:                tx.To(),
		Value:             tx.Value(),
		Data:              tx.Data(),
		AccessList:        tx.AccessList(),
		SkipAccountChecks: false,
		BlobHashes:        tx.BlobHashes(),
		BlobGasFeeCap:     tx.BlobGasFeeCap(),
	}

	msg.ParseVoucher()

	// If baseFee provided, set gasPrice to effectiveGasPrice.
	if baseFee != nil {
		msg.GasPrice = cmath.BigMin(msg.GasPrice.Add(msg.GasTipCap, baseFee), msg.GasFeeCap)
	}
	var err error
	msg.From, err = types.Sender(s, tx)
	// Set IsPow flag if the transaction is a PoW transaction
	msg.IsPow = tx.Type() == types.PowTxType
	fmt.Println("msg.IsPow=", msg.IsPow)
	return msg, err
}

// ApplyMessage computes the new state by applying the given message
// against the old state within the environment.
//
// ApplyMessage returns the bytes returned by any EVM execution (if it took place),
// the gas used (which includes gas refunds) and an error if it failed. An error always
// indicates a core error meaning that the message would always fail for that particular
// state and would never be accepted within a block.
func ApplyMessage(evm *vm.EVM, msg *Message, gp *GasPool) (*ExecutionResult, error) {
	return NewStateTransition(evm, msg, gp).TransitionDb()
}

// StateTransition represents a state transition.
//
// == The State Transitioning Model
//
// A state transition is a change made when a transaction is applied to the current world
// state. The state transitioning model does all the necessary work to work out a valid new
// state root.
//
//  1. Nonce handling
//  2. Pre pay gas
//  3. Create a new state object if the recipient is nil
//  4. Value transfer
//
// == If contract creation ==
//
//	4a. Attempt to run transaction data
//	4b. If valid, use result as code for the new state object
//
// == end ==
//
//  5. Run Script section
//  6. Derive new state root
type StateTransition struct {
	gp           *GasPool
	msg          *Message
	gasRemaining uint64
	initialGas   uint64
	state        vm.StateDB
	evm          *vm.EVM
}

// NewStateTransition initialises and returns a new state transition object.
func NewStateTransition(evm *vm.EVM, msg *Message, gp *GasPool) *StateTransition {
	return &StateTransition{
		gp:    gp,
		evm:   evm,
		msg:   msg,
		state: evm.StateDB,
	}
}

// to returns the recipient of the message.
func (st *StateTransition) to() common.Address {
	if st.msg == nil || st.msg.To == nil /* contract creation */ {
		return common.Address{}
	}
	return *st.msg.To
}

func (st *StateTransition) buyGas() error {
	// For PowTx, we don't buy gas but convert to powgas directly
	if st.msg.IsPow {
		// Set gas remaining and initial gas to powgas
		st.gasRemaining += st.evm.Context.PowGas
		st.initialGas = st.evm.Context.PowGas
		return nil
	}
	mgval := new(big.Int).SetUint64(st.msg.GasLimit)
	mgval = mgval.Mul(mgval, st.msg.GasPrice)
	balanceCheck := new(big.Int).Set(mgval)
	if st.msg.GasFeeCap != nil {
		balanceCheck.SetUint64(st.msg.GasLimit)
		balanceCheck = balanceCheck.Mul(balanceCheck, st.msg.GasFeeCap)
		balanceCheck.Add(balanceCheck, st.msg.Value)
	}
	if st.evm.ChainConfig().IsCancun(st.evm.Context.BlockNumber, st.evm.Context.Time) {
		if blobGas := st.blobGasUsed(); blobGas > 0 {
			// Check that the user has enough funds to cover blobGasUsed * tx.BlobGasFeeCap
			blobBalanceCheck := new(big.Int).SetUint64(blobGas)
			blobBalanceCheck.Mul(blobBalanceCheck, st.msg.BlobGasFeeCap)
			balanceCheck.Add(balanceCheck, blobBalanceCheck)
			// Pay for blobGasUsed * actual blob fee
			blobFee := new(big.Int).SetUint64(blobGas)
			blobFee.Mul(blobFee, st.evm.Context.BlobBaseFee)
			mgval.Add(mgval, blobFee)
		}
	}
	balanceCheckU256, overflow := uint256.FromBig(balanceCheck)
	if overflow {
		return fmt.Errorf("%w: address %v required balance exceeds 256 bits", ErrInsufficientFunds, st.msg.From.Hex())
	}

	BalanceOfGas := uint64(0)
	// Check account balance enough to pay gas
	if st.msg.FeeCurrency != nil {
		// Using voucher to buy gas
		balance := big.NewInt(0)
		BalanceOfGas, _ = voucher.BalanceOf.Execute(st.evm, &balance, &st.msg.From, uint256.NewInt(0), st.msg.tokenName, st.msg.From)
		// Legitimacy check
		if balance.Cmp(balanceCheck) < 0 {
			fmt.Println("Balance of account is insufficient!")
			return fmt.Errorf("%w: address %v have %v want %v", ErrInsufficientFunds, st.msg, balance, balanceCheck)
		} else {
			fmt.Printf("Call BalanceOf use %v unit of gas\n", BalanceOfGas)
		}
	} else {
		// At first, use interest to pay gas.
		insufficient := st.state.UseInterest(st.msg.From, balanceCheckU256)
		if insufficient != nil {
			// Only using native token to pay insufficient part
			fmt.Println("Using interest to pay part of gas. The remaining gas to be paid is:", insufficient)
			if have, want := st.state.GetBalance(st.msg.From), insufficient; have.Cmp(want) < 0 {
				return fmt.Errorf("%w: address %v have %v want %v", ErrInsufficientFunds, st.msg.From.Hex(), have, want)
			}
		} else {
			fmt.Printf("Using interest to pay all %v units of gas\n", balanceCheckU256)
		}
	}

	if err := st.gp.SubGas(st.msg.GasLimit); err != nil {
		return err
	}
	// If use voucher to pay,need pay the gas of call BalanceOf
	st.gasRemaining += (st.msg.GasLimit - BalanceOfGas)

	st.initialGas = st.msg.GasLimit
	mgvalU256, _ := uint256.FromBig(mgval)
	// If ETH is used to pay for Gas fees, the from account needs to be deducted here.
	// If the payment is made with a non-original token, no payment is made here
	// But rather, after tx is executed, the remain is calculated for offset and payment is made.
	if st.msg.FeeCurrency == nil {
		st.state.SubBalance(st.msg.From, mgvalU256)
	}
	return nil
}

func (st *StateTransition) preCheck() error {
	// Only check transactions that are not fake
	msg := st.msg
	// check address is locked?(it means the account sercurity level is 0)
	// ! Account lock
	fromSL := st.state.GetSecurityLevel(msg.From)
	fmt.Println("fromSL=", fromSL)
	if fromSL == 0 {
		return fmt.Errorf("%w: address %v, account is locked", ErrAccountLocked, msg.From.Hex())
	}
	if msg.To != nil {
		toSL := st.state.GetSecurityLevel(*msg.To)
		fmt.Println("toSL=", toSL)
		if toSL == 0 {
			return fmt.Errorf("%w: address %v, account is locked", ErrAccountLocked, msg.To.Hex())
		}
	}

	// System flag is 1, skip all the checks
	if msg.SystemFlag == 1 {
		return nil
	}

	if !msg.SkipAccountChecks {
		// Make sure this transaction's nonce is correct.
		stNonce := st.state.GetNonce(msg.From)
		if msgNonce := msg.Nonce; stNonce < msgNonce {
			return fmt.Errorf("%w: address %v, tx: %d state: %d", ErrNonceTooHigh,
				msg.From.Hex(), msgNonce, stNonce)
		} else if stNonce > msgNonce {
			return fmt.Errorf("%w: address %v, tx: %d state: %d", ErrNonceTooLow,
				msg.From.Hex(), msgNonce, stNonce)
		} else if stNonce+1 < stNonce {
			return fmt.Errorf("%w: address %v, nonce: %d", ErrNonceMax,
				msg.From.Hex(), stNonce)
		}
		// Make sure the sender is an EOA
		codeHash := st.state.GetCodeHash(msg.From)
		if codeHash != (common.Hash{}) && codeHash != types.EmptyCodeHash {
			return fmt.Errorf("%w: address %v, codehash: %s", ErrSenderNoEOA,
				msg.From.Hex(), codeHash)
		}
	}
	// Make sure that transaction 	 is greater than the baseFee (post london)
	if st.evm.ChainConfig().IsLondon(st.evm.Context.BlockNumber) {
		// Skip the checks if gas fields are zero and baseFee was explicitly disabled (eth_call)
		skipCheck := st.evm.Config.NoBaseFee && msg.GasFeeCap.BitLen() == 0 && msg.GasTipCap.BitLen() == 0
		if !skipCheck {
			if l := msg.GasFeeCap.BitLen(); l > 256 {
				return fmt.Errorf("%w: address %v, maxFeePerGas bit length: %d", ErrFeeCapVeryHigh,
					msg.From.Hex(), l)
			}
			if l := msg.GasTipCap.BitLen(); l > 256 {
				return fmt.Errorf("%w: address %v, maxPriorityFeePerGas bit length: %d", ErrTipVeryHigh,
					msg.From.Hex(), l)
			}
			if msg.GasFeeCap.Cmp(msg.GasTipCap) < 0 {
				return fmt.Errorf("%w: address %v, maxPriorityFeePerGas: %s, maxFeePerGas: %s", ErrTipAboveFeeCap,
					msg.From.Hex(), msg.GasTipCap, msg.GasFeeCap)
			}
			// This will panic if baseFee is nil, but basefee presence is verified
			// as part of header validation.
			if msg.GasFeeCap.Cmp(st.evm.Context.BaseFee) < 0 {
				return fmt.Errorf("%w: address %v, maxFeePerGas: %s, baseFee: %s", ErrFeeCapTooLow,
					msg.From.Hex(), msg.GasFeeCap, st.evm.Context.BaseFee)
			}
		}
	}
	// Check the blob version validity
	if msg.BlobHashes != nil {
		// The to field of a blob tx type is mandatory, and a `BlobTx` transaction internally
		// has it as a non-nillable value, so any msg derived from blob transaction has it non-nil.
		// However, messages created through RPC (eth_call) don't have this restriction.
		if msg.To == nil {
			return ErrBlobTxCreate
		}
		if len(msg.BlobHashes) == 0 {
			return ErrMissingBlobHashes
		}
		for i, hash := range msg.BlobHashes {
			if !kzg4844.IsValidVersionedHash(hash[:]) {
				return fmt.Errorf("blob %d has invalid hash version", i)
			}
		}
	}
	// Check that the user is paying at least the current blob fee
	if st.evm.ChainConfig().IsCancun(st.evm.Context.BlockNumber, st.evm.Context.Time) {
		if st.blobGasUsed() > 0 {
			// Skip the checks if gas fields are zero and blobBaseFee was explicitly disabled (eth_call)
			skipCheck := st.evm.Config.NoBaseFee && msg.BlobGasFeeCap.BitLen() == 0
			if !skipCheck {
				// This will panic if blobBaseFee is nil, but blobBaseFee presence
				// is verified as part of header validation.
				if msg.BlobGasFeeCap.Cmp(st.evm.Context.BlobBaseFee) < 0 {
					return fmt.Errorf("%w: address %v blobGasFeeCap: %v, blobBaseFee: %v", ErrBlobFeeCapTooLow,
						msg.From.Hex(), msg.BlobGasFeeCap, st.evm.Context.BlobBaseFee)
				}
			}
		}
	}
	return st.buyGas()
}

// TransitionDb will transition the state by applying the current message and
// returning the evm execution result with following fields.
//
//   - used gas: total gas used (including gas being refunded)
//   - returndata: the returned data from evm
//   - concrete execution error: various EVM errors which abort the execution, e.g.
//     ErrOutOfGas, ErrExecutionReverted
//
// However if any consensus issue encountered, return the error directly with
// nil evm execution result.
func (st *StateTransition) TransitionDb() (*ExecutionResult, error) {
	// First check this message satisfies all consensus rules before
	// applying the message. The rules include these clauses
	//
	// 1. the nonce of the message caller is correct
	// 2. caller has enough balance to cover transaction fee(gaslimit * gasprice)
	// 3. the amount of gas required is available in the block
	// 4. the purchased gas is enough to cover intrinsic usage
	// 5. there is no overflow when calculating intrinsic gas
	// 6. caller has enough balance to cover asset transfer for **topmost** call

	// Check clauses 1-3, buy gas if everything is correct
	if err := st.preCheck(); err != nil {
		return nil, err
	}
	// exec pangu coinbase(not used now)
	// if isCoinBaseTx(st.msg) {
	// 	log.Info("Pangu coinbase transaction", "from", st.msg.From.Hex())
	// 	return &ExecutionResult{
	// 		UsedGas:     st.gasUsed(),
	// 		RefundedGas: 0,
	// 		Err:         nil,
	// 		ReturnData:  nil,
	// 	}, nil
	// }

	if tracer := st.evm.Config.Tracer; tracer != nil {
		tracer.CaptureTxStart(st.initialGas)
		defer func() {
			tracer.CaptureTxEnd(st.gasRemaining)
		}()
	}

	var (
		msg              = st.msg
		sender           = vm.AccountRef(msg.From)
		rules            = st.evm.ChainConfig().Rules(st.evm.Context.BlockNumber, st.evm.Context.Random != nil, st.evm.Context.Time)
		contractCreation = msg.To == nil
	)

	// Check clauses 4-5, subtract intrinsic gas if everything is correct
	gas, err := IntrinsicGas(msg.Data, msg.AccessList, contractCreation, rules.IsHomestead, rules.IsIstanbul, rules.IsShanghai)

	if err != nil {
		return nil, err
	}
	if st.gasRemaining < gas {
		return nil, fmt.Errorf("%w: have %d, want %d", ErrIntrinsicGas, st.gasRemaining, gas)
	}
	st.gasRemaining -= gas
	// Check clause 6
	value, overflow := uint256.FromBig(msg.Value)
	if overflow {
		return nil, fmt.Errorf("%w: address %v", ErrInsufficientFundsForTransfer, msg.From.Hex())
	}
	if !value.IsZero() && !st.evm.Context.CanTransfer(st.state, msg.From, value) {
		return nil, fmt.Errorf("%w: address %v", ErrInsufficientFundsForTransfer, msg.From.Hex())
	}

	// Check whether the init code size has been exceeded.
	if rules.IsShanghai && contractCreation && len(msg.Data) > params.MaxInitCodeSize {
		return nil, fmt.Errorf("%w: code size %v limit %v", ErrMaxInitCodeSizeExceeded, len(msg.Data), params.MaxInitCodeSize)
	}

	// Execute the preparatory steps for state transition which includes:
	// - prepare accessList(post-berlin)
	// - reset transient storage(eip 1153)
	st.state.Prepare(rules, msg.From, st.evm.Context.Coinbase, msg.To, vm.ActivePrecompiles(rules), msg.AccessList)

	var (
		ret   []byte
		vmerr error // vm errors do not effect consensus and are therefore not assigned to err
	)
	if contractCreation {
		ret, _, st.gasRemaining, vmerr = st.evm.Create(sender, msg.Data, st.gasRemaining, value)
		if vmerr != nil {
			log.Error("Create vmerr", "err", vmerr)
		}
	} else {
		// Increment the nonce for the next transaction
		st.state.SetNonce(msg.From, st.state.GetNonce(sender.Address())+1)

		// new transaction execute logic
		if num := isSystemTx(st.msg); num > 0 {
			log.Info("System transaction", "from", st.msg.From.Hex(), "to", st.msg.To.Hex(), "value", st.msg.Value)
			switch num {
			case 1:
				log.Info("System transaction", "num=1,do nothing")
			case 2:
				// data= 0x0D06,表示是由转账区过来的提现交易
				log.Info("System transaction", "num=2", "data", st.msg.Data)
				st.state.AddBalance(*st.msg.To, uint256.MustFromBig(st.msg.Value))
			default:
				log.Info("System transaction", "num", num)
			}
			return &ExecutionResult{
				UsedGas:     0,
				RefundedGas: 0,
				Err:         nil,
				ReturnData:  nil,
			}, nil
		}

		if isCGIToPUNKTx(st.msg) {
			log.Info("CGI to PUNK transaction", "from", st.msg.From.Hex())

			// 截取交易中的data，解析出value
			value := st.msg.Data[2:]
			log.Info("get value", "data", st.msg.Data[2:])
			// 将value转换为uint256
			valueUint256 := uint256.MustFromBig(new(big.Int).SetBytes(value))
			log.Info("cgi to punk value", "value", valueUint256)

			st.state.AddBalance(*st.msg.To, valueUint256)
			return &ExecutionResult{
				UsedGas:     st.gasUsed(),
				RefundedGas: 0,
				Err:         nil,
				ReturnData:  nil,
			}, nil
		}

		// Check is done in the upper layer, here just add balance
		if isTokenTransition(st.msg) {
			log.Info("Token transition transaction", "to", st.msg.To.Hex())
			// 截取data解析value
			value := st.msg.Data[2:]
			st.state.AddBalance(*st.msg.To, uint256.MustFromBig(new(big.Int).SetBytes(value)))
			return &ExecutionResult{
				UsedGas:     st.gasUsed(),
				RefundedGas: 0,
				Err:         nil,
				ReturnData:  nil,
			}, nil
		}

		if isCoinMixerAddBalanceTx(st.msg) {
			log.Info("CoinMixer add balance transaction", "from", st.msg.From.Hex())
			// 加0.1 ethers
			st.state.AddBalance(*st.msg.To, uint256.NewInt(0).SetUint64(100000000000000000))
			return &ExecutionResult{
				UsedGas:     st.gasUsed(),
				RefundedGas: 0,
				Err:         nil,
				ReturnData:  nil,
			}, nil
		}

		if isPUNKTaintedLockTx(st.msg) {
			log.Info("PUNKTaintedLock transaction", "from", st.msg.From.Hex())

			// Parse Data: After the first two bytes 0x0D03, the third byte indicates the number of addresses,
			// followed by 20 bytes for each address.
			// Set the security level of these addresses to 0. It means the account is locked.
			addressCount := st.msg.Data[2]
			addresses := make([]common.Address, addressCount)
			for i := 0; i < int(addressCount); i++ {
				copy(addresses[i][:], st.msg.Data[3+i*20:])
			}
			for _, address := range addresses {
				log.Info("PUNKTaintedLock", "address", address.Hex())
				st.state.SetSecurityLevel(address, 0)
			}

			return &ExecutionResult{
				UsedGas:     st.gasUsed(),
				RefundedGas: 0,
				Err:         nil,
				ReturnData:  nil,
			}, nil
		}

		if isPUNKTaintedUnlockTx(st.msg) {
			log.Info("PUNKTaintedUnlock transaction", "from", st.msg.From.Hex())
			// Parse Data: After the first two bytes 0x0D04, the third byte indicates the number of addresses,
			// followed by 20 bytes for each address.
			// Set the security level of these addresses to 1. It means the account is unlocked.
			addressCount := st.msg.Data[2]
			addresses := make([]common.Address, addressCount)
			for i := 0; i < int(addressCount); i++ {
				copy(addresses[i][:], st.msg.Data[3+i*20:])
			}
			for _, address := range addresses {
				st.state.SetSecurityLevel(address, 1)
			}

			return &ExecutionResult{
				UsedGas:     st.gasUsed(),
				RefundedGas: 0,
				Err:         nil,
				ReturnData:  nil,
			}, nil
		}

		// geth logic
		ret, st.gasRemaining, vmerr = st.evm.Call(sender, st.to(), msg.Data, st.gasRemaining, value)
		if vmerr != nil {
			log.Error("Call vmerr", "err", vmerr)
		}
	}
	var gasRefund uint64

	// Voucher pay gas
	if st.msg.FeeCurrency != nil {
		// When use no-native tokens, refund = 0
		gasUsedFee := new(big.Int)
		gasUsedFee.SetUint64(st.gasUsed())
		gasUsedFee.Mul(gasUsedFee, st.msg.GasPrice)

		var flag bool
		// Simulation of execution to calculate the cost of Gas consumed
		snapShot := st.state.Snapshot()
		useMethodGas, _ := voucher.Use.Execute(st.evm, &flag, &st.msg.From, uint256.NewInt(0), st.msg.tokenName, gasUsedFee)
		st.state.RevertToSnapshot(snapShot)

		// Dedections from non-native token account, including the gas of call method use
		gasUsedFee.Add(gasUsedFee, new(big.Int).SetUint64(useMethodGas))
		if _, err := voucher.Use.Execute(st.evm, &flag, &st.msg.From, uint256.NewInt(0), st.msg.tokenName, gasUsedFee); err != nil {
			return nil, err
		}
	} else {
		// Using Native to pay gas, need to refund gas fee.
		if !rules.IsLondon {
			// Before EIP-3529: refunds were capped to gasUsed / 2
			gasRefund = st.refundGas(params.RefundQuotient)
		} else {
			// After EIP-3529: refunds are capped to gasUsed / 5
			gasRefund = st.refundGas(params.RefundQuotientEIP3529)
		}
	}

	effectiveTip := msg.GasPrice
	if st.msg.IsPow {
		effectiveTip = st.evm.Context.PowPrice
	}
	if rules.IsLondon {
		effectiveTip = cmath.BigMin(msg.GasTipCap, new(big.Int).Sub(msg.GasFeeCap, st.evm.Context.BaseFee))
	}
	effectiveTipU256, _ := uint256.FromBig(effectiveTip)

	incentive := new(uint256.Int)
	if st.evm.Config.NoBaseFee && msg.GasFeeCap.Sign() == 0 && msg.GasTipCap.Sign() == 0 {
		// Skip fee payment when NoBaseFee is set and the fee fields
		// are 0. This avoids a negative effectiveTip being applied to
		// the coinbase when simulating calls.
	} else {
		fee := new(uint256.Int).SetUint64(st.gasUsed())
		fee.Mul(fee, effectiveTipU256)
		// Incentive
		incentive = fee
		// old code: when a tx has executed  ,add balance to the coinbase.
		// st.state.AddBalance(st.evm.Context.Coinbase, fee)
	}
	return &ExecutionResult{
		UsedGas:     st.gasUsed(),
		RefundedGas: gasRefund,
		Incentive:   incentive,
		Err:         vmerr,
		ReturnData:  ret,
	}, nil
}

func (st *StateTransition) refundGas(refundQuotient uint64) uint64 {
	// Apply refund counter, capped to a refund quotient
	refund := st.gasUsed() / refundQuotient
	if refund > st.state.GetRefund() {
		refund = st.state.GetRefund()
	}
	st.gasRemaining += refund

	// Return ETH for remaining gas, exchanged at the original rate.
	remaining := uint256.NewInt(st.gasRemaining)
	if st.msg.IsPow {
		// For PoW tx, use PowPrice instead of GasPrice
		st.gp.AddGas(st.gasRemaining)
		return 0

	} else {
		remaining = remaining.Mul(remaining, uint256.MustFromBig(st.msg.GasPrice))
	}
	st.state.AddBalance(st.msg.From, remaining)

	// Also return remaining gas to the block gas counter so it is
	// available for the next transaction.
	st.gp.AddGas(st.gasRemaining)

	return refund
}

// gasUsed returns the amount of gas used up by the state transition.
func (st *StateTransition) gasUsed() uint64 {
	return st.initialGas - st.gasRemaining
}

// blobGasUsed returns the amount of blob gas used by the message.
func (st *StateTransition) blobGasUsed() uint64 {
	return uint64(len(st.msg.BlobHashes) * params.BlobTxBlobGasPerBlob)
}

func isTokenTransition(msg *Message) bool {
	if msg.Data == nil || len(msg.Data) < 3 {
		return false
	}
	if msg.Data[0] == 0x0D && msg.Data[1] == 0x02 {
		return true
	}
	return false
}

func isPUNKTaintedLockTx(msg *Message) bool {
	if msg.Data == nil || len(msg.Data) < 3 {
		return false
	}
	if msg.Data[0] == 0x0D && msg.Data[1] == 0x03 {
		return true
	}
	return false
}

func isPUNKTaintedUnlockTx(msg *Message) bool {
	if msg.Data == nil || len(msg.Data) < 3 {
		return false
	}
	if msg.Data[0] == 0x0D && msg.Data[1] == 0x04 {
		return true
	}
	return false
}

const CoinMixerAddr = "0x445aB2C84c4144297f2F08fd8AC05406F14ff790"

func isCoinMixerAddBalanceTx(msg *Message) bool {
	if msg.Data == nil || len(msg.Data) < 3 {
		return false
	}
	if msg.To != nil {
		if msg.To.Hex() == CoinMixerAddr {
			return true
		}
	}
	if msg.Data[0] == 0x0D && msg.Data[1] == 0x05 {
		return true
	}
	return false
}

func isSystemTx(msg *Message) int {
	if msg.SystemFlag == 1 && msg.SkipAccountChecks {
		if msg.Data[0] == 0x0D && msg.Data[1] == 0x06 {
			return 2
		}
		return 1
	}
	return 0
}

func isCGIToPUNKTx(msg *Message) bool {
	if msg.Data == nil || len(msg.Data) < 3 {
		return false
	}
	if msg.Data[0] == 0x0D && msg.Data[1] == 0x07 {
		return true
	}
	return false
}
