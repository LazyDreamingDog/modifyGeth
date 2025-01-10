package bls12381

import (
	"errors"
	"math/big"
)

func FromBytes(in []byte) (*Fe, error) {
	fe := &Fe{}
	if len(in) != fpByteSize {
		return nil, errors.New("input string must be equal 48 Bytes")
	}
	fe.setBytes(in)
	if !fe.isValid() {
		return nil, errors.New("must be less than modulus")
	}
	toMont(fe, fe)
	return fe, nil
}

func from64Bytes(in []byte) (*Fe, error) {
	if len(in) != 32*2 {
		return nil, errors.New("input string must be equal 64 Bytes")
	}
	a0 := make([]byte, fpByteSize)
	copy(a0[fpByteSize-32:fpByteSize], in[:32])
	a1 := make([]byte, fpByteSize)
	copy(a1[fpByteSize-32:fpByteSize], in[32:])
	e0, err := FromBytes(a0)
	if err != nil {
		return nil, err
	}
	e1, err := FromBytes(a1)
	if err != nil {
		return nil, err
	}
	// F = 2 ^ 256 * R
	F := Fe{
		0x75b3cd7c5ce820f,
		0x3ec6ba621c3edb0b,
		0x168a13d82bff6bce,
		0x87663c4bf8c449d2,
		0x15f34c83ddc8d830,
		0xf9628b49caa2e85,
	}

	mul(e0, e0, &F)
	add(e1, e1, e0)
	return e1, nil
}

func FromBig(in *big.Int) (*Fe, error) {
	fe := new(Fe).SetBig(in)
	if !fe.isValid() {
		return nil, errors.New("invalid input string")
	}
	toMont(fe, fe)
	return fe, nil
}

func fromString(in string) (*Fe, error) {
	fe, err := new(Fe).setString(in)
	if err != nil {
		return nil, err
	}
	if !fe.isValid() {
		return nil, errors.New("invalid input string")
	}
	toMont(fe, fe)
	return fe, nil
}

func ToBytes(e *Fe) []byte {
	e2 := new(Fe)
	fromMont(e2, e)
	return e2.Bytes()
}

func ToBig(e *Fe) *big.Int {
	e2 := new(Fe)
	fromMont(e2, e)
	return e2.big()
}

func toString(e *Fe) (s string) {
	e2 := new(Fe)
	fromMont(e2, e)
	return e2.string()
}

func toMont(c, a *Fe) {
	mul(c, a, r2)
}

func fromMont(c, a *Fe) {
	mul(c, a, &Fe{1})
}

func wfp2MulGeneric(c *wfe2, a, b *Fe2) {
	wt0, wt1 := new(wfe), new(wfe)
	t0, t1 := new(Fe), new(Fe)
	wmul(wt0, &a[0], &b[0])
	wmul(wt1, &a[1], &b[1])
	wsub(&c[0], wt0, wt1)
	lwaddAssign(wt0, wt1)
	ladd(t0, &a[0], &a[1])
	ladd(t1, &b[0], &b[1])
	wmul(wt1, t0, t1)
	lwsub(&c[1], wt1, wt0)
}

func wfp2SquareGeneric(c *wfe2, a *Fe2) {
	t0, t1, t2 := new(Fe), new(Fe), new(Fe)
	ladd(t0, &a[0], &a[1])
	sub(t1, &a[0], &a[1])
	ldouble(t2, &a[0])
	wmul(&c[0], t1, t0)
	wmul(&c[1], t2, &a[1])
}

func exp(c, a *Fe, e *big.Int) {
	z := new(Fe).Set(r1)
	for i := e.BitLen(); i >= 0; i-- {
		mul(z, z, z)
		if e.Bit(i) == 1 {
			mul(z, z, a)
		}
	}
	c.Set(z)
}

func inverse(inv, e *Fe) {
	if e.isZero() {
		inv.zero()
		return
	}
	u := new(Fe).Set(&modulus)
	v := new(Fe).Set(e)
	s := &Fe{1}
	r := &Fe{0}
	var k int
	var z uint64
	var found = false
	// Phase 1
	for i := 0; i < sixWordBitSize*2; i++ {
		if v.isZero() {
			found = true
			break
		}
		if u.isEven() {
			u.div2(0)
			s.mul2()
		} else if v.isEven() {
			v.div2(0)
			z += r.mul2()
		} else if u.cmp(v) == 1 {
			lsubAssign(u, v)
			u.div2(0)
			laddAssign(r, s)
			s.mul2()
		} else {
			lsubAssign(v, u)
			v.div2(0)
			laddAssign(s, r)
			z += r.mul2()
		}
		k += 1
	}

	if !found {
		inv.zero()
		return
	}

	if k < fpBitSize || k > fpBitSize+sixWordBitSize {
		inv.zero()
		return
	}

	if r.cmp(&modulus) != -1 || z > 0 {
		lsubAssign(r, &modulus)
	}
	u.Set(&modulus)
	lsubAssign(u, r)

	// Phase 2
	for i := k; i < 2*sixWordBitSize; i++ {
		double(u, u)
	}
	inv.Set(u)
}

