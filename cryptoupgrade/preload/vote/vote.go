package vote

import (
	"bytes"
	"crypto/rand"
	"errors"
	"math"

	. "github.com/ethereum/go-ethereum/cryptoupgrade/preload/bls12381"
)

type BallotCipher []*SingleElemCipher

func (ballot_cipher BallotCipher) Mul(encryption_ballot BallotCipher) (BallotCipher, error) {
	group1 := NewG1()
	if len(ballot_cipher) != len(encryption_ballot) {
		return nil, errors.New("wrong ballot syntex")
	}
	for i := 0; i < len(ballot_cipher); i++ {
		ballot_cipher[i].C1 = group1.Add(group1.New(), ballot_cipher[i].C1, encryption_ballot[i].C1)
		ballot_cipher[i].C2 = group1.Add(group1.New(), ballot_cipher[i].C2, encryption_ballot[i].C2)
	}
	return ballot_cipher, nil
}

func (ballotcipher BallotCipher) Size() int {
	s := 0
	group1 := NewG1()
	s += len(group1.ToBytes(ballotcipher[0].C1)) * 2 * len(ballotcipher)
	return s
}

type CoefficientOfLine [2]*Fr

func GenerateBallotZKP(ballot *Ballot, g *PointG1, pk *PointG1, ck *PointG1) *BallotZKP {
	group1 := NewG1()
	virtual_n := ballot.virtual_n
	index := ballot.index
	len := uint64(math.Log2(float64(virtual_n)))
	zkp := NewBallotZKP()
	zkp.len = len
	var res_bytes bytes.Buffer
	alpha_vec := make([]*Fr, 1)
	beta_vec := make([]*Fr, 1)
	gamma_vec := make([]*Fr, 1)
	delta_vec := make([]*Fr, 1)
	index_bin_representation := uint64_bin_representation(index, len)
	for l := uint64(1); l <= len; l++ {
		alpha, _ := NewFr().Rand(rand.Reader)
		beta, _ := NewFr().Rand(rand.Reader)
		gamma, _ := NewFr().Rand(rand.Reader)
		delta, _ := NewFr().Rand(rand.Reader)
		alpha_vec = append(alpha_vec, alpha)
		beta_vec = append(beta_vec, beta)
		gamma_vec = append(gamma_vec, gamma)
		delta_vec = append(delta_vec, delta)
		index_at_l := index_bin_representation[l]
		I := group1.Add(group1.New(), group1.MulScalar(group1.New(), g, index_at_l), group1.MulScalar(group1.New(), ck, alpha))
		B := group1.Add(group1.New(), group1.MulScalar(group1.New(), g, beta), group1.MulScalar(group1.New(), ck, gamma))
		A := group1.Add(group1.New(), group1.MulScalar(group1.New(), g, NewFr().Mul(index_at_l, beta)), group1.MulScalar(group1.New(), ck, delta))
		zkp.I = append(zkp.I, I)
		zkp.B = append(zkp.B, B)
		zkp.A = append(zkp.A, A)
		res_bytes.Write(group1.ToBytes(I))
		res_bytes.Write(group1.ToBytes(B))
		res_bytes.Write(group1.ToBytes(A))
	}
	y := HashToFr(res_bytes.Bytes())

	res_bytes.Reset()
	Rl_vec := make([]*Fr, 0)
	for l := uint64(0); l < len; l++ {
		Rl, _ := NewFr().Rand(rand.Reader)
		Rl_vec = append(Rl_vec, Rl)

		temp_sum := NewFr().Zero()
		for j := uint64(0); j < virtual_n; j++ {
			coefficients := coefficient_of_pj_polynomials(j, index, len, beta_vec)
			temp_sum = NewFr().Add(temp_sum, NewFr().Mul(coefficients[l], fr_pow(y, j)))
		}

		_, D := encrypt_m_r(g, pk, temp_sum, Rl)
		zkp.D = append(zkp.D, D)
		res_bytes.Write(group1.ToBytes(D))
	}
	res_bytes.Write(y.ToBytes())
	x := HashToFr(res_bytes.Bytes())

	R := NewFr().Zero()
	for j := uint64(0); j < virtual_n; j++ {
		R = NewFr().Add(R, NewFr().Mul(NewFr().Mul(ballot.r[j], fr_pow(x, len)), fr_pow(y, j)))
	}
	for l := uint64(0); l < len; l++ {
		R = NewFr().Add(R, NewFr().Mul(Rl_vec[l], fr_pow(x, l)))
	}
	zkp.R = R

	for l := uint64(1); l <= len; l++ {
		z := NewFr().Add(NewFr().Mul(index_bin_representation[l], x), beta_vec[l])
		w := NewFr().Add(NewFr().Mul(alpha_vec[l], x), gamma_vec[l])
		v := NewFr().Add(NewFr().Mul(alpha_vec[l], NewFr().Sub(x, z)), delta_vec[l])
		zkp.z = append(zkp.z, z)
		zkp.w = append(zkp.w, w)
		zkp.v = append(zkp.v, v)
	}
	return zkp
}

