package bls12381

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"math/big"
)

// Fe is base field element representation
type Fe /***			***/ [fpNumberOfLimbs]uint64

// Fe2 is element representation of 'fp2' which is quadratic extention of base field 'fp'
// Representation follows c[0] + c[1] * u encoding order.
type Fe2 /**			***/ [2]Fe

// fe6 is element representation of 'fp6' field which is cubic extention of 'fp2'
// Representation follows c[0] + c[1] * v + c[2] * v^2 encoding order.
type fe6 /**			***/ [3]Fe2

// Fe12 is element representation of 'fp12' field which is quadratic extention of 'fp6'
// Representation follows c[0] + c[1] * w encoding order.
type Fe12 /**			***/ [2]fe6

type wfe /***			***/ [fpNumberOfLimbs * 2]uint64
type wfe2 /**			***/ [2]wfe
type wfe6 /**			***/ [3]wfe2

func NewFe() *Fe {
	return &Fe{}
}

func NewFe12() *Fe12 {
	return new(Fe12).zero()
}

func (fe *Fe) SetBytesToMont(in []byte) *Fe {
	l := len(in)
	if l >= fpByteSize {
		l = fpByteSize
	}
	padded := make([]byte, fpByteSize)
	copy(padded[fpByteSize-l:], in[:])
	var a int
	for i := 0; i < fpNumberOfLimbs; i++ {
		a = i * 8
		fe[i] = uint64(padded[a]) | uint64(padded[a+1])<<8 |
			uint64(padded[a+2])<<16 | uint64(padded[a+3])<<24 |
			uint64(padded[a+4])<<32 | uint64(padded[a+5])<<40 |
			uint64(padded[a+6])<<48 | uint64(padded[a+7])<<56
	}
	toMont(fe, fe)
	return fe
}

func (fe *Fe) ToMont() *Fe {
	toMont(fe, fe)
	return fe
}

func (fe *Fe) FromMont() *Fe {
	fromMont(fe, fe)
	return fe
}

func (fe *Fe) setBytes(in []byte) *Fe {
	l := len(in)
	if l >= fpByteSize {
		l = fpByteSize
	}
	padded := make([]byte, fpByteSize)
	copy(padded[fpByteSize-l:], in[:])
	var a int
	for i := 0; i < fpNumberOfLimbs; i++ {
		a = fpByteSize - i*8
		fe[i] = uint64(padded[a-1]) | uint64(padded[a-2])<<8 |
			uint64(padded[a-3])<<16 | uint64(padded[a-4])<<24 |
			uint64(padded[a-5])<<32 | uint64(padded[a-6])<<40 |
			uint64(padded[a-7])<<48 | uint64(padded[a-8])<<56
	}
	return fe
}

func (fe *Fe) SetBig(a *big.Int) *Fe {
	fe.setBytes(a.Bytes())
	return fe
}

func (fe *Fe) setString(s string) (*Fe, error) {
	if s[:2] == "0x" {
		s = s[2:]
	}
	bytes, err := hex.DecodeString(s)
	if err != nil {
		return nil, err
	}
	return fe.setBytes(bytes), nil
}

func (fe *Fe) Set(fe2 *Fe) *Fe {
	fe[0] = fe2[0]
	fe[1] = fe2[1]
	fe[2] = fe2[2]
	fe[3] = fe2[3]
	fe[4] = fe2[4]
	fe[5] = fe2[5]
	return fe
}

func (fe *Fe) BytesFromMont() []byte {
	tmp := new(Fe)
	fromMont(tmp, fe)
	out := make([]byte, fpByteSize)
	var a int
	for i := 0; i < fpNumberOfLimbs; i++ {
		a = i * 8
		out[a] = byte(tmp[i])
		out[a+1] = byte(tmp[i] >> 8)
		out[a+2] = byte(tmp[i] >> 16)
		out[a+3] = byte(tmp[i] >> 24)
		out[a+4] = byte(tmp[i] >> 32)
		out[a+5] = byte(tmp[i] >> 40)
		out[a+6] = byte(tmp[i] >> 48)
		out[a+7] = byte(tmp[i] >> 56)
	}
	return out
}

func (fe *Fe) Bytes() []byte {
	out := make([]byte, fpByteSize)
	var a int
	for i := 0; i < fpNumberOfLimbs; i++ {
		a = fpByteSize - i*8
		out[a-1] = byte(fe[i])
		out[a-2] = byte(fe[i] >> 8)
		out[a-3] = byte(fe[i] >> 16)
		out[a-4] = byte(fe[i] >> 24)
		out[a-5] = byte(fe[i] >> 32)
		out[a-6] = byte(fe[i] >> 40)
		out[a-7] = byte(fe[i] >> 48)
		out[a-8] = byte(fe[i] >> 56)
	}
	return out
}

