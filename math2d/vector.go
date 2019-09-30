package math2d

type Vector3 [3]float64

func Vector3XY(x, y float64) Vector3 {
	return Vector3{x, y, 1}
}

func (v Vector3) XY() (x, y float64) {
	x = v[0]
	y = v[1]
	return x, y
}

func (v Vector3) Norm() Vector3 {

	m := v[2]
	if m == 0 {
		panic("invalid Vector3: (v[2] = 0)")
	}
	if m == 1 {
		return v
	}

	v[0] /= m
	v[1] /= m
	v[2] = 1

	return v
}
