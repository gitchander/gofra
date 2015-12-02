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
	im    draw.Image
	mutex sync.Mutex
}

func newSyncImage(im draw.Image) *syncImage {
	return &syncImage{im: im}
}

func (si *syncImage) Draw(r image.Rectangle, src image.Image) {

	si.mutex.Lock()
	defer si.mutex.Unlock()

	draw.Draw(si.im, r, src, r.Min, draw.Src)
}

func NewImageRGBA(Width, Height int) *image.RGBA {

	var (
		Dx = Width
		Dy = Height
	)

	if Dx <= 0 {
		Dx = 1
	}
	if Dy <= 0 {
		Dy = 1
	}

	return image.NewRGBA(image.Rect(0, 0, Dx, Dy))
}

func ImageSaveToPNG(im image.Image, filename string) error {

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	defer w.Flush()

	return png.Encode(w, im)
}
