package vote

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"errors"

	. "github.com/ethereum/go-ethereum/cryptoupgrade/preload/bls12381"
)

type Ballot struct {
	n         uint64
	virtual_n uint64
	index     uint64
	e         []*Fr
	r         []*Fr
}

func GenerateBallot(n uint64, index uint64) (*Ballot, error) {
	ballot := &Ballot{}
	ballot.n = n
	ballot.index = index
	if n < 2 {
		return nil, errors.New("wrong candidate number")
	}
	if index >= n {
		return nil, errors.New("wrong index")
	}
	virtual_n := uint64(1)
	for virtual_n < n {
		virtual_n *= 2
	}
	ballot.virtual_n = virtual_n
	for i := uint64(0); i < n; i++ {
		if i == index {
			ballot.e = append(ballot.e, NewFr().One())
		} else {
			ballot.e = append(ballot.e, NewFr().Zero())
		}
		r, _ := NewFr().Rand(rand.Reader)
		ballot.r = append(ballot.r, r)
	}
	for i := uint64(n); i < virtual_n; i++ {
		ballot.e = append(ballot.e, NewFr().Zero())
		ballot.r = append(ballot.r, NewFr().Zero())
	}
	return ballot, nil
}

func EncryptBallot(ballot *Ballot, g *PointG1, pk *PointG1) BallotCipher {
	var res BallotCipher
	virtual_n := ballot.virtual_n
	for i := uint64(0); i < virtual_n; i++ {
		cipher := &SingleElemCipher{}
		cipher.C1, cipher.C2 = encrypt_m_r(g, pk, ballot.e[i], ballot.r[i])
		res = append(res, cipher)
	}
	return res
}

type SingleElemCipher struct {
	C1 *PointG1
	C2 *PointG1
}

func NewSingleElemCipher() *SingleElemCipher {
	return &SingleElemCipher{}
}

func (cipher *SingleElemCipher) Serialize() []byte {
	var buf bytes.Buffer
	group1 := NewG1()
	buf.Write(group1.ToBytes(cipher.C1))
	buf.Write(group1.ToBytes(cipher.C2))
	return buf.Bytes()
}

func (cipher *SingleElemCipher) Deserialize(data []byte) (*SingleElemCipher, error) {
	group1 := NewG1()
	if len(data) != 96*2 {
		return nil, errors.New("invalid data")
	}
	cipher.C1, _ = group1.FromBytes(data[:96])
	cipher.C2, _ = group1.FromBytes(data[96:])
	return cipher, nil
}

func (cipher *SingleElemCipher) Equal(r_cipher *SingleElemCipher) bool {
	group1 := NewG1()
	return group1.Equal(cipher.C1, r_cipher.C1) && group1.Equal(cipher.C2, r_cipher.C2)
}

func (cipher *SingleElemCipher) Add(r_cipher *SingleElemCipher) {
	group1 := NewG1()
	cipher.C1 = group1.Add(group1.New(), cipher.C1, r_cipher.C1)
	cipher.C2 = group1.Add(group1.New(), cipher.C2, r_cipher.C2)
}

type BallotZKP struct {
	len uint64
	I   []*PointG1
	B   []*PointG1
	A   []*PointG1
	D   []*PointG1
	z   []*Fr
	w   []*Fr
	v   []*Fr
	R   *Fr
}

func NewBallotZKP() *BallotZKP {
	return &BallotZKP{
		len: 0,
		I:   make([]*PointG1, 1),
		B:   make([]*PointG1, 1),
		A:   make([]*PointG1, 1),
		D:   make([]*PointG1, 0),
		z:   make([]*Fr, 1),
		w:   make([]*Fr, 1),
		v:   make([]*Fr, 1),
		R:   NewFr(),
	}
}

