package main

import (
	"image"

	"github.com/gitchander/gofra"

	"github.com/vbauerster/mpb"
	"github.com/vbauerster/mpb/decor"
)

func renderWithProgress(m *image.RGBA, params *gofra.Parameters) {

	p := mpb.New()
	defer p.Stop()

	total := 100

	bar := p.AddBar(int64(total),
		mpb.PrependDecorators(
			decor.StaticName("Render fractal:", 0, 0),
			decor.Percentage(3, decor.DSyncSpace),
		),
		mpb.AppendDecorators(
			decor.ETA(2, 0),
		),
	)

	count := 0
	progress := func(percent int) {
		bar.Incr(percent - count)
		count = percent
	}

	gofra.RenderImageRGBA(m, *params, progress)
}
