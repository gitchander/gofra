package fcolor

// Lerp - Linear interpolation
// t= [0..1]
// (t == 0) => v0
// (t == 1) => v1
func lerp(v0, v1 float64, t float64) float64 {
	return (1.0-t)*v0 + t*v1
}

func cropFloat64(x float64, min, max float64) float64 {
	if max < min {
		panic("invalid interval")
	}
	if x < min {
		x = min
	}
	if x > max {
		x = max
	}
	return x
}
