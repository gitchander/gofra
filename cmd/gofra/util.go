package main

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func cropRange(x, min, max int) int {
	if min > max {
		min, max = max, min
	}
	if x < min {
		x = min
	}
	if x > max {
		x = max
	}
	return x
}

func cropRangef(x, min, max float64) float64 {
	if min > max {
		min, max = max, min
	}
	if x < min {
		x = min
	}
	if x > max {
		x = max
	}
	return x
}
