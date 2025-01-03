package tmps

import (
	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
	"math/big"
)

func Aggregation(pis []*Pi, vks []*Vk) (*Pi, *Vk) {
	//生成pis长度的随机系数
	r := GenerateRandomBigIntArray(Prime, len(pis))
	y := new(big.Int)
	pi1 := new(bn256.G2)
	pi2 := new(bn256.G2)
	vk_a := new(bn256.G2)

	for i := 0; i < len(pis); i++ {
		y.Add(y, new(big.Int).Mul(r[i], (*big.Int)(pis[i].Y)))
		pi1.Add(pi1, new(bn256.G2).ScalarMult((*bn256.G2)(pis[i].Pi1), r[i]))
		pi2.Add(pi2, new(bn256.G2).ScalarMult((*bn256.G2)(pis[i].Pi2), r[i]))
		vk_a.Add(vk_a, new(bn256.G2).ScalarMult((*bn256.G2)(vks[i].Vk), r[i]))
	}
	pi := Pi{
		Y:   (*BigIntAlias)(y),
		Pi1: (*G2Alias)(pi1),
		Pi2: (*G2Alias)(pi2),
	}

	vk1 := Vk{Challenge: vks[0].Challenge,
		Vk: (*G2Alias)(vk_a),
	}
	return &pi, &vk1
}
