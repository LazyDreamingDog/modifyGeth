package vm

import (
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto/tmps"
	"github.com/holiman/uint256"
)

var errInputLength = errors.New("error: input length must be valid JSON")

// 定义预编译合约结构体
type tmpsVerify struct{}

// 定义 Gas 消耗计算方法
func (d *tmpsVerify) RequiredGas(input []byte) uint64 {
	return 50000 // 假设固定的 Gas 消耗
}

// Run 方法：解析输入并调用 Verify
func (d *tmpsVerify) Run(input []byte, blkCtx BlockContext) ([]byte, error) {
	// fmt.Println("input:", input)
	// 调用 decode 函数解析输入
	pk, c, pi, vk, err := tmps.Decode(input)
	if err != nil {
		return nil, fmt.Errorf("decode error: %w", err)
	}

	// 调用 Verify 函数验证
	isValid := tmps.Verify(pk, c, pi, vk)

	// 将结果转换为字节数组返回
	result := uint256.NewInt(0)
	if isValid {
		result.SetUint64(1)
	}
	return result.Bytes(), nil
}
