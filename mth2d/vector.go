package mth2d

type Vector [3]float64

func (v *Vector) set_XY(x, y float64) {
	v[0] = x
	v[1] = y
	v[2] = 1
}

func (v Vector) get_XY() (x, y float64) {

	v.norm()

	x = v[0]
	y = v[1]

	return
}

func (v *Vector) norm() {

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
