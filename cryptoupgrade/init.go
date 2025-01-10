package cryptoupgrade

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
)

// Private attributes
var (
	pullCodeEventHash   = crypto.Keccak256Hash([]byte("pullcode(string)"))
	algoInfoPath        = "./plugin/algorithm_info.json"
	compressedPath      = "./plugin"
	codeStorageABI_json = `[
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "internalType": "string",
          "name": "voucherName",
          "type": "string"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "conversionRate",
          "type": "uint256"
        }
      ],
      "name": "VoucherCreated",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "internalType": "address",
          "name": "buyer",
          "type": "address"
        },
        {
          "indexed": false,
          "internalType": "string",
          "name": "voucherName",
          "type": "string"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "amount",
          "type": "uint256"
        }
      ],
      "name": "VoucherPurchased",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "internalType": "address",
          "name": "user",
          "type": "address"
        },
        {
          "indexed": false,
          "internalType": "string",
          "name": "voucherName",
          "type": "string"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "amount",
          "type": "uint256"
        }
      ],
      "name": "VoucherUsed",
      "type": "event"
    },
    {
      "inputs": [
        {
          "internalType": "string",
          "name": "name",
          "type": "string"
        },
        {
          "internalType": "address",
          "name": "user",
          "type": "address"
        }
      ],
      "name": "balanceOf",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "string",
          "name": "name",
          "type": "string"
        }
      ],
      "name": "buy",
      "outputs": [],
      "stateMutability": "payable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "string",
          "name": "name",
          "type": "string"
        },
        {
          "internalType": "uint256",
          "name": "conversionRate",
          "type": "uint256"
        }
      ],
      "name": "createVoucher",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "string",
          "name": "name",
          "type": "string"
        }
      ],
      "name": "getVoucherInfo",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "conversionRate",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "string",
          "name": "name",
          "type": "string"
        },
        {
          "internalType": "uint256",
          "name": "amount",
          "type": "uint256"
        }
      ],
      "name": "use",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    }
  ]`
)

// Public attributes
var (
	CodeStorageAddress = common.BytesToAddress([]byte{67})
	CodeStorageABI     abi.ABI
)

func init() {
	// Load contract ABI
	var err error
	CodeStorageABI, err = abi.JSON(strings.NewReader(codeStorageABI_json))
	if err != nil {
		log.Error("Load codestorage ABI err")
	}
	err = loadFromFile(algoInfoPath)
	if err != nil {
		log.Error("Load upgrade algorithm map error:", err)
	}
}
