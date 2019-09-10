package gofra

import (
	"math"

	"github.com/gitchander/gofra/fcolor"
)

type Palette struct {
	Colors     []fcolor.RGB `json:"colors"`
	SpaceColor fcolor.RGB   `json:"space_color"`
	Period     float64      `json:"period"`
	Shift      float64      `json:"shift"`
}

func newColorTable(iterations int, pal Palette) []fcolor.RGB {

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