func inverseBatch(in []Fe) {

	n, N, setFirst := 0, len(in), false

	for i := 0; i < len(in); i++ {
		if !in[i].isZero() {
			n++
		}
	}
	if n == 0 {
		return
	}

	tA := make([]Fe, n)
	tB := make([]Fe, n)

	for i, j := 0, 0; i < N; i++ {
		if !in[i].isZero() {
			if !setFirst {
				setFirst = true
				tA[j].Set(&in[i])
			} else {
				mul(&tA[j], &in[i], &tA[j-1])
			}
			j = j + 1
		}
	}

	inverse(&tB[n-1], &tA[n-1])
	for i, j := N-1, n-1; j != 0; i-- {
		if !in[i].isZero() {
			mul(&tB[j-1], &tB[j], &in[i])
			j = j - 1
		}
	}

	for i, j := 0, 0; i < N; i++ {
		if !in[i].isZero() {
			if setFirst {
				setFirst = false
				in[i].Set(&tB[j])
			} else {
				mul(&in[i], &tA[j-1], &tB[j])
			}
			j = j + 1
		}
	}
}

func rsqrt(c, a *Fe) bool {
	t0, t1 := new(Fe), new(Fe)
	sqrtAddchain(t0, a)
	mul(t1, t0, a)
	square(t1, t1)
	ret := t1.equal(a)
	c.Set(t0)
	return ret
}

func Sqrt(c, a *Fe) bool {
	u, v := new(Fe).Set(a), new(Fe)
	// a ^ (p - 3) / 4
	sqrtAddchain(c, a)
	// a ^ (p + 1) / 4
	mul(c, c, u)

	square(v, c)
	return u.equal(v)
}

func _sqrt(c, a *Fe) bool {
	u, v := new(Fe).Set(a), new(Fe)
	exp(c, a, pPlus1Over4)
	square(v, c)
	return u.equal(v)
}

func sqrtAddchain(c, a *Fe) {
	chain := func(c *Fe, n int, a *Fe) {
		for i := 0; i < n; i++ {
			square(c, c)
		}
		mul(c, c, a)
	}

	t := make([]Fe, 16)
	t[13].Set(a)
	square(&t[0], &t[13])
	mul(&t[8], &t[0], &t[13])
	square(&t[4], &t[0])
	mul(&t[1], &t[8], &t[0])
	mul(&t[6], &t[4], &t[8])
	mul(&t[9], &t[1], &t[4])
	mul(&t[12], &t[6], &t[4])
	mul(&t[3], &t[9], &t[4])
	mul(&t[7], &t[12], &t[4])
	mul(&t[15], &t[3], &t[4])
	mul(&t[10], &t[7], &t[4])
	mul(&t[2], &t[15], &t[4])
	mul(&t[11], &t[10], &t[4])
	square(&t[0], &t[3])
	mul(&t[14], &t[11], &t[4])
	mul(&t[5], &t[0], &t[8])
	mul(&t[4], &t[0], &t[1])

	chain(&t[0], 12, &t[15])
	chain(&t[0], 7, &t[7])
	chain(&t[0], 4, &t[1])
	chain(&t[0], 6, &t[6])
	chain(&t[0], 7, &t[11])
	chain(&t[0], 5, &t[4])
	chain(&t[0], 2, &t[8])
	chain(&t[0], 6, &t[3])
	chain(&t[0], 6, &t[3])
	chain(&t[0], 6, &t[9])
	chain(&t[0], 3, &t[8])
	chain(&t[0], 7, &t[3])
	chain(&t[0], 4, &t[3])
	chain(&t[0], 6, &t[7])
	chain(&t[0], 6, &t[14])
	chain(&t[0], 3, &t[13])
	chain(&t[0], 8, &t[3])
	chain(&t[0], 7, &t[11])
	chain(&t[0], 5, &t[12])
	chain(&t[0], 6, &t[3])
	chain(&t[0], 6, &t[5])
	chain(&t[0], 4, &t[9])
	chain(&t[0], 8, &t[5])
	chain(&t[0], 4, &t[3])
	chain(&t[0], 7, &t[11])
	chain(&t[0], 9, &t[10])
	chain(&t[0], 2, &t[8])
	chain(&t[0], 5, &t[6])
	chain(&t[0], 7, &t[1])
	chain(&t[0], 7, &t[9])
	chain(&t[0], 6, &t[11])
	chain(&t[0], 5, &t[5])
	chain(&t[0], 5, &t[10])
	chain(&t[0], 5, &t[10])
	chain(&t[0], 8, &t[3])
	chain(&t[0], 7, &t[2])
	chain(&t[0], 9, &t[7])
	chain(&t[0], 5, &t[3])
	chain(&t[0], 3, &t[8])
	chain(&t[0], 8, &t[7])
	chain(&t[0], 3, &t[8])
	chain(&t[0], 7, &t[9])
	chain(&t[0], 9, &t[7])
	chain(&t[0], 6, &t[2])
	chain(&t[0], 6, &t[4])
	chain(&t[0], 5, &t[4])
	chain(&t[0], 5, &t[4])
	chain(&t[0], 4, &t[3])
	chain(&t[0], 3, &t[8])
	chain(&t[0], 8, &t[2])
	chain(&t[0], 7, &t[4])
	chain(&t[0], 5, &t[4])
	chain(&t[0], 5, &t[4])
	chain(&t[0], 4, &t[7])
	chain(&t[0], 4, &t[6])
	chain(&t[0], 7, &t[4])
	chain(&t[0], 5, &t[5])
	chain(&t[0], 5, &t[4])
	chain(&t[0], 5, &t[4])
	chain(&t[0], 5, &t[4])
	chain(&t[0], 5, &t[4])
	chain(&t[0], 5, &t[4])
	chain(&t[0], 5, &t[4])
	chain(&t[0], 4, &t[3])
	chain(&t[0], 6, &t[2])
	chain(&t[0], 4, &t[1])
	square(c, &t[0])
}

func IsQuadraticNonResidue(a *Fe) bool {
	if a.isZero() {
		return true
	}
	return !Sqrt(new(Fe), a)
}
