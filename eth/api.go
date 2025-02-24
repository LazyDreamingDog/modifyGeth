// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package eth

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
)

// EthereumAPI provides an API to access Ethereum full node-related information.
type EthereumAPI struct {
	e *Ethereum
	*BlockDataAPI
}

// NewEthereumAPI creates a new Ethereum protocol API for full nodes.
func NewEthereumAPI(e *Ethereum) *EthereumAPI {
	return &EthereumAPI{
		e:            e,
		BlockDataAPI: NewBlockDataAPI(e.APIBackend),
	}
}

// Etherbase is the address that mining rewards will be sent to.
func (api *EthereumAPI) Etherbase() (common.Address, error) {
	return api.e.Etherbase()
}

// Coinbase is the address that mining rewards will be sent to (alias for Etherbase).
func (api *EthereumAPI) Coinbase() (common.Address, error) {
	return api.Etherbase()
}

// Hashrate returns the POW hashrate.
func (api *EthereumAPI) Hashrate() hexutil.Uint64 {
	return hexutil.Uint64(api.e.Miner().Hashrate())
}

// Mining returns an indication if this node is currently mining.
func (api *EthereumAPI) Mining() bool {
	return api.e.IsMining()
}

// BlockDataAPI 提供区块数据查询的API
type BlockDataAPI struct {
	backend BlockDataReader
}

// NewBlockDataAPI 创建新的BlockDataAPI实例
func NewBlockDataAPI(backend BlockDataReader) *BlockDataAPI {
	return &BlockDataAPI{backend: backend}
}

// GetPowDifficulty 返回区块的PoW难度
func (api *BlockDataAPI) GetPowDifficulty(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*hexutil.Big, error) {
	difficulty, err := api.backend.GetPowDifficultyByNumberOrHash(ctx, blockNrOrHash)
	if err != nil {
		return nil, err
	}
	return (*hexutil.Big)(difficulty), nil
}

// GetPowGas 返回区块的PoW Gas
func (api *BlockDataAPI) GetPowGas(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*hexutil.Uint64, error) {
	gas, err := api.backend.GetPowGasByNumberOrHash(ctx, blockNrOrHash)
	if err != nil {
		return nil, err
	}
	result := hexutil.Uint64(gas)
	return &result, nil
}

// GetPowPrice 返回区块的PoW Price
func (api *BlockDataAPI) GetPowPrice(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*hexutil.Big, error) {
	price, err := api.backend.GetPowPriceByNumberOrHash(ctx, blockNrOrHash)
	if err != nil {
		return nil, err
	}
	return (*hexutil.Big)(price), nil
}

// GetAvgRatioNumerator 返回区块的AvgRatioNumerator
func (api *BlockDataAPI) GetAvgRatioNumerator(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*hexutil.Uint64, error) {
	num, err := api.backend.GetAvgRatioNumeratorByNumberOrHash(ctx, blockNrOrHash)
	if err != nil {
		return nil, err
	}
	result := hexutil.Uint64(num)
	return &result, nil
}

// GetAvgRatioDenominator 返回区块的AvgRatioDenominator
func (api *BlockDataAPI) GetAvgRatioDenominator(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*hexutil.Uint64, error) {
	den, err := api.backend.GetAvgRatioDenominatorByNumberOrHash(ctx, blockNrOrHash)
	if err != nil {
		return nil, err
	}
	result := hexutil.Uint64(den)
	return &result, nil
}

// GetAvgGasNumerator 返回区块的AvgGasNumerator
func (api *BlockDataAPI) GetAvgGasNumerator(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*hexutil.Uint64, error) {
	num, err := api.backend.GetAvgGasNumeratorByNumberOrHash(ctx, blockNrOrHash)
	if err != nil {
		return nil, err
	}
	result := hexutil.Uint64(num)
	return &result, nil
}

// GetAvgGasDenominator 返回区块的AvgGasDenominator
func (api *BlockDataAPI) GetAvgGasDenominator(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*hexutil.Uint64, error) {
	den, err := api.backend.GetAvgGasDenominatorByNumberOrHash(ctx, blockNrOrHash)
	if err != nil {
		return nil, err
	}
	result := hexutil.Uint64(den)
	return &result, nil
}

// GetPoSLeader 返回区块的PoSLeader
func (api *BlockDataAPI) GetPoSLeader(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (common.Address, error) {
	return api.backend.GetPoSLeaderByNumberOrHash(ctx, blockNrOrHash)
}

// GetPoSVoting 返回区块的PoSVoting
func (api *BlockDataAPI) GetPoSVoting(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (hexutil.Bytes, error) {
	voting, err := api.backend.GetPoSVotingByNumberOrHash(ctx, blockNrOrHash)
	if err != nil {
		return nil, err
	}
	return voting, nil
}

// GetCommitTxLength 返回区块的CommitTxLength
func (api *BlockDataAPI) GetCommitTxLength(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*hexutil.Uint64, error) {
	length, err := api.backend.GetCommitTxLengthByNumberOrHash(ctx, blockNrOrHash)
	if err != nil {
		return nil, err
	}
	result := hexutil.Uint64(length)
	return &result, nil
}

// GetIncentive 返回区块的Incentive
func (api *BlockDataAPI) GetIncentive(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*hexutil.Big, error) {
	incentive, err := api.backend.GetIncentiveByNumberOrHash(ctx, blockNrOrHash)
	if err != nil {
		return nil, err
	}
	return (*hexutil.Big)(incentive.ToBig()), nil
}
