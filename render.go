package gofra

import (
	"image"
	"runtime"
	"sync"

	"github.com/gitchander/gofra/interval"
	"github.com/gitchander/gofra/mth2d"
)

func RenderImageRGBA(im *image.RGBA, params Parameters) {

	var (
		nX = im.Rect.Dx()
		nY = im.Rect.Dy()

		fi         = params.FractalInfo
		pixelWidth = 2 * fi.Location.Radius / float64(min(nX, nY))
		center     = fi.Location.Center
	)

	var transform mth2d.Matrix
	transform.InitIdendity()
	transform.Translate(-float64(nX)/2, -float64(nY)/2)
	transform.Scale(pixelWidth, pixelWidth)
	mth2d.ReflectAxisX(&transform)
	transform.Rotate(degToRad(fi.Location.Angle))
	transform.Translate(center.Re, center.Im)

	numCPU := runtime.NumCPU()
	countY, countX := nearDivs(numCPU)

	var (
		xIns = interval.Interval{Max: nX}.Split(countX)
		yIns = interval.Interval{Max: nY}.Split(countY)
	)

	si := newSyncImage(im)
	c := newColorСomputer(params, transform)

	wg := new(sync.WaitGroup)

	for _, yIn := range yIns {
		for _, xIn := range xIns {

			var (
				r = image.Rectangle{
					Min: image.Point{xIn.Min, yIn.Min},
					Max: image.Point{xIn.Max, yIn.Max},
				}
				cc = c.Clone()
			)

			wg.Add(1)
			go renderRectangle(wg, r, cc, si)
		}
	}

	wg.Wait()
}

func renderRectangle(wg *sync.WaitGroup, r image.Rectangle, cc colorСomputer, si *syncImage) {

	defer wg.Done()

	var (
		y0 = r.Min.Y
		yn = r.Max.Y

		x0 = r.Min.X
		xn = r.Max.X

		clСompute = cc.colorСompute

		ti = image.NewRGBA(r)
	)

	for y := y0; y < yn; y++ {
		for x := x0; x < xn; x++ {
			ti.Set(x, y, clСompute(x, y))
		}
	}

	si.Draw(r, ti)
}