func VerifyBallotZKPMain(ballot_cipher BallotCipher, zkp *BallotZKP, g *PointG1, pk *PointG1, ck *PointG1) bool {
	group1 := NewG1()
	len := zkp.len
	var res_bytes bytes.Buffer
	for i := uint64(1); i <= len; i++ {
		res_bytes.Write(group1.ToBytes(zkp.I[i]))
		res_bytes.Write(group1.ToBytes(zkp.B[i]))
		res_bytes.Write(group1.ToBytes(zkp.A[i]))
	}
	y := HashToFr(res_bytes.Bytes())

	res_bytes.Reset()
	for i := uint64(0); i < len; i++ {
		res_bytes.Write(group1.ToBytes(zkp.D[i]))
	}
	res_bytes.Write(y.ToBytes())
	x := HashToFr(res_bytes.Bytes())

	for i := uint64(1); i <= len; i++ {
		if !group1.Equal(group1.Add(group1.New(), group1.MulScalar(group1.New(), zkp.I[i], x), zkp.B[i]), group1.Add(group1.New(), group1.MulScalar(group1.New(), g, zkp.z[i]), group1.MulScalar(group1.New(), ck, zkp.w[i]))) {
			return false
		}
		if !group1.Equal(group1.Add(group1.New(), group1.MulScalar(group1.New(), zkp.I[i], NewFr().Sub(x, zkp.z[i])), zkp.A[i]), group1.Add(group1.New(), group1.MulScalar(group1.New(), g, NewFr().Zero()), group1.MulScalar(group1.New(), ck, zkp.v[i]))) {
			return false
		}
	}

	left_sum1 := group1.New().Zero()
	for j := uint64(0); j < uint64(math.Pow(2, float64(len))); j++ {
		j_bin_representation := uint64_bin_representation(j, len)
		mul_temp := NewFr().One()
		for l := uint64(1); l <= len; l++ {
			if j_bin_representation[l].Equal(NewFr().Zero()) {
				mul_temp = NewFr().Mul(mul_temp, NewFr().Sub(x, zkp.z[l]))
			} else {
				mul_temp = NewFr().Mul(mul_temp, zkp.z[l])
			}
		}
		mul_temp = NewFr().Sub(NewFr().Zero(), mul_temp)

		_, enc_temp := encrypt_m_r(g, pk, mul_temp, NewFr().Zero())

		temp := group1.Add(group1.New(), group1.MulScalar(group1.New(), ballot_cipher[j].C2, fr_pow(x, len)), enc_temp)

		left_sum1 = group1.Add(group1.New(), left_sum1, group1.MulScalar(group1.New(), temp, fr_pow(y, j)))
	}

	left_sum2 := group1.New().Zero()
	for l := uint64(0); l < len; l++ {
		left_sum2 = group1.Add(group1.New(), left_sum2, group1.MulScalar(group1.New(), zkp.D[l], fr_pow(x, l)))
	}

	_, right_c2 := encrypt_m_r(g, pk, NewFr().Zero(), zkp.R)
	return group1.Equal(group1.Add(group1.New(), left_sum1, left_sum2), right_c2)
}

