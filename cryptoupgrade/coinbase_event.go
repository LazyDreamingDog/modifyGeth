package cryptoupgrade

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
)

type coinBaseBackend interface {
	BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error)
	// * Extra interface: types.Block.Transactions()
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
}

func bindCoinBaseEvent(client client, contractAddress common.Address, eventHash common.Hash) (ethereum.Subscription, chan types.Log, error) {
	// Based 1.contract address and 2.event hash to lookup
	query := ethereum.FilterQuery{
		// Addresses: []common.Address{contractAddress},
		Topics: [][]common.Hash{{eventHash}}, // Event hash
	}
	logCh := make(chan types.Log)
	// Subscribe to logs that meet FilterQuery,and logs will be stored in the logCh
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logCh)
	return sub, logCh, err
}

func lookupIncentive(client coinBaseBackend, startBlock uint64, endBlock uint64) (map[common.Address]*big.Int, error) {
	coinbaseAddEventHash := crypto.Keccak256Hash([]byte("CoinbaseAdded(string,string,uint256,address[],uint256[])"))
	rewardMap := make(map[common.Address]*big.Int)
	// Traverse blocks
	for blockNumber := startBlock; blockNumber <= endBlock; blockNumber++ {
		// Get Txs from block
		block, err := client.BlockByNumber(context.Background(), big.NewInt(int64(blockNumber)))
		if err != nil {
			return nil, err
		}
		// Traverse receipt
		for _, tx := range block.Transactions() {
			// Get receipt
			receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
			if err != nil {
				return nil, err
			}
			// Get logs
			for _, eventLog := range receipt.Logs {
				// Check event hash
				if len(eventLog.Topics) > 0 && eventLog.Topics[0] == coinbaseAddEventHash {
					// ABI decode
					vmap := make(map[string]interface{})
					err := CoinBaseABI.UnpackIntoMap(vmap, "CoinbaseAdded", eventLog.Data)
					if err != nil {
						return nil, err
					}
					addrs := vmap["selectedAddresses"].([]common.Address)
					rewards := vmap["rewards"].([]*big.Int)

					// Legital check
					if len(addrs) != len(rewards) {
						return nil, fmt.Errorf("len of addrs and rewards is unequal")
					}

					// Update
					for i := range addrs {
						addr := addrs[i]
						reward := rewards[i]

						if rewardMap[addr] == nil {
							rewardMap[addr] = new(big.Int)
							rewardMap[addr].Set(reward)
						} else {
							rewardMap[addrs[i]].Add(rewardMap[addrs[i]], rewards[i])
						}
						log.Info(fmt.Sprintf("Add %v to %s", rewards[i], addrs[i]))
					}
				}
			}
		}
	}
	return rewardMap, nil
}
