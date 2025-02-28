package cryptoupgrade

import (
	"fmt"
	"math/big"
	"plugin"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

func TestAdd(t *testing.T) {
	sourceFile := "./plugin/src/Add.go"
	pluginFile := "./plugin/so/Add.so"

	compilePlugin(sourceFile, pluginFile)

	funcName := "Add"
	intType, _ := abi.NewType("int256", "", nil)
	args := abi.Arguments{
		abi.Argument{Type: intType},
		abi.Argument{Type: intType},
	}
	a := big.NewInt(100)
	b := big.NewInt(100)
	input, err := args.Pack(a, b)
	if err != nil {
		t.Error("Encode err:", err)
	} else {
		t.Logf("input(Hex):%x", input)
	}
	// Construct Info
	algoInfoMap[funcName] = algoInfo{
		code:  "",
		gas:   10,
		itype: "int256,int256",
		otype: "int256",
	}
	// Call plugin
	output, gas, err := callUpgradeAlgo(funcName, pluginFile, 100, input)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	t.Logf("output: %v\n", common.Bytes2Hex(output))
	t.Logf("gas: %v\n", gas)
}

// ! Go plugin don't approve struct export, only approve function and variable
func TestStructInPlugin(t *testing.T) {
	funcName := "NewStudent"
	pluginFile := ""
	p, _ := plugin.Open(pluginFile)
	NewStudentSymbol, err := p.Lookup(funcName)
	if err != nil {
		t.Error("Struct Name don't found")
	}

	type student struct {
		ID   int
		Name string
	}

	var newStudent func(int, string) student
	newStudent, ok := NewStudentSymbol.(func(int, string) student)
	if !ok {
		t.Error("Symbol is not of expected type")
	}
	instance := newStudent(123, "liuqi")
	fmt.Printf("instance: %v\n", instance)
}

func TestBlake2b(t *testing.T) {
	funcName := "Sum256"
	srcFile := "./plugin/src/blake2b.go"

	plugFile := "./plugin/so/" + funcName + ".so"
	compilePlugin(srcFile, plugFile)

	// Encode
	bytesType, _ := abi.NewType("bytes", "", nil)
	args := abi.Arguments{
		abi.Argument{Type: bytesType},
	}
	a := []byte("Hello world!")
	input, err := args.Pack(a)
	if err != nil {
		t.Error("Encode err:", err)
	}else{
		t.Logf("\ninput: %s\n input(Hex):%x \n", input, input)
	}
	// Construct Info
	algoInfoMap[funcName] = algoInfo{
		code:  "",
		gas:   10,
		itype: "bytes",
		otype: "bytes32",
	}
	// Call plugin
	output, gas, err := callUpgradeAlgo(funcName, plugFile, 100, input)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("output(Hex): %v\n", output)
	t.Logf("gas: %v\n", gas)
}

func TestBlake2s(t *testing.T) {
	funcName := "Sum256"
	srcFile := "./testgo/src/blake2s.go"

	plugFile := "./testgo/so/" + funcName + ".so"
	compilePlugin(srcFile, plugFile)

	// Encode input
	bytesType, _ := abi.NewType("bytes", "", nil)
	args := abi.Arguments{
		abi.Argument{Type: bytesType},
	}
	a := []byte("Hello world!")
	input, err := args.Pack(a)
	if err != nil {
		t.Error("Encode err:", err)
	}
	// Construct Info
	algoInfoMap[funcName] = algoInfo{
		code:  "",
		gas:   10,
		itype: "bytes",
		otype: "bytes32",
	}
	// Call plugin
	output, gas, _ := callUpgradeAlgo(funcName, plugFile, 100, input)
	t.Logf("output(Hex): %v\n", common.Bytes2Hex(output))
	t.Logf("gas: %v\n", gas)
}

func TestBls(t *testing.T) {
	funcName := "KeygenWithSeedAPI"
	// srcFile := "./testgo/src/bls.go"

	pluginFile := "./testgo/so/bls.so"
	// plugFile := "./testgo/so/" + funcName + ".so"
	// compilePlugin(srcFile, plugFile)

	// Encode input
	bytesType, _ := abi.NewType("bytes", "", nil)
	args := abi.Arguments{
		abi.Argument{Type: bytesType},
	}
	a := []byte("23234")
	input, err := args.Pack(a)
	if err != nil {
		t.Error("Encode err:", err)
	}
	// Construct Info
	algoInfoMap[funcName] = algoInfo{
		code:  "",
		gas:   10,
		itype: "bytes",
		otype: "bytes,bytes",
	}
	// Call plugin
	output, gas, _ := callUpgradeAlgo(funcName, pluginFile, 100, input)
	t.Logf("output(Hex): %v\n", common.Bytes2Hex(output))
	t.Logf("gas: %v\n", gas)
}
