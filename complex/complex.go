package complex

import "math"

type Complex struct {
	Re float64
	Im float64
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
	t := b.Norm()
	c.Re = (a.Re*b.Re + a.Im*b.Im) / t
	c.Im = (a.Im*b.Re - a.Re*b.Im) / t
	return
}

func (a Complex) AddScalar(t float64) (c Complex) {
	c.Re = a.Re + t
	c.Im = a.Im
	return
}

func (a Complex) SubScalar(t float64) (c Complex) {
	c.Re = a.Re - t
	c.Im = a.Im
	return
}

func (a Complex) MulScalar(t float64) (c Complex) {
	c.Re = a.Re * t
	c.Im = a.Im * t
	return
}

func (a Complex) DivScalar(t float64) (c Complex) {
	c.Re = a.Re / t
	c.Im = a.Im / t
	return
}

func (a Complex) Conjugate() (b Complex) {
	b.Re = a.Re
	b.Im = -a.Im
	return
}

func (a Complex) Norm() float64 {
	return a.Re*a.Re + a.Im*a.Im
}

func (a Complex) magnitude() float64 {
	return math.Sqrt(a.Norm())
}

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
	t := a.Norm()
	b.Re = a.Re / t
	b.Im = -a.Im / t
	return
}

func (a Complex) ToPolar() Polar {
	return Polar{
		Radius: a.Magnitude(),
		Angle:  a.Argument(),
	}
}

func (a Complex) Power(p float64) Complex {
	return Polar{
		Radius: math.Exp(p * 0.5 * math.Log(a.Norm())),
		Angle:  a.Argument() * p,
	}.ToComplex()
}

// Trigonometric form
type Polar struct {
	Radius float64
	Angle  float64
}

func (p Polar) ToComplex() Complex {
	sin, cos := math.Sincos(p.Angle)
	return Complex{
		Re: p.Radius * cos,
		Im: p.Radius * sin,
	}
}
