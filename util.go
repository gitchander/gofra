package gofra

import "math"

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// nearest dividers
func nearDivs(x int) (a, b int) {
	for d := 1; d*d <= x; d++ {
		if (x % d) == 0 {
			a = d
			b = x / d
		}
	}
	return
}

func degToRad(deg float64) float64 {
	return deg * math.Pi / 180
}

func mod(a, b int) int {
	m := a % b
	if m < 0 {
		m += b
	}
	return m
}

// [0..360)
func angleDegNorm(a int) int {
	return mod(a, 360)
}
