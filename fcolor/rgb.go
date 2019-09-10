package fcolor

import (
	"encoding/json"
	"image/color"
	"math"
)

type RGB struct {
	R, G, B float64
}

var _ color.Color = RGB{}

func (c RGB) RGBA() (r, g, b, a uint32) {
	r = uint32(math.Floor(c.R * maxUint16))
	g = uint32(math.Floor(c.G * maxUint16))
	b = uint32(math.Floor(c.B * maxUint16))
	a = maxUint16
	return
}

func (c RGB) MarshalJSON() ([]byte, error) {

	const k = math.MaxUint8

	var (
		r = uint8(cropFloat64(c.R, 0, 1) * k)
		g = uint8(cropFloat64(c.G, 0, 1) * k)
		b = uint8(cropFloat64(c.B, 0, 1) * k)
	)

	s := rgb_to_str(r, g, b)

	return json.Marshal(s)
}

func (c *RGB) UnmarshalJSON(data []byte) error {

	var s string

	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	var r, g, b uint8

	if err := str_to_rgb(s, &r, &g, &b); err != nil {
		return err
	}

	const k = 1 / float64(math.MaxUint8)

	c.R = float64(r) * k
	c.G = float64(g) * k
	c.B = float64(b) * k

	return nil
}

func (c RGB) Norm() RGB {
	return RGB{
		R: cropFloat64(c.R, 0, 1),
		G: cropFloat64(c.G, 0, 1),
		B: cropFloat64(c.B, 0, 1),
	}
}

// t=[0..1]
// t=0, c0
// t=1, c1
// t=0.5, (c0 * 0.5) + (c1 * 0.5)
func LerpRGB(c0, c1 RGB, t float64) RGB {
	return RGB{
		R: lerp(c0.R, c1.R, t),
		G: lerp(c0.G, c1.G, t),
		B: lerp(c0.B, c1.B, t),
	}
}

func SinerpRGB(c0, c1 RGB, t float64) RGB {
	t = (1 - math.Sin(math.Pi*(t+0.5))) * 0.5
	return LerpRGB(c0, c1, t)
}

func MixRGB(cs []RGB) RGB {

	var r, g, b float64

	for _, c := range cs {
		r += c.R
		g += c.G
		b += c.B
	}

	k := 1 / float64(len(cs))

	r *= k
	g *= k
	b *= k

	return RGB{r, g, b}
}

var RGBModel = color.ModelFunc(rgbModel)

func rgbModel(c color.Color) color.Color {
	if _, ok := c.(RGB); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	return RGB{
		R: float64(r) / maxUint16,
		G: float64(g) / maxUint16,
		B: float64(b) / maxUint16,
	}
}
