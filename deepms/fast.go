package main

import (
	"fmt"
	"image"
	"image/color"

	"github.com/gitchander/gofra"
)

func fastTest() error {

	im := gofra.NewImageRGBA(256, 256)

	err := render(im)
	if err != nil {
		return err
	}

	err = gofra.ImageSaveToPNG(im, "./test.png")
	if err != nil {
		return err
	}

	return nil
}

func fastTest11() error {

	render(nil)

	return nil
}

func render(im *image.RGBA) error {

	n := 100
	Center := gofra.Complex{
		Re: -1.8594017028808594,
		Im: 0.001809072494506836,
	}
	Radius := 3.1789143880208335e-06

	var X = make([]gofra.Complex, n)
	var A = make([]gofra.Complex, n)
	var B = make([]gofra.Complex, n)
	var C = make([]gofra.Complex, n)

	X[0] = Center
	A[0] = gofra.Complex{1, 0}
	B[0] = gofra.Complex{0, 0}
	C[0] = gofra.Complex{0, 0}

	for i := 0; i < n-1; i++ {
		A[i+1] = X[i].Mul(A[i]).MulScalar(2).AddScalar(1)
		B[i+1] = X[i].Mul(B[i]).MulScalar(2).Add(A[i].Mul(A[i]))
		C[i+1] = X[i].Mul(C[i]).MulScalar(2).Add(A[i].Mul(B[i]).MulScalar(2))

		fmt.Println(A[i+1], B[i+1], C[i+1])
	}

	var (
		nX = im.Rect.Dx()
		nY = im.Rect.Dy()

		pixelWidth = 2 * Radius / float64(min(nX, nY))

		Start = gofra.Complex{
			Re: (Center.Re - float64(nX)*pixelWidth/2),
			Im: (Center.Im - float64(nY)*pixelWidth/2),
		}
	)

	var Z gofra.Complex

	Z.Im = Start.Im
	for yi := 0; yi < nY; yi++ {
		Z.Re = Start.Re
		for xi := 0; xi < nX; xi++ {

			//-------------------------------

			d := Z.Sub(X[0])
			dd := d.Mul(d)
			ddd := dd.Mul(d)

			i := 0
			for i < n {

				di := A[i].Mul(d).Add(B[i].Mul(dd).Add(C[i].Mul(ddd)))

				Y := di.Add(X[i])

				if Y.Norm() > 4 {
					break
				}

				i++
			}

			//i := mset(Z, n)

			//-------------------------------

			if i%2 == 0 {
				im.Set(xi, yi, color.Black)
			} else {
				im.Set(xi, yi, color.White)
			}

			Z.Re += pixelWidth
		}
		Z.Im += pixelWidth
	}

	return nil
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func mset(Z gofra.Complex, n int) int {
	C := Z
	for i := 0; i < n; i++ {
		if Z.Norm() > 4 {
			return i
		}
		Z = Z.Mul(Z).Add(C)
	}
	return n
}
