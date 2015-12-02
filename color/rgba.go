package color

type RGBA struct {
	R, G, B, A float64
}

// Alpha blending
// c = a over b
func (a RGBA) Over(b RGBA) (c RGBA) {
	c.A = lerp(b.A, 1.0, a.A)
	c.R = lerp(b.R*b.A, a.R, a.A) / c.A
	c.G = lerp(b.G*b.A, a.G, a.A) / c.A
	c.B = lerp(b.B*b.A, a.B, a.A) / c.A
	return
}
