package common

import (
	"math/big"
)

// Rational represents a rational number as a fraction
type Rational struct {
	Numerator   uint64
	Denominator uint64
}

// adjustPrecision reduces the fraction in a deterministic way
func adjustPrecision(num, den *big.Int) (uint64, uint64) {
	// 持续检查并调整直到低于阈值
	for num.BitLen() > 64 || den.BitLen() > 64 {
		num.Rsh(num, 1)
		den.Rsh(den, 1)
	}

	// 确保分母不为0
	if den.Sign() == 0 {
		den.SetUint64(1)
	}

	return num.Uint64(), den.Uint64()
}

// NewRational creates a new Rational with deterministic precision adjustment
func NewRational(num, den uint64) *Rational {
	if den == 0 {
		den = 1
	}

	// 转换为big.Int进行GCD计算
	numBig := new(big.Int).SetUint64(num)
	denBig := new(big.Int).SetUint64(den)
	gcd := new(big.Int).GCD(nil, nil, numBig, denBig)

	numBig.Div(numBig, gcd)
	denBig.Div(denBig, gcd)

	return &Rational{
		Numerator:   numBig.Uint64(),
		Denominator: denBig.Uint64(),
	}
}

// Mul multiplies two rational numbers
func (r *Rational) Mul(other *Rational) *Rational {
	return NewRational(
		r.Numerator*other.Numerator,
		r.Denominator*other.Denominator,
	)
}

// Add adds two rationals with deterministic precision adjustment
func (r *Rational) Add(other *Rational) *Rational {
	num1 := new(big.Int).SetUint64(r.Numerator)
	den1 := new(big.Int).SetUint64(r.Denominator)
	num2 := new(big.Int).SetUint64(other.Numerator)
	den2 := new(big.Int).SetUint64(other.Denominator)

	newDen := new(big.Int).Mul(den1, den2)
	term1 := new(big.Int).Mul(num1, den2)
	term2 := new(big.Int).Mul(num2, den1)
	newNum := new(big.Int).Add(term1, term2)

	gcd := new(big.Int).GCD(nil, nil, newNum, newDen)
	newNum.Div(newNum, gcd)
	newDen.Div(newDen, gcd)

	num, den := adjustPrecision(newNum, newDen)
	return &Rational{Numerator: num, Denominator: den}
}

// Sub subtracts another rational number from this one
func (r *Rational) Sub(other *Rational) *Rational {
	// Similar to Add but with subtraction
	commonDenom := lcm(r.Denominator, other.Denominator)
	rMult := commonDenom / r.Denominator
	otherMult := commonDenom / other.Denominator

	return NewRational(
		r.Numerator*rMult-other.Numerator*otherMult,
		commonDenom,
	)
}

// Div divides this rational by another
func (r *Rational) Div(other *Rational) *Rational {
	if other.Numerator == 0 {
		panic("division by zero")
	}
	return NewRational(
		r.Numerator*other.Denominator,
		r.Denominator*other.Numerator,
	)
}

// ToUint64 converts rational to uint64
func (r *Rational) ToUint64() uint64 {
	return r.Numerator / r.Denominator
}

// Compare returns:
// -1 if r < other
// 0 if r == other
// 1 if r > other
func (r *Rational) Compare(other *Rational) int {
	// Convert to common denominator to compare
	commonDenom := lcm(r.Denominator, other.Denominator)
	rNum := r.Numerator * (commonDenom / r.Denominator)
	otherNum := other.Numerator * (commonDenom / other.Denominator)

	if rNum < otherNum {
		return -1
	} else if rNum > otherNum {
		return 1
	}
	return 0
}

// helper functions
func findGCD(a, b uint64) uint64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b uint64) uint64 {
	return a * b / findGCD(a, b)
}

// ToBigInt converts Rational to *big.Int
func (r *Rational) ToBigInt() *big.Int {
	if r.Denominator == 1 {
		return new(big.Int).SetUint64(r.Numerator)
	}
	num := new(big.Int).SetUint64(r.Numerator)
	den := new(big.Int).SetUint64(r.Denominator)
	return new(big.Int).Div(num, den)
}

// FromBigInt creates a new Rational from *big.Int
func FromBigInt(b *big.Int) *Rational {
	if b == nil {
		return NewRational(0, 1)
	}
	// Handle negative numbers
	if b.Sign() < 0 {
		panic("Rational does not support negative numbers")
	}

	// Convert to fraction form
	num := new(big.Int).Set(b)
	den := new(big.Int).SetUint64(1)

	// Adjust precision if needed
	n, d := adjustPrecision(num, den)
	return NewRational(n, d)
}

// ToBigFloat converts Rational to *big.Float for higher precision calculations
func (r *Rational) ToBigFloat() *big.Float {
	num := new(big.Float).SetUint64(r.Numerator)
	den := new(big.Float).SetUint64(r.Denominator)
	return new(big.Float).Quo(num, den)
}
