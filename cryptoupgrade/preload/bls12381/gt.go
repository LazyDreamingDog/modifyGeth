package bls12381

import (
	"math/big"
)

// E is type for target group element
type E = Fe12

// GT is type for target multiplicative group GT.
type GT struct {
	fp12 *fp12
}

func NewE() *E {
	return new(E)
}

// Set copies given value into the destination
func (e *E) Set(e2 *E) *E {
	return e.set(e2)
}

// One sets a new target group element to One
func (e *E) One() *E {
	e.one()
	return e
}

func (e *E) Zero() *E {
	e.zero()
	return e
}

// IsOne returns true if given element equals to One
func (e *E) IsOne() bool {
	return e.isOne()
}

// Equal returns true if given two element is equal, otherwise returns false
func (g *E) Equal(g2 *E) bool {
	return g.equal(g2)
}

// NewGT constructs new target group instance.
func NewGT() *GT {
	fp12 := newFp12(nil)
	return &GT{fp12}
}

// Q returns group order in big.Int.
func (g *GT) Q() *big.Int {
	return new(big.Int).Set(qBig)
}

// FromBytes expects 576 byte input and returns target group element
// FromBytes returns error if given element is not on correct subgroup.
func (g *GT) FromBytes(in []byte) (*E, error) {
	e, err := g.fp12.fromBytes(in)
	if err != nil {
		return nil, err
	}
	// if !g.IsValid(e) {
	// 	return e, errors.New("invalid element")
	// }
	return e, nil
}

// ToBytes serializes target group element.
func (g *GT) ToBytes(e *E) []byte {
	return g.fp12.toBytes(e)
}

// IsValid checks whether given target group element is in correct subgroup.
func (g *GT) IsValid(e *E) bool {
	r0, r1, r2 := g.New().Set(e), g.New(), g.New()
	g.fp12.frobeniusMap1(r0) // r0 = e^p
	r1.set(r0)               // r1 = e^p
	g.fp12.frobeniusMap1(r0) // r0 = e^(p^2)
	r2.set(r0)               // r2 = e^(p^2)
	g.fp12.frobeniusMap2(r0) // r0 = e^(p^4)
	g.Mul(r0, r0, e)         // r0 = e·e^(p^4)
	// cyclotomic test
	if !r0.Equal(r2) {
		return false
	}
	g.Exp(r0, e, bigFromHex("0xd201000000010000")) // r0 = e^-u
	g.Mul(r0, r0, r1)                              // r0 = e^-u · e^p = e^(p-u)

	// e^(p-u) = e^(p+1-t) == 1
	return r0.IsOne()
}

// New initializes a new target group element which is equal to One
func (g *GT) New() *E {
	return new(E).One()
}

// Add adds two field element `a` and `b` and assigns the result to the element in first argument.
func (g *GT) Add(c, a, b *E) {
	Fp12Add(c, a, b)
}

// Sub subtracts two field element `a` and `b`, and assigns the result to the element in first argument.
func (g *GT) Sub(c, a, b *E) {
	Fp12Sub(c, a, b)
}

// Mul multiplies two field element `a` and `b` and assigns the result to the element in first argument.
func (g *GT) Mul(c, a, b *E) {
	g.fp12.mul(c, a, b)
}

// Square squares an element `a` and assigns the result to the element in first argument.
func (g *GT) Square(c, a *E) {
	c.set(a)
	g.fp12.cyclotomicSquare(c)
}

// Exp exponents an element `a` by a scalar `s` and assigns the result to the element in first argument.
func (g *GT) Exp(c, a *E, s *big.Int) {
	g.fp12.cyclotomicExp(c, a, s)
}

// Inverse inverses an element `a` and assigns the result to the element in first argument.
func (g *GT) Inverse(c, a *E) {
	g.fp12.inverse(c, a)
}
