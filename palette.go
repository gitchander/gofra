package gofra

import (
	"math"

	"github.com/gitchander/gofra/fcolor"
	"github.com/gitchander/gofra/palgen"
)

type Palette1 struct {
	Colors     []fcolor.RGB `json:"colors"`
	SpaceColor fcolor.RGB   `json:"space_color"`
	Period     float64      `json:"period"`
	Shift      float64      `json:"shift"`
}

var DefaultPalette1 = Palette1{
	Colors: []fcolor.RGB{
		fcolor.RGB{R: 1, G: 1, B: 1}, // white
		fcolor.RGB{R: 0, G: 0, B: 0}, // black
	},
	SpaceColor: fcolor.RGB{R: 0, G: 0, B: 0}, // black
	Period:     30,
}

func newColorTable1(iterations int, pal Palette1) []fcolor.RGB {

	cs := pal.Colors

	if pal.Period <= 0 {
		pal.Period = 1
	}

	dX := float64(len(cs)) / float64(pal.Period)
	period := dX * pal.Period
	X := fmod(dX*pal.Shift, period)

	n := iterations
	p := make([]fcolor.RGB, n+1)

	//clerp := fcolor.LerpRGB
	clerp := fcolor.SinerpRGB

	for i := 0; i < n; i++ {

		floorX := math.Floor(X)

		j0 := int(floorX)
		j1 := j0 + 1
		if j1 == len(cs) {
			j1 = 0
		}

		p[i] = clerp(cs[j0], cs[j1], X-floorX).Norm()

		X = fmod(X+dX, period)
	}

	p[n] = pal.SpaceColor.Norm()

	return p
}

func fmod(x, n float64) float64 {
	for x >= n {
		x -= n
	}
	for x < 0 {
		x += n
	}
	return x
}

//------------------------------------------------------------------------------
type Palette struct {
	Params     palgen.Params `json:"params"`
	SpaceColor fcolor.RGB    `json:"space_color"`
	Period     float64       `json:"period"`
	Shift      float64       `json:"shift"`
}

var DefaultPalette = Palette{
	Params: palgen.Params{
		A: palgen.Vec3{0.5, 0.5, 0.5},
		B: palgen.Vec3{0.5, 0.5, 0.5},
		C: palgen.Vec3{1.0, 1.0, 1.0},
		D: palgen.Vec3{0.0, 0.0, 0.0},
	},
	SpaceColor: fcolor.RGB{R: 0, G: 0, B: 0}, // black
	Period:     30,
	Shift:      0,
}

func newColorTable(iterations int, pal Palette) []fcolor.RGB {

	if pal.Period <= 0 {
		pal.Period = 1
	}

	n := iterations
	cs := make([]fcolor.RGB, n+1)

	for i := 0; i < n; i++ {

		t := (float64(i) + pal.Shift) / pal.Period

		cs[i] = palgen.ColorByParams(pal.Params, t)
	}

	cs[n] = pal.SpaceColor.Norm()

	return cs
}
