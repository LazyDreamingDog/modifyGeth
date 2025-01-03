package tmps

import (
	"crypto/rand"
	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
	"math/big"
)

// generateRandomSubset 使用crypto/rand生成一个大小为l的{1,2,...,N}的随机子集
func generateRandomSubset(N int, l int) ([]int, error) {
	numbers := make([]int, N)
	for i := range numbers {
		numbers[i] = i + 1
	}

	subset := make([]int, 0, l)
	for len(subset) < l {
		// 使用crypto/rand生成安全的随机索引
		idxBig, err := rand.Int(rand.Reader, big.NewInt(int64(len(numbers))))
		if err != nil {
			return nil, err // 处理可能的错误
		}
		index := int(idxBig.Int64())

		subset = append(subset, numbers[index])
		// 移除已选择的元素
		numbers[index], numbers[len(numbers)-1] = numbers[len(numbers)-1], numbers[index]
		numbers = numbers[:len(numbers)-1]
	}

	return subset, nil
}

func Chal(pk *Pk, chal []*big.Int) (*big.Int, *Vk) {
	tmp := OneG2().ScalarMult((*bn256.G2)(pk.R_1), chal[0])
	vk_n := OneG2().Add(tmp, (*bn256.G2)(pk.R_0))

	vk := Vk{
		(*BigIntAlias)(chal[0]),
		(*G2Alias)(vk_n),
	}

	return chal[0], &vk
}
