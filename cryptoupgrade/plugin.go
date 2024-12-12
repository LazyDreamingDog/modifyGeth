package cryptoupgrade

import (
	"fmt"
	"os"
	"os/exec"
	"plugin"
	"reflect"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

// Compile source go file to .sp file, only approve run. Below one only approve debug
func compilePlugin(src string, outputPath string) error {
	cmd := exec.Command("go", "build", "-buildmode=plugin", "-tags", "urfave_cli_no_docs,ckzg", "-o", outputPath, src)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	GoversionCheck()
	return cmd.Run()
}

// Get go version
func GoversionCheck() string {
	// Get go version
	goverCmd := exec.Command("go", "version")
	output, err := goverCmd.Output()
	if err != nil {
		fmt.Printf("go version command err: %v\n", err)
	}
	return string(output)
}

// ! There may be conflicts between the go version and the go compiler version, so If you choose a plugin that supports dbg, it can only support debugging but not normal execution.
func compileModulePlugin(src string, outputPath string) error {
	cmd := exec.Command("go", "build", "-buildmode=plugin", "-gcflags=all=-N -l", "-o", outputPath, src)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Convert str's first char to capital
func capitalString(str string) string {
	bytes := []byte(str)
	if bytes[0] >= 'a' && bytes[0] <= 'z' {
		bytes[0] = bytes[0] - 32
	}
	return string(bytes)
}

func CallAlgorithm(funName string, gas uint64, encodedInput []byte) ([]byte, uint64) {
	fmt.Printf("Call algorithm %s,encodedinput: %s\n", funName, common.Bytes2Hex(encodedInput))
	pluginPath := compressedPath + "so/" + funName + ".so"
	// Decode input
	funcInfo := upgradeAlgorithmInfo[funName]
	inputType, outputType := funcInfo.getTypeList()
	fmt.Printf("Algorithm itype: %v ,otype:%v\n", inputType, outputType)
	input, err := UnpackInput(encodedInput, inputType)
	if err != nil {
		log.Error("Unpack input from callFunc error:", err)
	}
	// Attention 1.[]interface{} and ...interface{} 2.Algorithm capital
	output := callPlugin(pluginPath, capitalString(funName), input)
	encodedOutput, err := PackOutput(output, outputType)
	if err != nil {
		log.Error("Pack output error:", err)
	}
	remainGas := gas - uint64(funcInfo.gas)
	log.Info(fmt.Sprintf("Successful call upgrade algorithm. Gas:%d EncodeReturn: %v", gas, encodedOutput))
	return encodedOutput, remainGas
}

// Call fun in plugin, parameter and return are mutable types
func callPlugin(pluginPath string, funName string, args []interface{}) []interface{} {
	p, err := plugin.Open(pluginPath)
	if err != nil {
		log.Error(fmt.Sprintf("Failed to open plugin: path %s,err %v\n", pluginPath, err))
		// fmt.Printf("Failed to open plugin: path %s,err %v\n", pluginPath, err)
		return nil
	}
	fn, err := p.Lookup(funName)
	if err != nil {
		log.Error(fmt.Sprintf("Failed to find symbol %s in plugin %s: %v\n", funName, pluginPath, err))
		// fmt.Printf("Failed to find symbol %s in plugin %s: %v\n", funName, pluginPath, err)
		return nil
	}
	ReturnList, err := callFunction(fn, args)
	if err != nil {
		log.Error("Error in call fun ", funName, " in plugin ", pluginPath)
		// fmt.Println("Error in call fun ", funName, " in plugin ", pluginPath)
		return nil
	} else {
		return ReturnList
	}
}

// Provide a unified calling entry. Return fn's return as an interface{} list
func callFunction(fn interface{}, args []interface{}) ([]interface{}, error) {
	v := reflect.ValueOf(fn)
	fmt.Printf("Step 1: Checking if fn is a function: Kind = %s\n", v.Kind())
	// Chech whether fn is function type
	if v.Kind() != reflect.Func {
		return nil, fmt.Errorf("provided value is not a function")
	}
	fmt.Printf("Step 2: Construct input\n")
	// Construct parameters to input
	in := make([]reflect.Value, len(args))
	for i, arg := range args {
		fmt.Printf("Arg[%d]: Type = %T, Value = %v\n", i, arg, arg)
		in[i] = reflect.ValueOf(arg)
	}
	fmt.Printf("Step 3: Call\n")
	// Reflect Call
	result := v.Call(in)
	out := make([]interface{}, len(result))
	fmt.Printf("Step 4: Get return\n")
	for i, r := range result {
		// Convert reflect.Value to interface{}
		fmt.Printf("Result[%d]: Type = %v, Value = %v\n", i, r.Type(), r.Interface())
		out[i] = r.Interface()
	}
	return out, nil
}
