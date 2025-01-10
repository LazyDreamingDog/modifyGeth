package thresholdencrypt

import (
	"bytes"
	"crypto/rand"

	. "github.com/ethereum/go-ethereum/cryptoupgrade/preload/bls12381"
)

type DecryptionShareZKP struct {
	t1 *PointG1
	t2 *PointG1
	z  *Fr
}

func (zkp *DecryptionShareZKP) Serialize() []byte {
	group1 := NewG1()
	buf := bytes.NewBuffer([]byte{})
	buf.Write(group1.ToBytes(zkp.t1))
	buf.Write(group1.ToBytes(zkp.t2))
	buf.Write(zkp.z.ToBytes())
	return buf.Bytes()
}

func (zkp *DecryptionShareZKP) Deserialize([]byte) (*DecryptionShareZKP, error) {
	group1 := NewG1()
	buf := bytes.NewBuffer([]byte{})

	read_point_data := make([]byte, 96)
	buf.Read(read_point_data)
	point, _ := group1.FromBytes(read_point_data)
	zkp.t1 = point

	buf.Read(read_point_data)
	point, _ = group1.FromBytes(read_point_data)
	zkp.t2 = point

	read_fr_data := make([]byte, 32)
	buf.Read(read_fr_data)
	zkp.z = NewFr().FromBytes(read_fr_data)

	return zkp, nil
}

func NewDecryptionShareZKP() *DecryptionShareZKP {
	return &DecryptionShareZKP{}
}

type Param struct {
	node_number uint64
	threshold   uint64
	g           *PointG1
}

type CommiteeMember struct {
	id                    uint64
	pk                    *PointG1
	polynomial_commitment []*PointG1
}

type CommiteeClient struct {
	CommiteeMember
	sk            *Fr
	polynomial_fr PolynomialFr
	secret_share  *Fr
}

type PolynomialFr []*Fr

func (poly PolynomialFr) compute(x *Fr) *Fr {
	res := NewFr().Zero()
	power_res := NewFr().One()
	for i := 0; i < len(poly); i++ {
		res = NewFr().Add(res, NewFr().Mul(power_res, poly[i]))
		power_res = NewFr().Mul(power_res, x)
	}
	return res
}

func (p *CommiteeClient) Initialize(id uint64, param *Param) {
	group1 := NewG1()
	p.id = id
	for i := 0; i < int(param.threshold); i++ {
		r, _ := NewFr().Rand(rand.Reader)
		p.polynomial_fr = append(p.polynomial_fr, r)
		p.polynomial_commitment = append(p.polynomial_commitment, group1.MulScalar(group1.New(), param.g, r))
	}
}
