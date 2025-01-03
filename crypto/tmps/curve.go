package tmps

import (
	"math/big"

	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
)

func OneG1() *bn256.G1 {
	return new(bn256.G1).ScalarBaseMult(new(big.Int).SetInt64(1))
}

func OneG2() *bn256.G2 {
	return new(bn256.G2).ScalarBaseMult(new(big.Int).SetInt64(1))
}

func ZeroG1() *bn256.G1 {
	return new(bn256.G1).ScalarBaseMult(new(big.Int).SetInt64(0))
}

func ZeroG2() *bn256.G2 {
	return new(bn256.G2).ScalarBaseMult(new(big.Int).SetInt64(0))
}

func ZeroGT() *bn256.GT {
	return bn256.Pair(ZeroG1(), ZeroG2())
}
