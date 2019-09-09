package math2d

type vector3 [3]float64

func (v *vector3) setXY(x, y float64) {
	v[0] = x
	v[1] = y
	v[2] = 1
}

func (v vector3) getXY() (x, y float64) {
	v.norm()
	x = v[0]
	y = v[1]
	return
}

func (v *vector3) norm() {
	if m := v[2]; m != 1 {
		if m == 0 {
			return
		}
		inv_m := 1 / m
		v[0] *= inv_m
		v[1] *= inv_m
		v[2] = 1
	}
}
