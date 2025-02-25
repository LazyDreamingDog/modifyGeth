package cryptoupgrade

import (
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

// Decode ABI data, using contract abi.
func TestABICode(t *testing.T) {
	var (
		funcName = "vdf"
		input    []byte
	)
	input = []byte("Hello world")

	callFuncAbi := CodeStorageABI.Methods["callFunc"]
	funcSelector := callFuncAbi.ID
	encodedparams, err := callFuncAbi.Inputs.Pack(funcName, input)
	if err != nil {
		t.Error("Encode error")
	} else {
		t.Logf("encodedparams: %v\n", encodedparams)
	}
	temp := UnpackCall(append(funcSelector, encodedparams...))

	t.Logf("temp[0].(string): %v\n", temp[0].(string))
	t.Logf("temp[1].([]byte): %s\n", temp[1].([]byte))
}

func TestPackVar(t *testing.T) {
	intType, _ := abi.NewType("int256", "", nil)
	args := abi.Arguments{
		abi.Argument{Type: intType},
		abi.Argument{Type: intType},
	}
	t.Log(args)
	a := big.NewInt(100)
	b := big.NewInt(100)
	encodeData, err := args.Pack(a, b)
	if err != nil {
		t.Error("Encode err:", err)
	} else {
		hexData := common.Bytes2Hex(encodeData)
		t.Log(fmt.Sprintf("outputdata: %v , length: %d", hexData, len(hexData)))
	}
	// Decode
	params, err := args.Unpack(encodeData)
	if err != nil {
		t.Error("decode err:", err)
	} else {
		ra := params[0].(*big.Int)
		rb := params[1].(*big.Int)
		t.Log(ra, rb)
	}
}

// Try to decode abi data, only using params type not contract abi.
func TestDecode(t *testing.T) {
	// Define solidity type
	stringType, _ := abi.NewType("string", "", nil)
	uint256Type, _ := abi.NewType("uint256", "", nil)
	addressType, _ := abi.NewType("address", "", nil)
	intType, _ := abi.NewType("int256", "", nil)
	args := abi.Arguments{
		abi.Argument{Type: uint256Type},
		abi.Argument{Type: addressType},
		abi.Argument{Type: stringType},
		abi.Argument{Type: intType},
	}
	a := big.NewInt(11)
	b := common.BytesToAddress([]byte{10})
	c := "liuqi"
	d := big.NewInt(1)
	encodeData, err := args.Pack(a, b, c, d)
	if err != nil {
		t.Error("Encode err:", err)
	}
	decodeData, err := args.Unpack(encodeData)
	if err != nil {
		t.Error("Decode err", err)
	}
	t.Log(decodeData[0].(*big.Int))
	t.Log(decodeData[1].(common.Address))
	t.Log(decodeData[2].(string))
	t.Log(decodeData[3].(*big.Int))
}

func TestParseEvent(t *testing.T) {
	t.Run("codeupload", func(t *testing.T) {
		data := "0x000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000036164640000000000000000000000000000000000000000000000000000000000"
		var name string
		err := CodeStorageABI.UnpackIntoInterface(&name, "codeUploaded", common.Hex2Bytes(data))
		if err != nil {
			t.Error(err)
			t.Fail()
		} else {
			t.Log(name)
		}
	})

	t.Run("neg", func(t *testing.T) {
		CoinBaseABI, _ := abi.JSON(strings.NewReader(common.CoinbaseABI_json))
		// Attention data must delete "0x"
		data := "00000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000000100000000000000000000000057f96028ba3258ebfb4940d67443967cf23e3fc4000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000de"
		vmap := make(map[string]interface{})
		err := CoinBaseABI.UnpackIntoMap(vmap, "CoinbaseAdded", common.Hex2Bytes(data))
		if err != nil {
			t.Fatal("abi decode:", err)
		}
		t.Log(vmap)

		// selectedAddresses := vmap["selectedAddresses"].([]common.Address)
		// rewards := vmap["rewards"].([]*big.Int)
		// t.Log("selectedAddresses:", selectedAddresses)
		// t.Log("rewards:", rewards)
	})
}

func TestMethodID(t *testing.T) {
	// Check ABI is illegal
	for i, c := range CodeStorageABI.Methods {
		t.Logf("%s:%v\n", i, c)
	}
	t.Log(CodeStorageABI.Methods["callFunc"].ID)
}
