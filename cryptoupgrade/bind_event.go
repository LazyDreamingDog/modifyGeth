package cryptoupgrade

import (
	"bytes"
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
)

// Define Interface to avoid cricle import ethclient->core->upgradecrptoupgrade
type client interface {
	SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error)
	CallContract(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error)
}

// * parse event from receipt
func ParseReceipt(receipt *types.Receipt) {
	eventSignature := []byte("pullcode(string)")
	eventHash := crypto.Keccak256Hash(eventSignature)

	for _, vLog := range receipt.Logs {
		fmt.Println(vLog)
		if vLog.Topics[0] == eventHash {
			fmt.Println("Event found in transaction logs!")
			// Event in codestorage.sol
			type pullcode struct {
				name string
			}
			var event pullcode
			err := CodeStorageABI.UnpackIntoInterface(&event, "pullcode", vLog.Data)
			if err != nil {
				fmt.Printf("Failed to unpack log data: %v", err)
			}
			fmt.Printf("Event data: Name=%s\n", event.name)
		}
	}
}

func BindPullcode(client client) {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{CodeStorageAddress},
		Topics:    [][]common.Hash{{pullCodeEventHash}}, // Event hash
	}
	logCh := make(chan types.Log)
	// Subscribe to logs that meet FilterQuery,and logs will be stored in the logCh
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logCh)
	if err != nil {
		log.Error("Failed to subscribe to logs: %v", err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Info("Error while listening for logs: %v", err)
		case Log := <-logCh:
			// Parse name from event
			log.Info("Catch pull code event!")
			var name string
			err = CodeStorageABI.UnpackIntoInterface(&name, "pullcode", Log.Data)
			if err != nil {
				log.Error("Decode log data err!")
			} else if name == "" {
				log.Error("Err in parse name from pullcode event")
			}
			pc := lookupCodeInfo(client, name)
			if pc == nil {
				return
			}
			// Decompressed string to gofile
			sourcefilePath := compressedPath + "src/" + name + ".go"
			err = decompressStringToFile(pc.code, sourcefilePath)
			if err == nil {
				log.Info(fmt.Sprintf("Decompressed algorithm %s to %s", name, sourcefilePath))
			} else {
				log.Error(fmt.Sprintf("Error in decompressed algorithm %s to %s.Err:%v", name, sourcefilePath, err))
			}
			goVerison := GoversionCheck()
			// Compiled to .so file
			pluginfilePath := compressedPath + "so/" + name + ".so"
			err = compilePlugin(sourcefilePath, pluginfilePath)
			if err != nil {
				log.Error(fmt.Sprintf("Error in plugin compile:%v", err))
			} else {
				log.Info(fmt.Sprintf("Using go version: %s,compiled algorithm to %s.", goVerison, pluginfilePath))
				upgradeAlgorithmInfo[name] = *pc
			}
		}
	}

}

// * Through Client call contract, get the infomation of @name algorithm
func lookupCodeInfo(client client, name string) *codeInfo {

	// Must equal to method in codestorage contract
	lookupFuncName := "getInfo"
	input, err := CodeStorageABI.Pack(lookupFuncName, name)
	if err != nil {
		log.Error(fmt.Sprintf("Call contract codeStorage err! Algorithm:%s", name))
	}
	msg := ethereum.CallMsg{
		To:   &CodeStorageAddress,
		Data: input,
	}
	output, err := client.CallContract(context.Background(), msg, nil)
	if err != nil {
		log.Error(fmt.Sprintf("failed to call contract: %v", err))
	}

	ci, err := CodeStorageABI.Unpack("getInfo", output)
	if err != nil {
		log.Error("Failed to unpack ouput of getInfo in codeStorage contract:", err)
		return nil
	} else {
		log.Info(fmt.Sprintf("Successful get infomation of %s", name))
		// TODO exception handing
		return &codeInfo{
			code:  ci[0].(string),
			gas:   ci[1].(uint64),
			itype: ci[2].(string),
			otype: ci[3].(string),
		}
	}
}

// Check whether is callFunc in codestorage contract
func IsUpgradeAlgorithm(addr common.Address, funcSelector []byte) bool {
	return addr == CodeStorageAddress && bytes.Equal(CodeStorageABI.Methods["callFunc"].ID, funcSelector)
}

func convertBytesToString(array [][10]byte) []string {
	length := len(array)
	result := make([]string, length)
	for i := range array {
		// Convert array to slice
		result[i] = string(array[i][:])
	}
	return result
}
