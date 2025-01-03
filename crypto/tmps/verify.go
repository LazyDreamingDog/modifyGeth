package tmps

import (
	"math/big"

	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
)

func isEqual(gt1 *bn256.GT, gt2 *bn256.GT) bool {
	return gt1.String() == gt2.String()
}

func Verify(pk *Pk, c *big.Int, pi *Pi, vk *Vk) bool {
	left := bn256.Pair((*bn256.G1)(pk.G_1), new(bn256.G2).ScalarMult((*bn256.G2)(pk.G_2), (*big.Int)(pi.Y)))
	PrintGT(left, "left")

	right1 := bn256.Pair(new(bn256.G1).Add((*bn256.G1)(pk.B1), new(bn256.G1).ScalarMult((*bn256.G1)(pk.G_1), new(big.Int).Mul(c, c))), (*bn256.G2)(pi.Pi1))
	right2 := bn256.Pair(new(bn256.G1).Add((*bn256.G1)(pk.B2), new(bn256.G1).ScalarMult((*bn256.G1)(pk.G_1), new(big.Int).Mul(c, c))), (*bn256.G2)(pi.Pi2))
	right3 := bn256.Pair((*bn256.G1)(pk.G_1), (*bn256.G2)(vk.Vk))
	PrintGT(right1, "right1")
	PrintGT(right2, "right2")
	PrintGT(right3, "right3")

	temp := new(bn256.GT).Add(right1, right2)
	right := new(bn256.GT).Add(temp, right3)
	PrintGT(right, "right")

	return isEqual(left, right)
}