func (fe *Fe) ToBig() *big.Int {
	return new(big.Int).SetBytes(fe.BytesFromMont())
}

func (fe *Fe) big() *big.Int {
	return new(big.Int).SetBytes(fe.Bytes())
}

func (fe *Fe) string() (s string) {
	for i := fpNumberOfLimbs - 1; i >= 0; i-- {
		s = fmt.Sprintf("%s%16.16x", s, fe[i])
	}
	return "0x" + s
}

func (fe *Fe) zero() *Fe {
	fe[0] = 0
	fe[1] = 0
	fe[2] = 0
	fe[3] = 0
	fe[4] = 0
	fe[5] = 0
	return fe
}

func (fe *Fe) One() *Fe {
	return fe.Set(r1)
}

func (fe *Fe) Rand(r io.Reader) (*Fe, error) {
	bi, err := rand.Int(r, modulus.big())
	if err != nil {
		return nil, err
	}
	return fe.SetBig(bi), nil
}

func (fe *Fe) isValid() bool {
	return fe.cmp(&modulus) == -1
}

func (fe *Fe) isOdd() bool {
	var mask uint64 = 1
	return fe[0]&mask != 0
}

func (fe *Fe) isEven() bool {
	var mask uint64 = 1
	return fe[0]&mask == 0
}

func (fe *Fe) isZero() bool {
	return (fe[5] | fe[4] | fe[3] | fe[2] | fe[1] | fe[0]) == 0
}

func (fe *Fe) isOne() bool {
	return fe.equal(r1)
}

func (fe *Fe) cmp(fe2 *Fe) int {
	for i := fpNumberOfLimbs - 1; i >= 0; i-- {
		if fe[i] > fe2[i] {
			return 1
		} else if fe[i] < fe2[i] {
			return -1
		}
	}
	return 0
}

func (fe *Fe) equal(fe2 *Fe) bool {
	return fe2[0] == fe[0] && fe2[1] == fe[1] && fe2[2] == fe[2] && fe2[3] == fe[3] && fe2[4] == fe[4] && fe2[5] == fe[5]
}

func (e *Fe) signBE() bool {
	negZ, z := new(Fe), new(Fe)
	fromMont(z, e)
	neg(negZ, z)
	return negZ.cmp(z) > -1
}

func (e *Fe) sign() bool {
	r := new(Fe)
	fromMont(r, e)
	return r[0]&1 == 0
}

func (e *Fe) div2(u uint64) {
	e[0] = e[0]>>1 | e[1]<<63
	e[1] = e[1]>>1 | e[2]<<63
	e[2] = e[2]>>1 | e[3]<<63
	e[3] = e[3]>>1 | e[4]<<63
	e[4] = e[4]>>1 | e[5]<<63
	e[5] = e[5]>>1 | u<<63
}

func (e *Fe) mul2() uint64 {
	u := e[5] >> 63
	e[5] = e[5]<<1 | e[4]>>63
	e[4] = e[4]<<1 | e[3]>>63
	e[3] = e[3]<<1 | e[2]>>63
	e[2] = e[2]<<1 | e[1]>>63
	e[1] = e[1]<<1 | e[0]>>63
	e[0] = e[0] << 1
	return u
}

func NewFe2() *Fe2 {
	return &Fe2{}
}

func (e *Fe2) Zero() *Fe2 {
	e[0].zero()
	e[1].zero()
	return e
}

func (e *Fe2) One() *Fe2 {
	e[0].One()
	e[1].zero()
	return e
}

func (e *Fe2) Set(e2 *Fe2) *Fe2 {
	e[0].Set(&e2[0])
	e[1].Set(&e2[1])
	return e
}

func (e *Fe2) fromMont(a *Fe2) {
	fromMont(&e[0], &a[0])
	fromMont(&e[1], &a[1])
}

func (e *Fe2) toMont(a *Fe2) {
	toMont(&e[0], &a[0])
	toMont(&e[1], &a[1])
}

func (e *Fe2) fromWide(w *wfe2) {
	fromWide(&e[0], &w[0])
	fromWide(&e[1], &w[1])
}

func (e *Fe2) Rand(r io.Reader) (*Fe2, error) {
	a0, err := new(Fe).Rand(r)
	if err != nil {
		return nil, err
	}
	e[0].Set(a0)
	a1, err := new(Fe).Rand(r)
	if err != nil {
		return nil, err
	}
	e[1].Set(a1)
	return e, nil
}

func (e *Fe2) isOne() bool {
	return e[0].isOne() && e[1].isZero()
}

