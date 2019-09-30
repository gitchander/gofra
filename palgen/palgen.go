package palgen

import (
	"math"

	"github.com/gitchander/gofra/fcolor"
)

// http://www.iquilezles.org/www/articles/palettes/palettes.htm

// t = [0..1]
func ColorByParams(p Params, t float64) fcolor.RGB {
	v := palette(p.A, p.B, p.C, p.D, t)
	return fcolor.RGB{
		R: clampFloat64(v[0], 0, 1),
		G: clampFloat64(v[1], 0, 1),
		B: clampFloat64(v[2], 0, 1),
	}
}

func palette(a, b, c, d Vec3, t float64) Vec3 {

	// a + b * cos( 2*pi * (c*t + d) )

	const twoPi = 2 * math.Pi

	angle := c.MulScalar(t).Add(d).MulScalar(twoPi)

	cos := CosVec3(angle)

	return a.Add(b.Mul(cos))
}

func clampFloat64(x float64, min, max float64) float64 {
	if max < min { // It is empty intarval
		return 0 // Default value is zero
	}
	if x < min {
		x = min
	}
	if x > max {
		x = max
	}
	return x
}
