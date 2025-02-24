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
	codeUploaded   = crypto.Keccak256Hash([]byte("codeUploaded(string)"))
	algoInfoPath   = "./plugin/algorithm_info.json"
	compressedPath = "./plugin"
)

var codeStorageABI_json = `
[
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "internalType": "string",
          "name": "name",
          "type": "string"
        }
      ],
      "name": "codeUploaded",
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
          "internalType": "bytes",
          "name": "input",
          "type": "bytes"
        }
      ],
      "name": "callFunc",
      "outputs": [
        {
          "internalType": "bytes",
          "name": "",
          "type": "bytes"
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
      "name": "getCode",
      "outputs": [
        {
          "internalType": "string",
          "name": "",
          "type": "string"
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
      "name": "getGas",
      "outputs": [
        {
          "internalType": "uint64",
          "name": "",
          "type": "uint64"
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
      "name": "getInfo",
      "outputs": [
        {
          "internalType": "string",
          "name": "code",
          "type": "string"
        },
        {
          "internalType": "uint64",
          "name": "gas",
          "type": "uint64"
        },
        {
          "internalType": "string",
          "name": "itype",
          "type": "string"
        },
        {
          "internalType": "string",
          "name": "otype",
          "type": "string"
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
          "internalType": "uint64",
          "name": "_gas",
          "type": "uint64"
        }
      ],
      "name": "updataGas",
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
        },
        {
          "internalType": "string",
          "name": "code",
          "type": "string"
        },
        {
          "internalType": "uint64",
          "name": "gas",
          "type": "uint64"
        },
        {
          "internalType": "string",
          "name": "itype",
          "type": "string"
        },
        {
          "internalType": "string",
          "name": "otype",
          "type": "string"
        }
      ],
      "name": "uploadCode",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    }
  ]
`

var coinbaseABI_json = `
[
    {
      "inputs": [],
      "stateMutability": "nonpayable",
      "type": "constructor"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "string",
          "name": "source",
          "type": "string"
        },
        {
          "indexed": true,
          "internalType": "string",
          "name": "rewardType",
          "type": "string"
        },
        {
          "indexed": true,
          "internalType": "uint256",
          "name": "timestamp",
          "type": "uint256"
        },
        {
          "indexed": false,
          "internalType": "address[]",
          "name": "selectedAddresses",
          "type": "address[]"
        },
        {
          "indexed": false,
          "internalType": "uint256[]",
          "name": "rewards",
          "type": "uint256[]"
        }
      ],
      "name": "CoinbaseAdded",
      "type": "event"
    },
    {
      "inputs": [],
      "name": "STAKING_ADDRESS",
      "outputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "string",
          "name": "source",
          "type": "string"
        },
        {
          "internalType": "address[]",
          "name": "rewardAddresses",
          "type": "address[]"
        },
        {
          "internalType": "string",
          "name": "rewardType",
          "type": "string"
        },
        {
          "internalType": "uint256",
          "name": "totalAmount",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "numWinners",
          "type": "uint256"
        }
      ],
      "name": "addCoinbase",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "string",
          "name": "source",
          "type": "string"
        },
        {
          "internalType": "address[]",
          "name": "rewardAddresses",
          "type": "address[]"
        },
        {
          "internalType": "uint256[]",
          "name": "rewardAmounts",
          "type": "uint256[]"
        },
        {
          "internalType": "string",
          "name": "rewardType",
          "type": "string"
        }
      ],
      "name": "addCoinbaseDirectly",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "_address",
          "type": "address"
        }
      ],
      "name": "addToWhitelist",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "name": "whitelist",
      "outputs": [
        {
          "internalType": "bool",
          "name": "",
          "type": "bool"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    }
  ]`

// Public attributes
var (
	CodeStorageAddress          = common.BytesToAddress([]byte{67})
	CodeStorageABI, CoinBaseABI abi.ABI
)

func init() {
	// Load contract ABI
	var err error
	CodeStorageABI, err = abi.JSON(strings.NewReader(codeStorageABI_json))
	if err != nil {
		log.Error("Load codestorage ABI err")
	}

	CoinBaseABI, err = abi.JSON(strings.NewReader(coinbaseABI_json))
	if err != nil {
		log.Error("Load codestorage ABI err")
	}

	// Load stashed map
	if err = loadFromFile(algoInfoPath); err != nil {
		log.Error("Load upgrade algorithm map error:", err)
	}

	// Create directory
	directoryInit()
}
