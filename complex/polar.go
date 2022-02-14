package complex

import (
	"math"
)

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
