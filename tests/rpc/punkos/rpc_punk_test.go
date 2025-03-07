package punkos

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"math/big"
	"net/http"
	"testing"
)

type RPCRequest struct {
	Jsonrpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	ID      int           `json:"id"`
}

type RPCResponse struct {
	Jsonrpc string          `json:"jsonrpc"`
	ID      int             `json:"id"`
	Result  json.RawMessage `json:"result"`
	Error   *RPCError       `json:"error"`
}

type RPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func sendRPCRequest(url, method string, params []interface{}) (*RPCResponse, error) {
	reqBody := RPCRequest{
		Jsonrpc: "2.0",
		Method:  method,
		Params:  params,
		ID:      1,
	}
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var rpcResp RPCResponse
	err = json.Unmarshal(respBytes, &rpcResp)
	if err != nil {
		return nil, err
	}

	if rpcResp.Error != nil {
		return nil, errors.New(rpcResp.Error.Message)
	}

	return &rpcResp, nil
}

var block = "latest"

func TestGetBalance(t *testing.T) {
	address := sender.Address.Hex()
	rpcResp, err := sendRPCRequest(rpcUrl, "eth_getBalance", []interface{}{address, block})
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}

	var balance string
	err = json.Unmarshal(rpcResp.Result, &balance)
	if err != nil {
		t.Fatalf("Failed to unmarshal result: %v", err)
	}

	balanceInt := new(big.Int)
	balanceInt.SetString(balance[2:], 16)

	t.Logf("Balance: %s", balanceInt)
}

// 测试获取利息信息
func TestGetInterest(t *testing.T) {
	address := sender.Address.Hex()
	rpcResp, err := sendRPCRequest(rpcUrl, "eth_getInterest", []interface{}{address, block})
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}

	var stateDb string
	err = json.Unmarshal(rpcResp.Result, &stateDb)
	if err != nil {
		t.Fatalf("Failed to unmarshal result: %v", err)
	}

	t.Logf("StateDb: %s", stateDb)
}

func TestGetPledgeInfo(t *testing.T) {
	address := "0x153af978Feff7bb0af214573dC88383578b3548f"
	rpcResp, err := sendRPCRequest(rpcUrl, "eth_getPledgeInfo", []interface{}{address, block})
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}

	var stateDb string
	err = json.Unmarshal(rpcResp.Result, &stateDb)
	if err != nil {
		t.Fatalf("Failed to unmarshal result: %v", err)
	}

	t.Logf("StateDb: %s", stateDb)
}

// 测试获取区块链的Pow难度
func TestGetPowDifficulty(t *testing.T) {
	blockNumber := "latest"
	rpcResp, err := sendRPCRequest(rpcUrl, "eth_getPowDifficulty", []interface{}{blockNumber})
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}

	var pd string
	err = json.Unmarshal(rpcResp.Result, &pd)
	if err != nil {
		t.Fatalf("Failed to unmarshal result: %v", err)
	}

	t.Logf("PowDifficulty: %s", pd)
}
