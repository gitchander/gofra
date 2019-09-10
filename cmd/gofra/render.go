package main

import (
	"image"
	"time"

	"github.com/gitchander/gofra"

	"github.com/vbauerster/mpb"
	"github.com/vbauerster/mpb/decor"
)

func renderWithProgress(m *image.RGBA, c *gofra.Config) {

	p := mpb.New()

	total := 100

	bar := p.AddBar(int64(total),
		mpb.PrependDecorators(
			decor.Name("Render fractal:", decor.WC{W: 0, C: 0}),
			decor.Percentage(decor.WC{W: 0, C: 0}),
		),
		mpb.AppendDecorators(
			decor.Elapsed(decor.ET_STYLE_GO),
			// decor.Elapsed(decor.ET_STYLE_MMSS),
			// decor.Elapsed(decor.ET_STYLE_HHMMSS),
		),
	)

	count := 0
	start := time.Now()
	progress := func(percent int) {
		now := time.Now()
		bar.IncrInt64(int64(percent-count), now.Sub(start))
		start = now
		count = percent
	}

	gofra.RenderImageRGBA(m, *c, progress)

	p.Wait()
}
