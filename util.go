package gofra

import "math"

func min(a, b int) int {
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