func (e *Fe2) isZero() bool {
	return e[0].isZero() && e[1].isZero()
}

func (e *Fe2) equal(e2 *Fe2) bool {
	return e[0].equal(&e2[0]) && e[1].equal(&e2[1])
}

func (e *Fe2) signBE() bool {
	if !e[1].isZero() {
		return e[1].signBE()
	}
	return e[0].signBE()
}

func (e *Fe2) sign() bool {
	r := new(Fe)
	if !e[0].isZero() {
		fromMont(r, &e[0])
		return r[0]&1 == 0
	}
	fromMont(r, &e[1])
	return r[0]&1 == 0
}

func (e *fe6) zero() *fe6 {
	e[0].Zero()
	e[1].Zero()
	e[2].Zero()
	return e
}

func (e *fe6) one() *fe6 {
	e[0].One()
	e[1].Zero()
	e[2].Zero()
	return e
}

func (e *fe6) set(e2 *fe6) *fe6 {
	e[0].Set(&e2[0])
	e[1].Set(&e2[1])
	e[2].Set(&e2[2])
	return e
}

func (e *fe6) fromMont(a *fe6) {
	e[0].fromMont(&a[0])
	e[1].fromMont(&a[1])
	e[2].fromMont(&a[2])
}

func (e *fe6) toMont(a *fe6) {
	e[0].toMont(&a[0])
	e[1].toMont(&a[1])
	e[2].toMont(&a[2])
}

func (e *fe6) fromWide(w *wfe6) {
	e[0].fromWide(&w[0])
	e[1].fromWide(&w[1])
	e[2].fromWide(&w[2])
}

func (e *fe6) rand(r io.Reader) (*fe6, error) {
	a0, err := new(Fe2).Rand(r)
	if err != nil {
		return nil, err
	}
	e[0].Set(a0)
	a1, err := new(Fe2).Rand(r)
	if err != nil {
		return nil, err
	}
	e[1].Set(a1)
	a2, err := new(Fe2).Rand(r)
	if err != nil {
		return nil, err
	}
	e[2].Set(a2)
	return e, nil
}

func (e *fe6) isOne() bool {
	return e[0].isOne() && e[1].isZero() && e[2].isZero()
}

func (e *fe6) isZero() bool {
	return e[0].isZero() && e[1].isZero() && e[2].isZero()
}

func (e *fe6) equal(e2 *fe6) bool {
	return e[0].equal(&e2[0]) && e[1].equal(&e2[1]) && e[2].equal(&e2[2])
}

func (e *Fe12) zero() *Fe12 {
	e[0].zero()
	e[1].zero()
	return e
}

func (e *Fe12) one() *Fe12 {
	e[0].one()
	e[1].zero()
	return e
}

func (e *Fe12) set(e2 *Fe12) *Fe12 {
	e[0].set(&e2[0])
	e[1].set(&e2[1])
	return e
}

func (e *Fe12) FromMont(a *Fe12) {
	e[0].fromMont(&a[0])
	e[1].fromMont(&a[1])
}

func (e *Fe12) ToMont(a *Fe12) {
	e[0].toMont(&a[0])
	e[1].toMont(&a[1])
}

func (e *Fe12) Rand(r io.Reader) (*Fe12, error) {
	a0, err := new(fe6).rand(r)
	if err != nil {
		return nil, err
	}
	e[0].set(a0)
	a1, err := new(fe6).rand(r)
	if err != nil {
		return nil, err
	}
	e[1].set(a1)
	return e, nil
}

func (e *Fe12) isOne() bool {
	return e[0].isOne() && e[1].isZero()
}

func (e *Fe12) isZero() bool {
	return e[0].isZero() && e[1].isZero()
}

func (e *Fe12) equal(e2 *Fe12) bool {
	return e[0].equal(&e2[0]) && e[1].equal(&e2[1])
}

func (fe *wfe) set(fe2 *wfe) *wfe {
	fe[0] = fe2[0]
	fe[1] = fe2[1]
	fe[2] = fe2[2]
	fe[3] = fe2[3]
	fe[4] = fe2[4]
	fe[5] = fe2[5]
	fe[6] = fe2[6]
	fe[7] = fe2[7]
	fe[8] = fe2[8]
	fe[9] = fe2[9]
	fe[10] = fe2[10]
	fe[11] = fe2[11]
	return fe
}

func (fe *wfe2) set(fe2 *wfe2) *wfe2 {
	fe[0].set(&fe2[0])
	fe[1].set(&fe2[1])
	return fe
}

func (fe *wfe6) set(fe2 *wfe6) *wfe6 {
	fe[0].set(&fe2[0])
	fe[1].set(&fe2[1])
	fe[2].set(&fe2[2])
	return fe
}
