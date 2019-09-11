package palgen

import (
	"image"

	"github.com/gitchander/gofra/fcolor"
)

func Draw1(m *image.RGBA) {
	r := m.Bounds()
	cs := []fcolor.RGB{
		{R: 0, G: 0, B: 0},
		{R: 1, G: 1, B: 1},
	}
	for x := r.Min.X; x < r.Max.X; x++ {
		t := float64(x-r.Min.X) / float64(r.Dx())
		c := fcolor.LerpRGB(cs[0], cs[1], t)
		for y := r.Min.Y; y < r.Max.Y; y++ {
			m.Set(x, y, c)
		}
	}
}

func Draw2(m *image.RGBA) {

	var (
		bounds = m.Bounds()

		x0 = bounds.Min.X
		x1 = bounds.Max.X

		y0 = bounds.Min.Y
		y1 = bounds.Max.Y

		dx = bounds.Dx()
	)

	var p Params

	RandParams(NewRandNow(), &p)
	//p = ps[1]

	for x := x0; x < x1; x++ {
		t := float64(x-x0) / float64(dx) // [0..1)
		c := ColorByParams(p, t)
		for y := y0; y < y1; y++ {
			m.Set(x, y, c)
		}
	}
}
