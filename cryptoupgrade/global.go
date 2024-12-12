package cryptoupgrade

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// * I hope the address of the contract is fixed, which involves the design of the underlying code
var (
	CodeStorageAddress = common.BytesToAddress([]byte{67})
	pullCodeEventHash  = crypto.Keccak256Hash([]byte("pullcode(string)"))
	// ! Last '/' Essential
	compressedPath       = "./plugin/"
	CodeStorageABI, _, _ = abi.LoadHardhatContract("/home/ubuntu/project/modifyGeth/cryptoupgrade/contract/CodeStorage.json")
)

// TODO If geth restart, map will be nil
var upgradeAlgorithmInfo = make(map[string]codeInfo)

type codeInfo struct {
	code  string
	gas   uint64
	itype string
	otype string
}

func (c *codeInfo) getTypeList() ([]string, []string) {
	ilist := strings.Split(c.itype, ",")
	olist := strings.Split(c.otype, ",")
	return ilist, olist
}
