package common

// Address
var (
	CodeStorageAddress = BytesToAddress([]byte{67}) // 0000000000000000000000000000000000000043
	MutiVoucherAddress = BytesToAddress([]byte{68}) // 0000000000000000000000000000000000000044
	CoinBaseAddress    = BytesToAddress([]byte{68}) // 0000000000000000000000000000000000000045
)

var MutivoucherABI_json = `[
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "internalType": "string",
          "name": "name",
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
          "name": "name",
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
          "name": "name",
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
      "inputs": [],
      "name": "decimals",
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

// ABI json
var CoinbaseABI_json = `
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

var CodeStorageABI_json = `
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

var CoinBaseDeployedCode = "0x608060405234801561001057600080fd5b50600436106100575760003560e01c806321004afa1461005c5780638c20570b146100785780639b19251a14610096578063e43252d7146100c6578063f653c70e146100e2575b600080fd5b61007660048036038101906100719190610ae6565b6100fe565b005b6100806104c1565b60405161008d9190610c03565b60405180910390f35b6100b060048036038101906100ab9190610c4a565b6104d9565b6040516100bd9190610c92565b60405180910390f35b6100e060048036038101906100db9190610c4a565b6104f9565b005b6100fc60048036038101906100f79190610d03565b6105c2565b005b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610189576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161018090610e49565b60405180910390fd5b600086869050116101cf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101c690610eb5565b60405180910390fd5b6000811180156101e25750858590508111155b610221576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161021890610f21565b60405180910390fd5b600061026e878780806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f8201169050808301925050505050505083610753565b905060008267ffffffffffffffff81111561028c5761028b610f41565b5b6040519080825280602002602001820160405280156102ba5781602001602082028036833780820191505090505b50905060008490506001841061032d5760646032826102d99190610f9f565b6102e39190611010565b826000815181106102f7576102f6611041565b5b6020026020010181815250508160008151811061031757610316611041565b5b60200260200101518161032a9190611070565b90505b60028410610398576064601e826103449190610f9f565b61034e9190611010565b8260018151811061036257610361611041565b5b6020026020010181815250508160018151811061038257610381611041565b5b6020026020010151816103959190611070565b90505b60028411156104105760006002856103b09190611070565b826103bb9190611010565b90506000600290505b8581101561040d57818482815181106103e0576103df611041565b5b60200260200101818152505081836103f89190611070565b92508080610405906110a4565b9150506103c4565b50505b600081111561044a57808260008151811061042e5761042d611041565b5b6020026020010181815161044291906110ec565b915081815250505b42878760405161045b92919061115f565b60405180910390208c8c60405161047392919061115f565b60405180910390207fb10ac242af9c4ae0bd95fcd80859174de85c68cdd6541b00a6ab5cad0d92e43a86866040516104ac9291906112f4565b60405180910390a45050505050505050505050565b7363bc05bc6fcab99af9a4c215b2e92a9c6f45d41f81565b60006020528060005260406000206000915054906101000a900460ff1681565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610568576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161055f90611377565b60405180910390fd5b60016000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1661064d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161064490610e49565b60405180910390fd5b60008686905011610693576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161068a90610eb5565b60405180910390fd5b8383905086869050146106db576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106d2906113e3565b60405180910390fd5b4282826040516106ec92919061115f565b6040518091039020898960405161070492919061115f565b60405180910390207fb10ac242af9c4ae0bd95fcd80859174de85c68cdd6541b00a6ab5cad0d92e43a8989898960405161074194939291906114f8565b60405180910390a45050505050505050565b60608251821115610799576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610790906115a5565b60405180910390fd5b60008267ffffffffffffffff8111156107b5576107b4610f41565b5b6040519080825280602002602001820160405280156107e35781602001602082028036833780820191505090505b5090506000845167ffffffffffffffff81111561080357610802610f41565b5b6040519080825280602002602001820160405280156108315781602001602082028036833780820191505090505b50905060005b85518110156108bb5785818151811061085357610852611041565b5b602002602001015182828151811061086e5761086d611041565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505080806108b3906110a4565b915050610837565b5060005b848110156109df5760008183516108d69190611070565b446108e191906115c5565b90508281815181106108f6576108f5611041565b5b602002602001015184838151811061091157610910611041565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505082826001855161095b9190611070565b6109659190611070565b8151811061097657610975611041565b5b602002602001015183828151811061099157610990611041565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250505080806109d7906110a4565b9150506108bf565b50819250505092915050565b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b60008083601f840112610a1a57610a196109f5565b5b8235905067ffffffffffffffff811115610a3757610a366109fa565b5b602083019150836001820283011115610a5357610a526109ff565b5b9250929050565b60008083601f840112610a7057610a6f6109f5565b5b8235905067ffffffffffffffff811115610a8d57610a8c6109fa565b5b602083019150836020820283011115610aa957610aa86109ff565b5b9250929050565b6000819050919050565b610ac381610ab0565b8114610ace57600080fd5b50565b600081359050610ae081610aba565b92915050565b60008060008060008060008060a0898b031215610b0657610b056109eb565b5b600089013567ffffffffffffffff811115610b2457610b236109f0565b5b610b308b828c01610a04565b9850985050602089013567ffffffffffffffff811115610b5357610b526109f0565b5b610b5f8b828c01610a5a565b9650965050604089013567ffffffffffffffff811115610b8257610b816109f0565b5b610b8e8b828c01610a04565b94509450506060610ba18b828c01610ad1565b9250506080610bb28b828c01610ad1565b9150509295985092959890939650565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610bed82610bc2565b9050919050565b610bfd81610be2565b82525050565b6000602082019050610c186000830184610bf4565b92915050565b610c2781610be2565b8114610c3257600080fd5b50565b600081359050610c4481610c1e565b92915050565b600060208284031215610c6057610c5f6109eb565b5b6000610c6e84828501610c35565b91505092915050565b60008115159050919050565b610c8c81610c77565b82525050565b6000602082019050610ca76000830184610c83565b92915050565b60008083601f840112610cc357610cc26109f5565b5b8235905067ffffffffffffffff811115610ce057610cdf6109fa565b5b602083019150836020820283011115610cfc57610cfb6109ff565b5b9250929050565b6000806000806000806000806080898b031215610d2357610d226109eb565b5b600089013567ffffffffffffffff811115610d4157610d406109f0565b5b610d4d8b828c01610a04565b9850985050602089013567ffffffffffffffff811115610d7057610d6f6109f0565b5b610d7c8b828c01610a5a565b9650965050604089013567ffffffffffffffff811115610d9f57610d9e6109f0565b5b610dab8b828c01610cad565b9450945050606089013567ffffffffffffffff811115610dce57610dcd6109f0565b5b610dda8b828c01610a04565b92509250509295985092959890939650565b600082825260208201905092915050565b7f43616c6c6572206973206e6f742077686974656c697374656400000000000000600082015250565b6000610e33601983610dec565b9150610e3e82610dfd565b602082019050919050565b60006020820190508181036000830152610e6281610e26565b9050919050565b7f4e6f206164647265737365732070726f76696465640000000000000000000000600082015250565b6000610e9f601583610dec565b9150610eaa82610e69565b602082019050919050565b60006020820190508181036000830152610ece81610e92565b9050919050565b7f496e76616c6964206e756d626572206f662077696e6e65727300000000000000600082015250565b6000610f0b601983610dec565b9150610f1682610ed5565b602082019050919050565b60006020820190508181036000830152610f3a81610efe565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610faa82610ab0565b9150610fb583610ab0565b9250828202610fc381610ab0565b91508282048414831517610fda57610fd9610f70565b5b5092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600061101b82610ab0565b915061102683610ab0565b92508261103657611035610fe1565b5b828204905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600061107b82610ab0565b915061108683610ab0565b925082820390508181111561109e5761109d610f70565b5b92915050565b60006110af82610ab0565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036110e1576110e0610f70565b5b600182019050919050565b60006110f782610ab0565b915061110283610ab0565b925082820190508082111561111a57611119610f70565b5b92915050565b600081905092915050565b82818337600083830152505050565b60006111468385611120565b935061115383858461112b565b82840190509392505050565b600061116c82848661113a565b91508190509392505050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6111ad81610be2565b82525050565b60006111bf83836111a4565b60208301905092915050565b6000602082019050919050565b60006111e382611178565b6111ed8185611183565b93506111f883611194565b8060005b8381101561122957815161121088826111b3565b975061121b836111cb565b9250506001810190506111fc565b5085935050505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b61126b81610ab0565b82525050565b600061127d8383611262565b60208301905092915050565b6000602082019050919050565b60006112a182611236565b6112ab8185611241565b93506112b683611252565b8060005b838110156112e75781516112ce8882611271565b97506112d983611289565b9250506001810190506112ba565b5085935050505092915050565b6000604082019050818103600083015261130e81856111d8565b905081810360208301526113228184611296565b90509392505050565b7f496e76616c696420616464726573730000000000000000000000000000000000600082015250565b6000611361600f83610dec565b915061136c8261132b565b602082019050919050565b6000602082019050818103600083015261139081611354565b9050919050565b7f417272617973206c656e677468206d69736d6174636800000000000000000000600082015250565b60006113cd601683610dec565b91506113d882611397565b602082019050919050565b600060208201905081810360008301526113fc816113c0565b9050919050565b6000819050919050565b600061141c6020840184610c35565b905092915050565b6000602082019050919050565b600061143d8385611183565b935061144882611403565b8060005b858110156114815761145e828461140d565b61146888826111b3565b975061147383611424565b92505060018101905061144c565b5085925050509392505050565b600080fd5b82818337505050565b60006114a88385611241565b93507f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8311156114db576114da61148e565b5b6020830292506114ec838584611493565b82840190509392505050565b60006040820190508181036000830152611513818688611431565b9050818103602083015261152881848661149c565b905095945050505050565b7f43616e6e6f742073656c656374206d6f7265206164647265737365732074686160008201527f6e2070726f766964656400000000000000000000000000000000000000000000602082015250565b600061158f602a83610dec565b915061159a82611533565b604082019050919050565b600060208201905081810360008301526115be81611582565b9050919050565b60006115d082610ab0565b91506115db83610ab0565b9250826115eb576115ea610fe1565b5b82820690509291505056fea264697066735822122012f543141eb3a72fd6a02179e8760c970840d48dd9271b6c166db8d675bfe80b64736f6c63430008120033"
