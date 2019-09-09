package main

import (
	"image"
	"time"

	"github.com/gitchander/gofra"

	"github.com/vbauerster/mpb"
	"github.com/vbauerster/mpb/decor"
)

func renderWithProgress(m *image.RGBA, params *gofra.Parameters) {

	p := mpb.New()

	total := 100

	bar := p.AddBar(int64(total),
		mpb.PrependDecorators(
			decor.Name("Render fractal:", decor.WC{W: 0, C: 0}),
			//decor.StaticName("Render fractal:", decor.WC{W: 0, C: 0}),
			decor.Percentage(decor.WC{W: 0, C: 0}),
		),
		//		mpb.AppendDecorators(
		//			decor.EwmaETA() ETA(2, 0),
		//		),
	)

	count := 0
	progress := func(percent int) {
		bar.IncrInt64(int64(percent-count), 10*time.Millisecond)
		count = percent
	}

	gofra.RenderImageRGBA(m, *params, progress)

	p.Wait()
}
