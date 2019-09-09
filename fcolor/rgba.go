package fcolor

import (
	"image/color"
	"math"
)

const maxUint16 = math.MaxUint16

type RGBA struct {
	R, G, B, A float64
}

var _ color.Color = RGBA{}

func (c RGBA) RGBA() (r, g, b, a uint32) {
	r = uint32(math.Floor(c.R * maxUint16))
	g = uint32(math.Floor(c.G * maxUint16))
	b = uint32(math.Floor(c.B * maxUint16))
	a = uint32(math.Floor(c.A * maxUint16))
	return
}

// Alpha blending
// c = a over b
func (a RGBA) Over(b RGBA) RGBA {
	cA := lerp(b.A, 1.0, a.A)
	return RGBA{
		R: lerp(b.R*b.A, a.R, a.A) / cA,
		G: lerp(b.G*b.A, a.G, a.A) / cA,
		B: lerp(b.B*b.A, a.B, a.A) / cA,
		A: cA,
	}
}

var RGBAModel = color.ModelFunc(rgbaModel)

func rgbaModel(c color.Color) color.Color {
	if _, ok := c.(RGBA); ok {
		return c
	}
	r, g, b, a := c.RGBA()
	return RGBA{
		R: float64(r) / maxUint16,
		G: float64(g) / maxUint16,
		B: float64(b) / maxUint16,
		A: float64(a) / maxUint16,
	}
}
