package math2d

import "math"

type Matrix3x2 [3 * 2]float64

func (m *Matrix3x2) InitIdendity() {
	*m = Matrix3x2{
		1, 0,
		0, 1,
		0, 0,
	}
}

func (m *Matrix3x2) InitTranslate(x, y float64) {
	*m = Matrix3x2{
		1, 0,
		0, 1,
		x, y,
	}
}

func (m *Matrix3x2) InitScale(x, y float64) {
	*m = Matrix3x2{
		x, 0,
		0, y,
		0, 0,
	}
}

func (m *Matrix3x2) InitRotate(angle float64) {
	sin, cos := math.Sincos(angle)
	*m = Matrix3x2{
		cos, -sin,
		sin, cos,
		0, 0,
	}
}

func (m *Matrix3x2) InitReflectAxisX() {
	*m = Matrix3x2{
		1, 0,
		0, -1,
		0, 0,
	}
}

func (m *Matrix3x2) InitReflectAxisY() {
	*m = Matrix3x2{
		-1, 0,
		0, 1,
		0, 0,
	}
}

func (m *Matrix3x2) Translate(x, y float64) {
	m[4] += x
	m[5] += y
}

func (m *Matrix3x2) Scale(x, y float64) {
	m[0] *= x
	m[1] *= y

	m[2] *= x
	m[3] *= y

	m[4] *= x
	m[5] *= y
}

func (m *Matrix3x2) Rotate(angle float64) {
	var r Matrix3x2
	r.InitRotate(angle)
	mat3x2Mul(m[:], r[:], m[:])
}

func (m *Matrix3x2) ReflectAxisX() {
	var n Matrix3x2
	n.InitReflectAxisX()
	mat3x2Mul(m[:], n[:], m[:])
}

func (m *Matrix3x2) ReflectAxisY() {
	var n Matrix3x2
	n.InitReflectAxisY()
	mat3x2Mul(m[:], n[:], m[:])
}

func (m *Matrix3x2) TransformPoint(x, y float64) (tx, ty float64) {
	tx = x*m[0] + y*m[2] + m[4]
	ty = x*m[1] + y*m[3] + m[5]
	return
}

// z = x * y
// (z == x) || (z != x)
// (z != y)

func mat3x2Mul(x, y, z []float64) {

	a, b := x[0], x[1]

	z[0] = a*y[0] + b*y[2]
	z[1] = a*y[1] + b*y[3]

	a, b = x[2], x[3]

	z[2] = a*y[0] + b*y[2]
	z[3] = a*y[1] + b*y[3]

	a, b = x[4], x[5]

	z[4] = a*y[0] + b*y[2]
	z[5] = a*y[1] + b*y[3]
}
