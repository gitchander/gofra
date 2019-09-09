package gofra

import (
	"bufio"
	"image"
	"image/draw"
	"image/png"
	"os"
	"sync"
)

type syncImage struct {
	guard sync.Mutex
	m     draw.Image
}

func newSyncImage(m draw.Image) *syncImage {
	return &syncImage{m: m}
}

func (p *syncImage) Draw(r image.Rectangle, src image.Image) {
	p.guard.Lock()
	draw.Draw(p.m, r, src, r.Min, draw.Src)
	p.guard.Unlock()
}

func NewImageRGBA(size image.Point) *image.RGBA {

	var (
		Dx = size.X
		Dy = size.Y
	)

	if Dx <= 0 {
		Dx = 1
	}
	if Dy <= 0 {
		Dy = 1
	}

	return image.NewRGBA(image.Rect(0, 0, Dx, Dy))
}

func ImageSaveToPNG(m image.Image, filename string) error {

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	defer w.Flush()

	return png.Encode(w, m)
}
