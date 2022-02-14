package bigm

import (
	"math"
	"math/big"
)

type Matrix3x2 struct {
	fs []*big.Float
}

func NewMatrix3x2() *Matrix3x2 {
	return &Matrix3x2{
		fs: makeFloats(6),
	}
}

func makeFloats(n int) []*big.Float {
	return makeFloatsPrec(n, 128)
}

func makeFloatsPrec(n int, prec uint) []*big.Float {
	xs := make([]*big.Float, n)
	for i := range xs {
		x := new(big.Float)
		x.SetPrec(prec)
		xs[i] = x
	}
	return xs
}

func setFloats(xs []*big.Float, a float64) {
	for _, x := range xs {
		x.SetFloat64(a)
	}
}

func (p *Matrix3x2) InitIdendity() {

	// -----------
	// | 1  0  - |
	// | 0  1  - |
	// | 0  0  - |
	// -----------

	as := []float64{
		1, 0,
		0, 1,
		0, 0,
	}
	for i, a := range as {
		p.fs[i].SetFloat64(a)
	}
}

func (p *Matrix3x2) InitTranslate(x, y *big.Float) {

	// -----------
	// | 1  0  - |
	// | 0  1  - |
	// | x  y  - |
	// -----------

	as := []float64{
		1, 0,
		0, 1,
	}
	for i, a := range as {
		p.fs[i].SetFloat64(a)
	}

	p.fs[4].Set(x)
	p.fs[5].Set(y)
}

func (p *Matrix3x2) InitScale(x, y *big.Float) {

	// -----------
	// | x  0  - |
	// | 0  y  - |
	// | 0  0  - |
	// -----------

	setFloats(p.fs[:6], 0)
	p.fs[0].Set(x)
	p.fs[3].Set(y)
}

func (p *Matrix3x2) InitRotate(angle float64) {

	// ----------------
	// | cos  -sin  - |
	// | sin  cos   - |
	// |  0    0    - |
	// ----------------

	sin, cos := math.Sincos(angle)
	as := []float64{
		cos, -sin,
		sin, cos,
		0, 0,
	}
	for i, a := range as {
		p.fs[i].SetFloat64(a)
	}
}

func (p *Matrix3x2) InitReflectAxisX() {

	// ------------
	// | 1   0  - |
	// | 0  -1  - |
	// | 0   0  - |
	// ------------

	as := []float64{
		1, 0,
		0, -1,
		0, 0,
	}
	for i, a := range as {
		p.fs[i].SetFloat64(a)
	}
}

func (p *Matrix3x2) InitReflectAxisY() {

	// ------------
	// | -1  0  - |
	// |  0  1  - |
	// |  0  0  - |
	// ------------

	as := []float64{
		-1, 0,
		0, 1,
		0, 0,
	}
	for i, a := range as {
		p.fs[i].SetFloat64(a)
	}
}

func (p *Matrix3x2) Translate(x, y *big.Float) {

	// m[4] += x
	// m[5] += y

	p.fs[4].Add(p.fs[4], x)
	p.fs[5].Add(p.fs[5], y)
}

func (p *Matrix3x2) Scale(x, y *big.Float) {

	//-------------------------------------
	// m[0] *= x
	// m[1] *= y

	// m[2] *= x
	// m[3] *= y

	// m[4] *= x
	// m[5] *= y
	//-------------------------------------

	p.fs[0].Mul(p.fs[0], x)
	p.fs[1].Mul(p.fs[1], y)

	p.fs[2].Mul(p.fs[2], x)
	p.fs[3].Mul(p.fs[3], y)

	p.fs[4].Mul(p.fs[4], x)
	p.fs[5].Mul(p.fs[5], y)
}
