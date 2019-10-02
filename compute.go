package gofra

import (
	"image/color"

	. "github.com/gitchander/gofra/complex"

	"github.com/gitchander/gofra/fcolor"
	"github.com/gitchander/gofra/math2d"
)

type colorСomputer interface {
	colorСompute(x, y int) color.RGBA
	Clone() colorСomputer
}

type aliasingСomputer struct {
	iterations   int
	colorTable   []fcolor.RGB
	spaceColor   fcolor.RGB
	orbitFactory OrbitFactory
	transform    math2d.Matrix
}

func (c aliasingСomputer) colorСompute(x, y int) color.RGBA {
	var Z Complex
	Z.Re, Z.Im = c.transform.TransformPoint(float64(x), float64(y))

	orbit := c.orbitFactory.NewOrbit(Z)
	iter := TraceOrbit(orbit, c.iterations)

	var fc fcolor.RGB
	if iter == -1 {
		fc = c.spaceColor
	} else {
		fc = c.colorTable[iter]
	}

	return color.RGBAModel.Convert(fc).(color.RGBA)
}

func (c aliasingСomputer) Clone() colorСomputer {

	v := aliasingСomputer{}

	v.iterations = c.iterations

	v.colorTable = make([]fcolor.RGB, len(c.colorTable))
	copy(v.colorTable, c.colorTable)

	v.spaceColor = c.spaceColor

	v.orbitFactory = c.orbitFactory

	v.transform = c.transform

	return v
}

type antiAliasingСomputer struct {
	aliasingСomputer
	spTable  []spPoint    // subpixel shifts
	spColors []fcolor.RGB // subpixel colors
}

func (c antiAliasingСomputer) colorСompute(x, y int) color.RGBA {

	var Z Complex

	for i, dz := range c.spTable {

		Z.Re, Z.Im = c.transform.TransformPoint(
			float64(x)+dz.X,
			float64(y)+dz.Y,
		)

		orbit := c.orbitFactory.NewOrbit(Z)
		iter := TraceOrbit(orbit, c.iterations)

		var fc fcolor.RGB
		if iter == -1 {
			fc = c.spaceColor
		} else {
			fc = c.colorTable[iter]
		}

		c.spColors[i] = fc
	}

	fc := fcolor.MixRGB(c.spColors)

	return color.RGBAModel.Convert(fc).(color.RGBA)
}

func (c antiAliasingСomputer) Clone() colorСomputer {

	v := antiAliasingСomputer{}

	v.aliasingСomputer = c.aliasingСomputer.Clone().(aliasingСomputer)

	v.spTable = make([]spPoint, len(c.spTable))
	copy(v.spTable, c.spTable)

	v.spColors = make([]fcolor.RGB, len(c.spColors))
	copy(v.spColors, c.spColors)

	return v
}

func newColorСomputer(config Config, transform math2d.Matrix) colorСomputer {

	ac := aliasingСomputer{
		iterations:   config.Calculation.Iterations,
		colorTable:   newColorTable(config.Calculation.Iterations, config.Palette),
		spaceColor:   config.Palette.SpaceColor,
		orbitFactory: newOrbitFactory(config.FractalInfo),
		transform:    transform,
	}

	spTable := makeSubpixelTable(config.Calculation.AntiAliasing)
	if len(spTable) == 0 {
		return ac
	}

	return antiAliasingСomputer{
		aliasingСomputer: ac,
		spTable:          spTable,
		spColors:         make([]fcolor.RGB, len(spTable)),
	}
}
