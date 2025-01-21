package rpc_punkos

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var sender = `{"address":"72c5da1e05cc7bba112a1cf55982bb63ca026bbd","crypto":{"cipher":"aes-128-ctr","ciphertext":"3fdfe9d5bcd265b3e27025a791d2ffdc567522b45eb15492d6381f0a9d097b77","cipherparams":{"iv":"2c12e45fc166d2563c1f3ba81efaf761"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"f7b46acde1ec6e498c7591126789b0cdffc2f6d5268ea6f23771b602135eb656"},"mac":"bd757c7c7cafb42ca674b4cc377a80b45fc20629a5f8f7ce76d81cec0194c592"},"id":"bcec0e55-76b6-4760-8b9f-c256fff86964","version":3}`
var receiver = "0xd1b194ac5281ccdb83d3058e034ca676f54269fe"

func TestSendTransaction(t *testing.T) {
	client, err := ethclient.Dial("http://localhost:36054")
	if err != nil {
		t.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// JSON string containing the key
	// Decrypt the key
	key, err := keystore.DecryptKey([]byte(sender), "423785")
	if err != nil {
		t.Fatalf("Failed to decrypt key: %v", err)
	}
	privateKey := key.PrivateKey
	accountAddress := key.Address

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
	toAddress := common.HexToAddress(receiver)
	value := big.NewInt(100000000000000000) // 1 ETH
	gasLimit := uint64(1000000)             // Adjusted gas limit
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

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
