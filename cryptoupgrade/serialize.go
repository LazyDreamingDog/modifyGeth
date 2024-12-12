package cryptoupgrade

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/log"
)

// Decode callFunc(name,input), which is abi code
func UnpackCall(data []byte) []interface{} {
	// First four byte is function selector, remain are params encoding with abi.
	abiEncodedInput := data[4:]
	callFuncAbi := CodeStorageABI.Methods["callFunc"]
	temp, err := callFuncAbi.Inputs.Unpack(abiEncodedInput)
	if err != nil {
		log.Error("Failed to unpack inputs: ", err)
	}

	return temp
}

// Decode input in callFunc(name,input)
func UnpackInput(encodedInput []byte, paramsType []string) ([]interface{}, error) {
	var args abi.Arguments
	// Construct arguments list
	for i := range paramsType {
		NType, err := abi.NewType(paramsType[i], "", nil)
		if err == nil {
			args = append(args, abi.Argument{Type: NType})
		} else {
			return nil, err
		}
	}
	fmt.Printf("Args: %v\n", args)
	params, err := args.Unpack(encodedInput)
	if err != nil {
		return nil, err
	}
	return params, nil
}

// Encode output with abi
func PackOutput(output []interface{}, returnType []string) ([]byte, error) {
	var args abi.Arguments
	var decodeOutput []byte
	// Construct arguments list
	for i := range returnType {
		NType, err := abi.NewType(returnType[i], "", nil)
		if err == nil {
			args = append(args, abi.Argument{Type: NType})
		} else {
			return nil, err
		}
	}
	decodeOutput, err := args.Pack(output...)
	if err != nil {
		return nil, err
	}
	return decodeOutput, err
}
