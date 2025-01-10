package thresholdencrypt

import (
	"bytes"
	"crypto/rand"

	. "github.com/ethereum/go-ethereum/cryptoupgrade/preload/bls12381"
	"github.com/ethereum/go-ethereum/cryptoupgrade/preload/vote"
)

func (p *CommiteeClient) GenerateDecryptionShare(x_index []uint64, cipher *vote.SingleElemCipher) *PointG1 {
	group1 := NewG1()
	w_id := lagrange_id_at_zero(p.id, x_index)
	m_id := group1.MulScalar(group1.New(), cipher.C1, NewFr().Sub(NewFr().Zero(), NewFr().Mul(w_id, p.secret_share)))
	return m_id
}

func (p *CommiteeClient) GenerateDecryptionShareZKP(cipher *vote.SingleElemCipher, g *PointG1, x_index []uint64) *DecryptionShareZKP {
	group1 := NewG1()
	zkp := NewDecryptionShareZKP()
	w_id := lagrange_id_at_zero(p.id, x_index)

	var hash_bytes bytes.Buffer

	r, _ := NewFr().Rand(rand.Reader)
	zkp.t1 = group1.MulScalar(group1.New(), g, r)
	zkp.t2 = group1.MulScalar(group1.New(), group1.MulScalar(group1.New(), cipher.C1, NewFr().Sub(NewFr().Zero(), w_id)), r)

	// hash_bytes.Write(group1.ToBytes(group1.MulScalar(group1.New(), p.cipher.C1, NewFr().Sub(NewFr().Zero(), w_id))))
	// hash_bytes.Write(group1.ToBytes(p.decryption_share))
	hash_bytes.Write(group1.ToBytes(zkp.t1))
	hash_bytes.Write(group1.ToBytes(zkp.t2))
	hash_bytes.Write(group1.ToBytes(p.pk))
	e := HashToFr(hash_bytes.Bytes())
	zkp.z = NewFr().Add(r, NewFr().Mul(e, p.secret_share))
	return zkp
}

func VerifyDecryptionShareMain(node_id uint64, x_index []uint64, cipher *vote.SingleElemCipher, decryption_share *PointG1, zkp *DecryptionShareZKP, commitee_pk, g *PointG1) bool {
	group1 := NewG1()
	var hash_bytes bytes.Buffer

	hash_bytes.Write(group1.ToBytes(zkp.t1))
	hash_bytes.Write(group1.ToBytes(zkp.t2))
	e := HashToFr(hash_bytes.Bytes())
	if !group1.Equal(group1.MulScalar(group1.New(), g, zkp.z), group1.Add(group1.New(), zkp.t1, group1.MulScalar(group1.New(), commitee_pk, e))) {
		return false
	}
	w_id := lagrange_id_at_zero(node_id, x_index)
	return group1.Equal(
		group1.MulScalar(group1.New(), group1.MulScalar(group1.New(), cipher.C1, NewFr().Sub(NewFr().Zero(), w_id)), zkp.z),
		group1.Add(group1.New(), zkp.t2, group1.MulScalar(group1.New(), decryption_share, e)))
}
