package main

import (
	"image"

	"github.com/gitchander/gofra"

	"github.com/vbauerster/mpb"
	"github.com/vbauerster/mpb/decor"
)

func renderWithProgress(m *image.RGBA, c *gofra.Config) {

	p := mpb.New()

	total := 100

	bar := p.AddBar(int64(total),
		mpb.PrependDecorators(
			decor.Name("Rendering fractal: ", decor.WC{W: 0, C: 0}),
			decor.Percentage(decor.WC{W: 0, C: 0}),
		),
		mpb.AppendDecorators(
			decor.Elapsed(decor.ET_STYLE_GO),
			// decor.Elapsed(decor.ET_STYLE_MMSS),
			// decor.Elapsed(decor.ET_STYLE_HHMMSS),
		),
	)

	count := 0
	progress := func(percent int) {
		bar.IncrInt64(int64(percent - count))
		count = percent
	}

	gofra.RenderImageRGBA(m, *c, progress)

	p.Wait()
}
