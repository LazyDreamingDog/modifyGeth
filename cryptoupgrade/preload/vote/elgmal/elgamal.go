package ec_elgamal

import (
	. "github.com/ethereum/go-ethereum/cryptoupgrade/preload/bls12381"
	"crypto/rand"
	"math/big"
)

// TODO

var (

	// modulus of Fr
	modulus, _ = new(big.Int).SetString("0x30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001", 16)
)

type Proof struct {
	// generator
	g *PointG1
	// public key
	PK *PointG1
	// cipher
	C1 *PointG1
	C2 *PointG1
	// proof
	CM1 *PointG1
	CM2 *PointG1
	T1  *PointG1
	T2  *Fr
}

func Keygen(g *PointG1) (sk *Fr, pk *PointG1) {
	group1 := NewG1()
	sk, _ = NewFr().Rand(rand.Reader)
	pk = group1.MulScalar(group1.New(), g, sk)
	return sk, pk
}

func Encrypt(g *PointG1, pk *PointG1, mPoint *PointG1) (c1 *PointG1, c2 *PointG1) {
	group1 := NewG1()
	// 生成随机数 r
	r, _ := NewFr().Rand(rand.Reader)
	// C1 = g ^ r
	c1 = group1.MulScalar(group1.New(), g, r)
	// C2 = mPoint * y ^ r
	c2 = group1.Add(group1.New(), mPoint, group1.MulScalar(group1.New(), pk, r))
	return c1, c2
}

func EncryptWithRandAndProve(g *PointG1, pk *PointG1, mPoint *PointG1) *Proof {
	// 生成随机数 r
	r, _ := NewFr().Rand(rand.Reader)
	return EncryptAndProve(g, pk, mPoint, r)
}

func EncryptAndProve(g *PointG1, pk *PointG1, mPoint *PointG1, r *Fr) *Proof {
	group1 := NewG1()
	// C1 = group1 ^ r
	c1 := group1.MulScalar(group1.New(), g, r)
	// C2 = mPoint * y ^ r
	c2 := group1.Add(group1.New(), mPoint, group1.MulScalar(group1.New(), pk, r))
	// 选取随机点 S1
	tmp, _ := NewFr().Rand(rand.Reader)
	S1 := group1.MulScalar(group1.New(), g, tmp)
	// 选取随机数 s2
	s2, _ := NewFr().Rand(rand.Reader)
	// 计算承诺 CM1 = g^s2
	cm1 := group1.MulScalar(group1.New(), g, s2)
	// 计算承诺 CM2 = S1 · y^s2
	cm2 := group1.Add(group1.New(), S1, group1.MulScalar(group1.New(), pk, s2))
	// 计算哈希值 e = Hash(y, C1, C2, CM1, CM2)
	hashInput := group1.ToBytes(pk)
	hashInput = append(hashInput, group1.ToBytes(c1)...)
	hashInput = append(hashInput, group1.ToBytes(c2)...)
	hashInput = append(hashInput, group1.ToBytes(cm1)...)
	hashInput = append(hashInput, group1.ToBytes(cm2)...)
	e := HashToFr(hashInput)
	// 计算响应 T1 = M^e · S1
	t1 := group1.Add(group1.New(), group1.MulScalar(group1.New(), mPoint, e), S1)
	// 计算响应 T2 = er + s2
	t2 := NewFr().Mul(e, r)
	t2.Add(t2, s2)
	return &Proof{
		g:   g,
		PK:  pk,
		C1:  c1,
		C2:  c2,
		CM1: cm1,
		CM2: cm2,
		T1:  t1,
		T2:  t2,
	}
}

func Verify(proof *Proof) bool {
	group1 := NewG1()
	// 计算哈希值 e = Hash(y, C1, C2, CM1, CM2)
	hashInput := group1.ToBytes(proof.PK)
	hashInput = append(hashInput, group1.ToBytes(proof.C1)...)
	hashInput = append(hashInput, group1.ToBytes(proof.C2)...)
	hashInput = append(hashInput, group1.ToBytes(proof.CM1)...)
	hashInput = append(hashInput, group1.ToBytes(proof.CM2)...)
	e := HashToFr(hashInput)
	// 验证 C1^e · CM1 = g^T2
	left1 := group1.Add(group1.New(), group1.MulScalar(group1.New(), proof.C1, e), proof.CM1)
	right1 := group1.MulScalar(group1.New(), proof.g, proof.T2)
	if !group1.Equal(left1, right1) {
		return false
	}
	// 验证 C2^e · CM2 = T1 · y^T2
	left2 := group1.Add(group1.New(), group1.MulScalar(group1.New(), proof.C2, e), proof.CM2)
	right2 := group1.Add(group1.New(), proof.T1, group1.MulScalar(group1.New(), proof.PK, proof.T2))
	return group1.Equal(left2, right2)
}

func Decrypt(sk *Fr, c1 *PointG1, c2 *PointG1) (m *PointG1) {
	group1 := NewG1()
	// c2 / c1^x
	return group1.Sub(group1.New(), c2, group1.MulScalar(group1.New(), c1, sk))
}
