package voucher

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/holiman/uint256"
)

// Method in contract, which will be defined at initial mitigation
type Method struct {
	abi      *abi.ABI
	readOnly bool
	name     string
	maxGas   uint64
}

func (m Method) Name() string { return m.name }

func NewMethod(abi *abi.ABI, method string, readOnly bool, maxGas uint64) Method {
	return Method{
		abi:      abi,
		name:     method,
		readOnly: readOnly,
		maxGas:   maxGas,
	}
}

// Bind method to a case of contract
func (m Method) Bind(contractAddress *common.Address) *BoundMethod {
	return &BoundMethod{
		// Attention shallow copy
		Method:          m,
		contractAddress: contractAddress,
	}
}

// encodeCall will encodes the msg into []byte format for EVM consumption
func (am Method) encodeCall(args ...interface{}) ([]byte, error) {
	return am.abi.Pack(am.name, args...)
}

// decodeResult will decode the result of msg execution into the result parameter
func (am Method) decodeResult(result interface{}, output []byte) error {
	if result == nil {
		return nil
	}
	return am.abi.UnpackIntoInterface(result, am.name, output)
}

// Generates ABI for a given method and its arguments.
func GetEncodedAbi(methodSelector []byte, varAbis [][]byte) []byte {
	encodedVarsAbiByteSize := 0
	for _, varAbi := range varAbis {
		encodedVarsAbiByteSize += len(varAbi)
	}
	encodedAbi := make([]byte, len(methodSelector)+encodedVarsAbiByteSize)

	copy(encodedAbi[0:len(methodSelector)], methodSelector[:])

	copyCursor := len(methodSelector)
	for _, varAbi := range varAbis {
		copy(encodedAbi[copyCursor:copyCursor+len(varAbi)], varAbi[:])
		copyCursor += len(varAbi)
	}

	return encodedAbi
}

// Method which is bounded to a case of contract
type BoundMethod struct {
	Method
	contractAddress *common.Address
}

// Construct function
func NewBoundMethod(contractAddress *common.Address, abi *abi.ABI, methodName string, readOnly bool, maxGas uint64) *BoundMethod {
	return NewMethod(abi, methodName, readOnly, maxGas).Bind(contractAddress)
}

// Execute executes the method with the given EVM and unpacks the return value into result.
// Result need to match the contract return type, and is a pointer.
func (bm *BoundMethod) Execute(evm *vm.EVM, result interface{}, caller *common.Address, value *uint256.Int, args ...interface{}) (uint64, error) {
	// fmt.Printf("call contract:%s, methodName:%s\n", bm.contractAddress, bm.Name())

	// Encode input
	input, err := bm.encodeCall(args...)
	if err != nil {
		return 0, err
	}

	// Execute triage based on read or write
	var output []byte
	gasRemain := uint64(0)
	if bm.readOnly {
		output, gasRemain, err = evm.StaticCall(vm.AccountRef(*caller), *bm.contractAddress, input, bm.maxGas)
	} else {
		output, gasRemain, err = evm.Call(vm.AccountRef(*caller), *bm.contractAddress, input, bm.maxGas, value)
	}

	// Output decode
	if err != nil {
		return 0, err
	}

	if len(output) != 0 {
		if err = bm.decodeResult(result, output); err != nil {
			return 0, err
		}
	}
	// fmt.Println("EVM call successful", "input", hexutil.Encode(input), "output", hexutil.Encode(output))
	return bm.maxGas - gasRemain, nil
}
