package cryptoupgrade

import (
	"fmt"
	"math/big"
	"plugin"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

var (
	sourceFile       = "./testgo/add.go"
	pluginFile       = "./testgo/add.so"
	modulepluginFile = "./testgo/module_main.so"
)

func TestCompilePlugin(t *testing.T) {
	compilePlugin(sourceFile, pluginFile)
	// compileModulePlugin(sourceFile, modulepluginFile)
}

func TestCallPlugin(t *testing.T) {
	funcName := "Add"
	input := []interface{}{1, 1}
	out := callPlugin(pluginFile, funcName, input)
	t.Logf("out: %v\n", out)
	// Mutiple return
	funcName = "Pair"
	input = []interface{}{1, "Hello world!"}
	out = callPlugin(pluginFile, funcName, input)
	t.Logf("out: %v\n", out)
}

func TestPluginInProj(t *testing.T) {
	// Encode
	pluginFile = "/home/ubuntu/project/modifyGeth/build/bin/plugin/so/add.so"
	funcName := "add"
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
	}
	// Construct Info
	upgradeAlgorithmInfo[funcName] = codeInfo{
		code:  "",
		gas:   10,
		itype: "int256,int256",
		otype: "int256",
	}
	// Call plugin
	output, gas := CallAlgorithm(funcName, 100, input)
	fmt.Printf("output(Hex): %v\n", common.Bytes2Hex(output))
	fmt.Printf("gas: %v\n", gas)
}

type student struct {
	ID   int
	Name string
}

func TestStructInPPlugin(t *testing.T) {
	funcName := "NewStudent"
	p, _ := plugin.Open(pluginFile)
	NewStudentSymbol, err := p.Lookup(funcName)
	if err != nil {
		t.Error("Struct Name don't found")
	}
	var newStudent func(int, string) student
	newStudent, ok := NewStudentSymbol.(func(int, string) student)
	if !ok {
		t.Error("Symbol is not of expected type")
	}
	instance := newStudent(123, "liuqi")
	fmt.Printf("instance: %v\n", instance)
}
