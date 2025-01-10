package vote

import (
	. "github.com/ethereum/go-ethereum/cryptoupgrade/preload/bls12381"
)

func fr_pow(x *Fr, k uint64) *Fr {
	res := NewFr().One()
	for i := uint64(0); i < k; i++ {
		res = NewFr().Mul(res, x)
	}
	return res
}

func coefficient_of_pj_polynomials(j, index, len uint64, beta_vec []*Fr) []*Fr {
	index_bin_representation := uint64_bin_representation(index, len)
	j_bin_representation := uint64_bin_representation(j, len)
	polynomials := make([]CoefficientOfLine, len)
	for l := uint64(1); l <= len; l++ {
		if j_bin_representation[l].Equal(NewFr().One()) {
			polynomials[l-1][0] = index_bin_representation[l]
			polynomials[l-1][1] = beta_vec[l]
		} else {
			polynomials[l-1][0] = NewFr().Sub(NewFr().One(), index_bin_representation[l])
			polynomials[l-1][1] = NewFr().Sub(NewFr().Zero(), beta_vec[l])
		}
	}
	return coefficient_of_mul_polynomials(polynomials)
}

func encrypt_m_r(g *PointG1, pk *PointG1, m *Fr, r *Fr) (c1 *PointG1, c2 *PointG1) {
	group1 := NewG1()
	// C1 = g ^ r
	c1 = group1.MulScalar(group1.New(), g, r)
	// C2 = mPoint * y ^ r
	c2 = group1.Add(group1.New(), group1.MulScalar(group1.New(), g, m), group1.MulScalar(group1.New(), pk, r))
	return c1, c2
}

func coefficient_of_mul_polynomials(polynomials []CoefficientOfLine) []*Fr {
	n := len(polynomials)
	f := make([][]*Fr, n+1)
	for i := range f {
		f[i] = make([]*Fr, n+1)
	}
	// 边界条件
	f[0][0] = NewFr().One()

	// 递推计算
	for i := 1; i <= n; i++ {

		a, b := polynomials[i-1][0], polynomials[i-1][1] // 第 i 个多项式的系数
		for j := 0; j <= i; j++ {
			if f[i-1][j] == nil {
				f[i-1][j] = NewFr().Zero()
			}
			f[i][j] = NewFr().Mul(f[i-1][j], b) // 不选 x 项
			if j > 0 {
				f[i][j] = NewFr().Add(f[i][j], NewFr().Mul(f[i-1][j-1], a)) // 选 x 项
			}
		}
	}
	return f[n]
}

func uint64_bin_representation(k, len uint64) []*Fr {
	bin_representation := make([]*Fr, len+1)
	for i := uint64(1); i <= len; i++ {
		bin_representation[i] = FrFromInt(int((k >> (len - i)) & 1))
	}
	return bin_representation
}
