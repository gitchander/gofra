package math2d

import "math"

type Matrix3x3 [3 * 3]float64

func (m *Matrix3x3) InitIdendity() {
	*m = Matrix3x3{
		1, 0, 0,
		0, 1, 0,
		0, 0, 1,
	}
}

func (m *Matrix3x3) InitTranslate(x, y float64) {
	*m = Matrix3x3{
		1, 0, 0,
		0, 1, 0,
		x, y, 1,
	}
}

func (m *Matrix3x3) InitScale(x, y float64) {
	*m = Matrix3x3{
		x, 0, 0,
		0, y, 0,
		0, 0, 1,
	}
}

func (m *Matrix3x3) InitRotate(angle float64) {
	sin, cos := math.Sincos(angle)
	*m = Matrix3x3{
		cos, -sin, 0,
		sin, cos, 0,
		0, 0, 1,
	}
}

func (m *Matrix3x3) InitReflectAxisX() {
	*m = Matrix3x3{
		1, 0, 0,
		0, -1, 0,
		0, 0, 1,
	}
}

func (m *Matrix3x3) InitReflectAxisY() {
	*m = Matrix3x3{
		-1, 0, 0,
		0, 1, 0,
		0, 0, 1,
	}
}

func (m *Matrix3x3) Translate(x, y float64) {

	temp := m[2]
	m[0] += x * temp
	m[1] += y * temp

	temp = m[5]
	m[3] += x * temp
	m[4] += y * temp

	temp = m[8]
	m[6] += x * temp
	m[7] += y * temp
}

func (m *Matrix3x3) Scale(x, y float64) {

	m[0] *= x
	m[1] *= y

	m[3] *= x
	m[4] *= y

	m[6] *= x
	m[7] *= y
}

func (m *Matrix3x3) Rotate(angle float64) {
	var n Matrix3x3
	n.InitRotate(angle)
	mat3x3Mul(m[:], n[:], m[:])
}

func (m *Matrix3x3) ReflectAxisX() {
	var n Matrix3x3
	n.InitReflectAxisX()
	mat3x3Mul(m[:], n[:], m[:])
}

func (m *Matrix3x3) ReflectAxisY() {
	var n Matrix3x3
	n.InitReflectAxisY()
	mat3x3Mul(m[:], n[:], m[:])
}

// matrix * vector
func (m *Matrix3x3) mulVectorL(x, y float64) (tx, ty float64) {
	v := Vector3XY(x, y)
	var w Vector3
	matrix3x3MulVector3(m[:], v[:], w[:])
	tx, ty = w.XY()
	return
}

// vector * matrix
func (m *Matrix3x3) mulVectorR(x, y float64) (tx, ty float64) {
	v := Vector3XY(x, y)
	var w Vector3
	vector3MulMatrix3x3(v[:], m[:], w[:])
	tx, ty = w.XY()
	return
}

func (m *Matrix3x3) TransformPoint(x, y float64) (tx, ty float64) {
	tx, ty = m.mulVectorR(x, y)
	return
}

func (m *Matrix3x3) Invert() {
	var i Matrix3x3
	invert(m[:], i[:])
	*m = i
}

// z = x * y
func mat3x3Mul(x, y, z []float64) {

	a, b, c := x[0], x[1], x[2]

	z[0] = a*y[0] + b*y[3] + c*y[6]
	z[1] = a*y[1] + b*y[4] + c*y[7]
	z[2] = a*y[2] + b*y[5] + c*y[8]

	a, b, c = x[3], x[4], x[5]

	z[3] = a*y[0] + b*y[3] + c*y[6]
	z[4] = a*y[1] + b*y[4] + c*y[7]
	z[5] = a*y[2] + b*y[5] + c*y[8]

	a, b, c = x[6], x[7], x[8]

	z[6] = a*y[0] + b*y[3] + c*y[6]
	z[7] = a*y[1] + b*y[4] + c*y[7]
	z[8] = a*y[2] + b*y[5] + c*y[8]
}

// w = v * m
func vector3MulMatrix3x3(v, m, w []float64) {

	a, b, c := v[0], v[1], v[2]

	w[0] = a*m[0] + b*m[3] + c*m[6]
	w[1] = a*m[1] + b*m[4] + c*m[7]
	w[2] = a*m[2] + b*m[5] + c*m[8]
}

// w = m * v
func matrix3x3MulVector3(m, v, w []float64) {

	a, b, c := v[0], v[1], v[2]

	w[0] = a*m[0] + b*m[1] + c*m[2]
	w[1] = a*m[3] + b*m[4] + c*m[5]
	w[2] = a*m[6] + b*m[7] + c*m[8]
}

func det(m []float64) (d float64) {

	d += m[0] * (m[4]*m[8] - m[5]*m[7])
	d -= m[1] * (m[3]*m[8] - m[5]*m[6])
	d += m[2] * (m[3]*m[7] - m[4]*m[6])

	return
}

func transpose(m []float64) {
	m[1], m[3] = m[3], m[1]
	m[2], m[6] = m[6], m[2]
	m[5], m[7] = m[7], m[5]
}

func invert(n, m []float64) {

	detInv := 1.0 / det(n)

	m[0] = +detInv * (n[4]*n[8] - n[5]*n[7])
	m[1] = -detInv * (n[3]*n[8] - n[5]*n[6])
	m[2] = +detInv * (n[3]*n[7] - n[4]*n[6])

	m[3] = -detInv * (n[1]*n[8] - n[2]*n[7])
	m[4] = +detInv * (n[0]*n[8] - n[2]*n[6])
	m[5] = -detInv * (n[0]*n[7] - n[1]*n[6])

	m[6] = +detInv * (n[1]*n[5] - n[2]*n[4])
	m[7] = -detInv * (n[0]*n[5] - n[2]*n[3])
	m[8] = +detInv * (n[0]*n[4] - n[1]*n[3])

	transpose(m)
}
