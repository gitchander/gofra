package main

import (
	"bytes"
	"image"
	"image/png"
	"io/ioutil"
	"log"

	"github.com/gitchander/gofra/palgen"
)

func main() {
	size := image.Point{512, 32}
	m := NewRGBASize(size)

	//palgen.Draw1(m)
	palgen.Draw2(m)

	err := SaveImagePNG("palette.png", m)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func NewRGBASize(size image.Point) *image.RGBA {
	r := image.Rectangle{Max: size}
	return image.NewRGBA(r)
}

func SaveImagePNG(filename string, m image.Image) error {
	var buf bytes.Buffer
	err := png.Encode(&buf, m)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, buf.Bytes(), 0666)
}
