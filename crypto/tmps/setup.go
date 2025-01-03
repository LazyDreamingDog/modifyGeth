package tmps

import (
	"crypto/rand"
	"math/big"

	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
)

// SetupCommon 用于生成公共部分
func SetupCommon(d int, N int) *CommonParams {
	// 随机生成 B1, B2
	b1, _ := rand.Int(rand.Reader, Prime)
	b2, _ := rand.Int(rand.Reader, Prime)

	// 生成 G_1, G_2, G_T
	g_1 := OneG1()
	g_2 := OneG2()
	G_T := bn256.Pair(g_1, g_2)

	// 生成 b_1, b_2
	b_1 := OneG1().ScalarMult(g_1, b1)
	b_2 := OneG1().ScalarMult(g_1, b2)

	return &CommonParams{
		G1:    g_1,
		G2:    g_2,
		GT:    G_T,
		B1:    b_1,
		B2:    b_2,
		B1Int: b1,
		B2Int: b2,
		N:     N,
		D:     d,
	}
}

func Setup(bf_M *Polynomial, com *CommonParams) (*Pk, *Ek) {
	W_x := new(Polynomial)

	W_x = NewPolynomial(GenerateRandomBigIntArray(Prime, com.D), Prime)
	//W_x = NewPolynomial([]*big.Int{big.NewInt(3), big.NewInt(3), big.NewInt(1), big.NewInt(2)}, Prime)
	Z_x := new(Polynomial)
	Z_x = bf_M.Subtract(W_x)

	//构造多项式X^2 + b0,X^2 + B1
	B1 := NewPolynomial([]*big.Int{big.NewInt(1), big.NewInt(0), com.B1Int}, Prime)
	B2 := NewPolynomial([]*big.Int{big.NewInt(1), big.NewInt(0), com.B2Int}, Prime)

	//使用欧几里得除法算法，将B_1,B_2进行扩展
	Q1, R1 := W_x.Divide(B1)
	Q2, R2 := Z_x.Divide(B2)

	r1_1 := new(big.Int)
	r1_0 := new(big.Int)
	r2_1 := new(big.Int)
	r2_0 := new(big.Int)
	if len(R1.coefficients) >= 2 {
		r1_1 = R1.coefficients[0]
		r1_0 = R1.coefficients[1]
	} else if len(R1.coefficients) == 1 {
		//r1是0
		r1_1 = new(big.Int).SetInt64(0)
		r1_0 = R1.coefficients[0]
	} else {
		//r1是0
		r1_1 = new(big.Int).SetInt64(0)
		r1_0 = new(big.Int).SetInt64(0)
	}

	if len(R2.coefficients) >= 2 {
		r2_1 = R2.coefficients[0]
		r2_0 = R2.coefficients[1]
	} else if len(R2.coefficients) == 1 {
		//r2是0
		r2_1 = new(big.Int).SetInt64(0)
		r2_0 = R2.coefficients[0]
	} else {
		//r2是0
		r2_1 = new(big.Int).SetInt64(0)
		r2_0 = new(big.Int).SetInt64(0)
	}

	r1 := new(big.Int).Add(r1_1, r2_1)
	r0 := new(big.Int).Add(r1_0, r2_0)
	r1.Mod(r1, Prime)
	r0.Mod(r0, Prime)

	q1_n := make([]*bn256.G2, com.D-1)
	q2_n := make([]*bn256.G2, com.D-1)

	r_0 := OneG2().ScalarMult(com.G2, r0)
	r_1 := OneG2().ScalarMult(com.G2, r1)
	for i := 0; i <= com.D-2; i++ {
		q1_n[i] = OneG2().ScalarMult(com.G2, Q1.coefficients[i])
	}
	for i := 0; i <= com.D-2; i++ {
		q2_n[i] = OneG2().ScalarMult(com.G2, Q2.coefficients[i])
	}

	pk := Pk{
		(*G1Alias)(com.G1),
		(*G2Alias)(com.G2),
		(*GTAlias)(com.GT),
		(*G1Alias)(com.B1),
		(*G1Alias)(com.B2),
		com.N,
		com.D,
		(*G2Alias)(r_0),
		(*G2Alias)(r_1),
	}

	ek := Ek{
		bf_M.coefficients,
		q1_n,
		q2_n,
	}
	return &pk, &ek
}
