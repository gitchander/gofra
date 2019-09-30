package gofra

import (
	"bytes"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io/ioutil"
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

func SaveImagePNG(filename string, m image.Image) error {
	var buf bytes.Buffer
	err := png.Encode(&buf, m)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, buf.Bytes(), 0666)
}

func SaveImageJPEG(filename string, m image.Image) error {
	var buf bytes.Buffer
	err := jpeg.Encode(&buf, m, nil)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, buf.Bytes(), 0666)
}
