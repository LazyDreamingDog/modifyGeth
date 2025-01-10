package ec_elgamal

import (
	"crypto/rand"
	"testing"

	"github.com/ethereum/go-ethereum/cryptoupgrade/preload/bls12381"
	"github.com/stretchr/testify/assert"
)

func TestDecrypt(t *testing.T) {
	group1 := bls12381.NewG1()
	g := group1.One()
	// modular/(2^33)
	sk, pk := Keygen(g)
	m, _ := bls12381.NewFr().Rand(rand.Reader)
	mPoint := group1.MulScalar(group1.New(), group1.One(), m)
	c1, c2 := Encrypt(g, pk, mPoint)
	mPointRecover := Decrypt(sk, c1, c2)
	assert.True(t, group1.Equal(mPoint, mPointRecover))
}

func TestProof(t *testing.T) {
	group1 := bls12381.NewG1()
	g := group1.One()
	sk, pk := Keygen(g)
	m, _ := bls12381.NewFr().Rand(rand.Reader)
	mPoint := group1.MulScalar(group1.New(), group1.One(), m)
	proof := EncryptWithRandAndProve(g, pk, mPoint)
	assert.True(t, Verify(proof))
	mPointRecover := Decrypt(sk, proof.C1, proof.C2)
	assert.True(t, group1.Equal(mPoint, mPointRecover))
}
