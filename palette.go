package gofra

import (
	"math"

	fcolor "github.com/gitchander/gofra/color"
)

type Palette struct {
	Colors     []fcolor.RGB
	SpaceColor fcolor.RGB
	Period     float64
	Shift      float64
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

	clerp := fcolor.LerpRGB

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
