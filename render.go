package gofra

import (
	"image"
	"runtime"
	"sync"
	"time"

	"github.com/gitchander/gofra/interval"
	"github.com/gitchander/gofra/math2d"
)

func RenderImageRGBA(m *image.RGBA, config Config, progress func(percent int)) {

	var (
		nX = m.Rect.Dx()
		nY = m.Rect.Dy()

		fi         = config.FractalInfo
		pixelWidth = 2 * fi.Location.Radius / float64(minInt(nX, nY))
		center     = fi.Location.Center
	)

	var transform math2d.Matrix
	transform.InitIdendity()
	transform.Translate(-float64(nX)/2, -float64(nY)/2)
	transform.Scale(pixelWidth, pixelWidth)
	transform.ReflectAxisX()
	transform.Rotate(degToRad(float64(fi.Location.AngleDeg)))
	transform.Translate(center.Re, center.Im)

	numCPU := runtime.NumCPU()
	countY, countX := nearDivs(numCPU)

	var (
		xIns = interval.Interval{Max: nX}.Split(countX)
		yIns = interval.Interval{Max: nY}.Split(countY)
	)

	si := newSyncImage(m)
	c := newColorСomputer(config, transform)

	var wg sync.WaitGroup
	points := make(chan int)

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
			go renderRectangle(&wg, r, cc, si, points)
		}
	}

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	total := nX * nY
	count := 0
progressLoop:
	for {
		select {
		case <-ticker.C:
			percent := count * 100 / total
			progress(percent)

		case n := <-points:
			count += n
			if count >= total {
				progress(100) // 100%
				break progressLoop
			}
		}
	}
	wg.Wait()
}

func renderRectangle(wg *sync.WaitGroup, r image.Rectangle, cc colorСomputer, si *syncImage, points chan<- int) {

	var (
		y0 = r.Min.Y
		yn = r.Max.Y

		x0 = r.Min.X
		xn = r.Max.X

		dx = xn - x0

		clСompute = cc.colorСompute

		ti = image.NewRGBA(r)
	)

	for y := y0; y < yn; y++ {
		for x := x0; x < xn; x++ {
			ti.Set(x, y, clСompute(x, y))
		}
		points <- dx
	}

	si.Draw(r, ti)
	wg.Done()
}
