package punkos

import (
	"context"
	"log"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"path/filepath"
)

// 公共测试账号
var sender, receiver *keystore.Key

// rpc url
var rpcUrl = "http://localhost:36054"

func init() {
	senderKeyPath := filepath.Join("..", "testData", "UTC--2025-01-21T10-52-30.971266200Z--f61bdf96dc06685065337b76659bebac9cbb53bb")
	receiverKeyPath := filepath.Join("..", "testData", "UTC--2025-01-21T09-25-14.260753200Z--d1b194ac5281ccdb83d3058e034ca676f54269fe")

	senderKeyJson, err := ioutil.ReadFile(senderKeyPath)
	if err != nil {
		log.Fatalf("Failed to read sender key file: %v", err)
	}

	receiverKeyJson, err := ioutil.ReadFile(receiverKeyPath)
	if err != nil {
		log.Fatalf("Failed to read receiver key file: %v", err)
	}

	sender, err = keystore.DecryptKey(senderKeyJson, "123456")
	if err != nil {
		log.Fatalf("Failed to decrypt sender key: %v", err)
	}

	receiver, err = keystore.DecryptKey(receiverKeyJson, "123456")
	if err != nil {
		log.Fatalf("Failed to decrypt receiver key: %v", err)
	}
}

func TestSendTransaction(t *testing.T) {
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		t.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// JSON string containing the key
	privateKey := sender.PrivateKey
	accountAddress := sender.Address

	// Query the initial balance
	initialBalance, err := client.BalanceAt(context.Background(), accountAddress, nil)
	if err != nil {
		t.Fatalf("Failed to get initial balance: %v", err)
	}
	t.Logf("Initial Balance: %s", initialBalance.String())

	// Ensure the account has sufficient funds
	if initialBalance.Cmp(big.NewInt(1000000000000000000)) < 0 {
		t.Fatalf("Insufficient funds in the account")
	}

	// Get the nonce
	nonce, err := client.PendingNonceAt(context.Background(), accountAddress)
	if err != nil {
		t.Fatalf("Failed to get nonce: %v", err)
	}

	// Get the gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		t.Fatalf("Failed to get gas price: %v", err)
	}

	// Create the transaction
	toAddress := common.HexToAddress(receiver.Address.Hex())
	value := big.NewInt(100000000000000000) // 1 ETH
	gasLimit := uint64(1000000)             // Adjusted gas limit
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

	// TODO 发送添加Interest的交易

	// Sign the transaction
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		t.Fatalf("Failed to get chain ID: %v", err)
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		t.Fatalf("Failed to sign transaction: %v", err)
	}

	// Send the transaction
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		t.Fatalf("Failed to send transaction: %v", err)
	}

	t.Logf("Transaction sent: %s", signedTx.Hash().Hex())

	// Wait for the transaction to be mined
	receipt, err := waitForReceipt(client, signedTx.Hash())
	if err != nil {
		t.Fatalf("Failed to get transaction receipt: %v", err)
	}

	if receipt.Status != 1 {
		t.Fatalf("Transaction failed")
	}

	// Query the final balance
	finalBalance, err := client.BalanceAt(context.Background(), accountAddress, nil)
	if err != nil {
		t.Fatalf("Failed to get final balance: %v", err)
	}
	t.Logf("Final Balance: %s", finalBalance.String())

	// Check if the balance has decreased
	if finalBalance.Cmp(initialBalance) >= 0 {
		t.Fatalf("Balance did not decrease after transaction")
	}
}

func waitForReceipt(client *ethclient.Client, txHash common.Hash) (*types.Receipt, error) {
	for {
		receipt, err := client.TransactionReceipt(context.Background(), txHash)
		if err == nil {
			return receipt, nil
		}
		if err.Error() != "transaction indexing is in progress" {
			return nil, err
		}
		time.Sleep(time.Second)
	}
}
