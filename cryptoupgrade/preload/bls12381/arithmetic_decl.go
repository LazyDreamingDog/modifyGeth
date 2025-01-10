//go:build amd64 && !generic
// +build amd64,!generic

package bls12381

import (
	"golang.org/x/sys/cpu"
)

func init() {
	if !cpu.X86.HasADX || !cpu.X86.HasBMI2 {
		mul = mulNoADX
		wmul = wmulNoADX
		fromWide = montRedNoADX
		mulFR = mulNoADXFR
		wmulFR = wmulNoADXFR
		wfp2Mul = wfp2MulGeneric
		wfp2Square = wfp2SquareGeneric
	}
}

var mul func(c, a, b *Fe) = mulADX
var wmul func(c *wfe, a, b *Fe) = wmulADX
var fromWide func(c *Fe, w *wfe) = montRedADX
var wfp2Mul func(c *wfe2, a, b *Fe2) = wfp2MulADX
var wfp2Square func(c *wfe2, b *Fe2) = wfp2SquareADX

func square(c, a *Fe) {
	mul(c, a, a)
}

func neg(c, a *Fe) {
	if a.isZero() {
		c.Set(a)
	} else {
		_neg(c, a)
	}
}

//go:noescape
func add(c, a, b *Fe)

//go:noescape
func addAssign(a, b *Fe)

//go:noescape
func ladd(c, a, b *Fe)

//go:noescape
func laddAssign(a, b *Fe)

//go:noescape
func double(c, a *Fe)

//go:noescape
func doubleAssign(a *Fe)

//go:noescape
func ldouble(c, a *Fe)

//go:noescape
func ldoubleAssign(a *Fe)

//go:noescape
func sub(c, a, b *Fe)

//go:noescape
func subAssign(a, b *Fe)

//go:noescape
func lsubAssign(a, b *Fe)

//go:noescape
func _neg(c, a *Fe)

//go:noescape
func mulNoADX(c, a, b *Fe)

//go:noescape
func mulADX(c, a, b *Fe)

//go:noescape
func wmulNoADX(c *wfe, a, b *Fe)

//go:noescape
func wmulADX(c *wfe, a, b *Fe)

//go:noescape
func montRedNoADX(a *Fe, w *wfe)

//go:noescape
func montRedADX(a *Fe, w *wfe)

//go:noescape
func lwadd(c, a, b *wfe)

//go:noescape
func lwaddAssign(a, b *wfe)

//go:noescape
func wadd(c, a, b *wfe)

//go:noescape
func lwdouble(c, a *wfe)

//go:noescape
func wdouble(c, a *wfe)

//go:noescape
func lwsub(c, a, b *wfe)

//go:noescape
func lwsubAssign(a, b *wfe)

//go:noescape
func wsub(c, a, b *wfe)

//go:noescape
func fp2Add(c, a, b *Fe2)

//go:noescape
func fp2AddAssign(a, b *Fe2)

//go:noescape
func fp2Ladd(c, a, b *Fe2)

//go:noescape
func fp2LaddAssign(a, b *Fe2)

//go:noescape
func fp2DoubleAssign(a *Fe2)

//go:noescape
func fp2Double(c, a *Fe2)

//go:noescape
func fp2Sub(c, a, b *Fe2)

//go:noescape
func fp2SubAssign(a, b *Fe2)

//go:noescape
func mulByNonResidue(c, a *Fe2)

//go:noescape
func mulByNonResidueAssign(a *Fe2)

//go:noescape
func wfp2Add(c, a, b *wfe2)

//go:noescape
func wfp2AddAssign(a, b *wfe2)

//go:noescape
func wfp2Ladd(c, a, b *wfe2)

//go:noescape
func wfp2LaddAssign(a, b *wfe2)

//go:noescape
func wfp2AddMixed(c, a, b *wfe2)

//go:noescape
func wfp2AddMixedAssign(a, b *wfe2)

//go:noescape
func wfp2Sub(c, a, b *wfe2)

//go:noescape
func wfp2SubAssign(a, b *wfe2)

//go:noescape
func wfp2SubMixed(c, a, b *wfe2)

//go:noescape
func wfp2SubMixedAssign(a, b *wfe2)

//go:noescape
func wfp2Double(c, a *wfe2)

//go:noescape
func wfp2DoubleAssign(a *wfe2)

//go:noescape
func wfp2MulByNonResidue(c, a *wfe2)

//go:noescape
func wfp2MulByNonResidueAssign(a *wfe2)

//go:noescape
func wfp2SquareADX(c *wfe2, a *Fe2)

//go:noescape
func wfp2MulADX(c *wfe2, a, b *Fe2)

var mulFR func(c, a, b *Fr) = mulADXFR
var wmulFR func(c *wideFr, a, b *Fr) = wmulADXFR

func squareFR(c, a *Fr) {
	mulFR(c, a, a)
}

func negFR(c, a *Fr) {
	if a.IsZero() {
		c.Set(a)
	} else {
		_negFR(c, a)
	}
}

//go:noescape
func addFR(c, a, b *Fr)

//go:noescape
func laddAssignFR(a, b *Fr)

//go:noescape
func doubleFR(c, a *Fr)

//go:noescape
func subFR(c, a, b *Fr)

//go:noescape
func lsubAssignFR(a, b *Fr)

//go:noescape
func _negFR(c, a *Fr)

//go:noescape
func mulNoADXFR(c, a, b *Fr)

//go:noescape
func mulADXFR(c, a, b *Fr)

//go:noescape
func wmulADXFR(c *wideFr, a, b *Fr)

//go:noescape
func wmulNoADXFR(c *wideFr, a, b *Fr)

//go:noescape
func waddFR(a, b *wideFr)
