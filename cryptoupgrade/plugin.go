package cryptoupgrade

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"plugin"
	"reflect"

	"github.com/ethereum/go-ethereum/cryptoupgrade/preload"
	"github.com/ethereum/go-ethereum/log"
)

// Compile source go file to .sp file, only approve run. Below one only approve debug
func compilePlugin(src string, outputPath string) error {
	cmd := exec.Command("go", "build", "-buildmode=plugin", "-tags=purego", "-o", outputPath, src)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// fmt.Printf("go version: %v\n", runtime.Version())
	return cmd.Run()
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

func CallAlgorithm(algoName string, gas uint64, encodedInput []byte) ([]byte, uint64, error) {
	algoName = capitalString(algoName)

	log.Info(fmt.Sprintf("Call algorithm %s,encodedinput: %s\n", algoName, encodedInput))

	p, ok := preload.IsPreLoad(algoName)
	if ok {
		return callPreloadAlgo(p, gas, encodedInput)
	} else {
		pluginPath := sofilePath(algoName)
		return callUpgradeAlgo(algoName, pluginPath, gas, encodedInput)
	}

}

func callUpgradeAlgo(funcName string, pluginPath string, gas uint64, encodedInput []byte) ([]byte, uint64, error) {
	// Get algorithm info
	funcInfo := algoInfoMap[funcName]
	inputType, outputType := funcInfo.getTypeList()
	log.Info(fmt.Sprintf("Algorithm itype: %v ,otype:%v\n", inputType, outputType))

	// Gas sufficient check
	if gas < funcInfo.gas {
		return nil, gas, errors.New("out of gas")
	}

	
	// Decode input
	input, err := UnpackInput(encodedInput, inputType)
	if err != nil {
		log.Error("Unpack input from callFunc error:", err)
		fmt.Println("Unpack input from callFunc error:", err)
		return nil, gas, err
	}

	// Call algorithm. Attention 1.[]interface{} and ...interface{} 2.Algorithm name must be capital
	output, err := callPlugin(pluginPath, funcName, input)
	if err != nil {
		log.Error("Call Plugin error", err)
		fmt.Println("Call Plugin error", err)
		return nil, gas, err
	}

	// Output encode
	encodedOutput, err := PackOutput(output, outputType)
	if err != nil {
		log.Error("Pack output error:", err)
		fmt.Println("Pack output error:", err)
		return nil, gas, err
	}

	// Gas deduction and return output
	remainGas := gas - uint64(funcInfo.gas)
	log.Info(fmt.Sprintf("Successful call upgrade algorithm. Gas:%d EncodeReturn: %v", gas, encodedOutput))
	return encodedOutput, remainGas, nil
}

func callPreloadAlgo(algo preload.PreLoadAlgorithm, gas uint64, encodedInput []byte) ([]byte, uint64, error) {
	// Gas sufficient check
	if gas < algo.RequiredGas() {
		return nil, gas, errors.New("out of gas")
	}

	// Get type list
	itype, otype := algo.GetTypeList()

	// Decode input
	input, err := UnpackInput(encodedInput, itype)
	if err != nil {
		log.Error("Unpack input from callFunc error:", err)
		return nil, gas, err
	}

	// Call algorithm
	p := algo.TargetFunc()
	output, err := callFunction(p, input)
	if err != nil {
		log.Error("Pack output error:", err)
		return nil, gas, err
	}

	// Decode output
	encodedOutput, err := PackOutput(output, otype)
	if err != nil {
		log.Error("Pack output error:", err)
		return nil, gas, err
	}

	// Gas deduction
	remainGas := gas - algo.RequiredGas()
	log.Info(fmt.Sprintf("Successful call upgrade algorithm. Gas:%d EncodeReturn: %v", gas, encodedOutput))
	return encodedOutput, remainGas, err
}

// Call fun in plugin, parameter and return are mutable types
func callPlugin(pluginPath string, funName string, args []interface{}) ([]interface{}, error) {
	// Load plugin
	p, err := plugin.Open(pluginPath)
	if err != nil {
		log.Error(fmt.Sprintf("Failed to open plugin: path is %s,err %v\n", pluginPath, err))
		return nil, err
	}

	// Lookup func in plugin
	fn, err := p.Lookup(funName)
	if err != nil {
		log.Error(fmt.Sprintf("Failed to find symbol %s in plugin %s: %v\n", funName, pluginPath, err))
		return nil, err
	}

	// Call func and return
	ReturnList, err := callFunction(fn, args)
	if err != nil {
		log.Error("Error in call fun ", funName, " in plugin ", pluginPath)
		return nil, err
	} else {
		return ReturnList, nil
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

	fmt.Printf("Step 3: Call function\n")
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
