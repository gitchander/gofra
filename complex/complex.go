package complex

import "math"

type Complex struct {
	Re float64 `json:"re"`
	Im float64 `json:"im"`
}

func (a Complex) Add(b Complex) (c Complex) {
	c.Re = a.Re + b.Re
	c.Im = a.Im + b.Im
	return
}

func (a Complex) Sub(b Complex) (c Complex) {
	c.Re = a.Re - b.Re
	c.Im = a.Im - b.Im
	return
}

func (a Complex) Mul(b Complex) (c Complex) {
	c.Re = a.Re*b.Re - a.Im*b.Im
	c.Im = a.Im*b.Re + a.Re*b.Im
	return
}

func (a Complex) Div(b Complex) (c Complex) {
	norm := b.Norm()
	c.Re = (a.Re*b.Re + a.Im*b.Im) / norm
	c.Im = (a.Im*b.Re - a.Re*b.Im) / norm
	return
}

func (a Complex) AddScalar(scalar float64) (c Complex) {
	c.Re = a.Re + scalar
	c.Im = a.Im
	return
}

func (a Complex) SubScalar(scalar float64) (c Complex) {
	c.Re = a.Re - scalar
	c.Im = a.Im
	return
}

func (a Complex) MulScalar(scalar float64) (c Complex) {
	c.Re = a.Re * scalar
	c.Im = a.Im * scalar
	return
}

func (a Complex) DivScalar(scalar float64) (c Complex) {
	c.Re = a.Re / scalar
	c.Im = a.Im / scalar
	return
}

func (a Complex) Conjugate() (b Complex) {
	b.Re = a.Re
	b.Im = -a.Im
	return
}

func (a Complex) Norm() float64 {
	return (a.Re * a.Re) + (a.Im * a.Im)
}

// Simple Magnitude
func (a Complex) _magnitude() float64 {
	return math.Sqrt(a.Norm())
}

// Better Magnitude
func (z Complex) Magnitude() float64 {

	var (
		a = math.Abs(z.Re)
		b = math.Abs(z.Im)
	)

	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}

	var m float64
	if a >= b {
		m = b / a
		m = a * math.Sqrt(1+m*m)
	} else {
		m = a / b
		m = b * math.Sqrt(1+m*m)
	}

	return m
}

func (a Complex) Argument() float64 {
	return math.Atan2(a.Im, a.Re)
}

// b = 1 / a
func (a Complex) Invert() (b Complex) {
	norm := a.Norm()
	return Complex{
		Re: a.Re / norm,
		Im: -a.Im / norm,
	}
}

func (a Complex) Polar() Polar {
	return Polar{
		Radius: a.Magnitude(),
		Angle:  a.Argument(),
	}
}

func (a Complex) Power(p float64) Complex {
	return Polar{
		Radius: math.Exp(p * 0.5 * math.Log(a.Norm())),
		Angle:  a.Argument() * p,
	}.Complex()
}

// Trigonometric form
type Polar struct {
	Radius float64
	Angle  float64
}

func (p Polar) Complex() Complex {
	sin, cos := math.Sincos(p.Angle)
	return Complex{
		Re: p.Radius * cos,
		Im: p.Radius * sin,
	}
}
