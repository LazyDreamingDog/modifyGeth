package cryptoupgrade

import (
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type AddAlgo struct{}

func (v *AddAlgo) RequiredGas() uint64     { return 100 }
func (v *AddAlgo) TargetFunc() interface{} { return interface{}(Add) }
func (v *AddAlgo) GetTypeList() ([]string, []string) {
	itype := []string{"uint64", "uint64"}
	otype := []string{"uint64"}
	return itype, otype
}

func Add(a, b uint64) uint64 {
	return a + b
}

func TestPreloadAdd(t *testing.T) {

	intType, _ := abi.NewType("uint64", "", nil)
	args := abi.Arguments{
		abi.Argument{Type: intType},
		abi.Argument{Type: intType},
	}

	// Construct input
	a := uint64(1)
	b := uint64(1)
	input, err := args.Pack(a, b)
	if err != nil {
		t.Error("Encode err:", err)
		t.Fail()
	}

	// Construct generic function
	p := &AddAlgo{}
	// Call plugin
	output, gas, err := callPreloadAlgo(p, 100, input)
	if err != nil || gas != 0 {
		t.Error(err)
		t.Fail()
	}

	t.Logf("output(Hex): %v\n", common.Bytes2Hex(output))
	t.Logf("gas: %v\n", gas)
}
