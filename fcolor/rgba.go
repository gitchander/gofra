package fcolor

import (
	"image/color"
	"math"
)

type RGBA struct {
	R, G, B, A float64
}

var _ color.Color = RGBA{}

func (c RGBA) RGBA() (r, g, b, a uint32) {
	const k = math.MaxUint16
	r = uint32(math.Floor(c.R * k))
	g = uint32(math.Floor(c.G * k))
	b = uint32(math.Floor(c.B * k))
	a = uint32(math.Floor(c.A * k))
	return
}

// Alpha blending
// c = a over b
func (a RGBA) Over(b RGBA) (c RGBA) {
	c.A = lerp(b.A, 1.0, a.A)
	c.R = lerp(b.R*b.A, a.R, a.A) / c.A
	c.G = lerp(b.G*b.A, a.G, a.A) / c.A
	c.B = lerp(b.B*b.A, a.B, a.A) / c.A
	return
}

var RGBAModel = color.ModelFunc(rgbaModel)

func rgbaModel(c color.Color) color.Color {
	if _, ok := c.(RGBA); ok {
		return c
	}
	r, g, b, a := c.RGBA()
	const k = 1 / float64(math.MaxUint16)
	return RGBA{
		R: float64(r) * k,
		G: float64(g) * k,
		B: float64(b) * k,
		A: float64(a) * k,
	}
}