func (ballot_zkp *BallotZKP) Serialize() []byte {
	group1 := NewG1()
	var data bytes.Buffer
	uintbytes := make([]byte, 8)
	binary.BigEndian.PutUint64(uintbytes, ballot_zkp.len)
	data.Write(uintbytes)
	for i := 1; i < len(ballot_zkp.I); i++ {
		data.Write(group1.ToBytes(ballot_zkp.I[i]))
		data.Write(group1.ToBytes(ballot_zkp.B[i]))
		data.Write(group1.ToBytes(ballot_zkp.A[i]))
		data.Write(group1.ToBytes(ballot_zkp.D[i-1]))
		data.Write(ballot_zkp.z[i].ToBytes())
		data.Write(ballot_zkp.w[i].ToBytes())
		data.Write(ballot_zkp.v[i].ToBytes())
	}
	data.Write(ballot_zkp.R.ToBytes())
	return data.Bytes()
}

func (ballot_zkp *BallotZKP) Deserialize(data []byte) (*BallotZKP, error) {
	ballot_zkp = NewBallotZKP()
	group1 := NewG1()
	ballot_zkp.len = binary.BigEndian.Uint64(data[:8])
	buf := bytes.NewBuffer(data[8:])
	for i := 0; i < int(ballot_zkp.len); i++ {
		read_point_data := make([]byte, 96)
		buf.Read(read_point_data)
		point, _ := group1.FromBytes(read_point_data)
		ballot_zkp.I = append(ballot_zkp.I, point)

		buf.Read(read_point_data)
		point, _ = group1.FromBytes(read_point_data)
		ballot_zkp.B = append(ballot_zkp.B, point)

		buf.Read(read_point_data)
		point, _ = group1.FromBytes(read_point_data)
		ballot_zkp.A = append(ballot_zkp.A, point)

		buf.Read(read_point_data)
		point, _ = group1.FromBytes(read_point_data)
		ballot_zkp.D = append(ballot_zkp.D, point)

		read_fr_data := make([]byte, 32)
		buf.Read(read_fr_data)
		fr := NewFr().FromBytes(read_fr_data)
		ballot_zkp.z = append(ballot_zkp.z, fr)

		buf.Read(read_fr_data)
		fr = NewFr().FromBytes(read_fr_data)
		ballot_zkp.w = append(ballot_zkp.w, fr)

		buf.Read(read_fr_data)
		fr = NewFr().FromBytes(read_fr_data)
		ballot_zkp.v = append(ballot_zkp.v, fr)
	}
	read_fr_data := make([]byte, 32)
	buf.Read(read_fr_data)
	fr := NewFr().FromBytes(read_fr_data)
	ballot_zkp.R = fr
	return ballot_zkp, nil
}

func (ballot_zkp *BallotZKP) Equal(r_zkp *BallotZKP) bool {
	group1 := NewG1()
	if ballot_zkp.len != r_zkp.len {
		return false
	}
	for i := 0; i < int(ballot_zkp.len); i++ {
		if !group1.Equal(ballot_zkp.I[i+1], r_zkp.I[i+1]) {
			return false
		}
		if !group1.Equal(ballot_zkp.B[i+1], r_zkp.B[i+1]) {
			return false
		}
		if !group1.Equal(ballot_zkp.A[i+1], r_zkp.A[i+1]) {
			return false
		}
		if !group1.Equal(ballot_zkp.D[i], r_zkp.D[i]) {
			return false
		}
		if !ballot_zkp.z[i+1].Equal(r_zkp.z[i+1]) {
			return false
		}
		if !ballot_zkp.w[i+1].Equal(r_zkp.w[i+1]) {
			return false
		}
		if !ballot_zkp.v[i+1].Equal(r_zkp.v[i+1]) {
			return false
		}
	}
	return ballot_zkp.R.Equal(r_zkp.R)
}

func (ballpt_zkp *BallotZKP) Size() int {
	s := 0
	group1 := NewG1()
	s += 8
	s += len(group1.ToBytes(ballpt_zkp.I[1]))*(len(ballpt_zkp.I)-1)*3 + len(group1.ToBytes(ballpt_zkp.D[0]))*len(ballpt_zkp.D) + len(ballpt_zkp.v[1].ToBytes())*((len(ballpt_zkp.v)-1)*3+1)
	return s
}
