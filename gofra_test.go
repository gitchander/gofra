package gofra

import (
	"math/rand"
	"testing"
	"time"
)

var seedBench int64 = rand.Int63()

type traceOrbitFn func(Z Complex, n int) int

func benchmarkTraceOrbit(fn traceOrbitFn, b *testing.B) {

	n := 100
	r := rand.New(rand.NewSource(seedBench))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		Z := Complex{
			Re: r.Float64(),
			Im: r.Float64(),
		}.MulScalar(2)

		fn(Z, n)
	}
}

func BenchmarkMandelbrotSample1(b *testing.B) {
	benchmarkTraceOrbit(mandelbrot_traceOrbit1, b)
}

func BenchmarkMandelbrotSample2(b *testing.B) {
	benchmarkTraceOrbit(mandelbrot_traceOrbit2, b)
}

func BenchmarkMandelbrotSample3(b *testing.B) {
	benchmarkTraceOrbit(mandelbrot_traceOrbit3, b)
}

func Benchmark1MandelbrotPow3Sample1(b *testing.B) {
	benchmarkTraceOrbit(mandelbrotPow3_traceOrbit1, b)
}

func Benchmark1MandelbrotPow3Sample2(b *testing.B) {
	benchmarkTraceOrbit(mandelbrotPow3_traceOrbit2, b)
}

func Benchmark2MandelbrotPow3(b *testing.B) {
	benchmarkTraceOrbit(mandelbrotPow3_traceOrbit3, b)
}

func testTraceOrbit(fns []traceOrbitFn, t *testing.T) {

	m := 100
	n := len(fns)
	var in = make([]int, n)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 100000; i++ {

		Z := Complex{
			Re: r.Float64(),
			Im: r.Float64(),
		}.MulScalar(2)

		for i, fn := range fns {
			in[i] = fn(Z, m)
		}

		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if i != j {
					if in[i] != in[j] {
						t.Fatalf("(fn[%d](Z) != fn[%d](Z)) where (Z = %+v)", i, j, Z)
					}
				}
			}
		}
	}
}

func TestTraceOrbitMandelbrot(t *testing.T) {

	fns := []traceOrbitFn{
		mandelbrot_traceOrbit1,
		mandelbrot_traceOrbit2,
		mandelbrot_traceOrbit3,
	}

	testTraceOrbit(fns, t)
}

func TestTraceOrbitMandelbrotPow3(t *testing.T) {

	fns := []traceOrbitFn{
		mandelbrotPow3_traceOrbit1,
		mandelbrotPow3_traceOrbit2,
		mandelbrotPow3_traceOrbit3,
	}

	testTraceOrbit(fns, t)
}

func mandelbrot_traceOrbit1(Z Complex, n int) int {
	C := Z
	for i := 0; i < n; i++ {
		if Z.Norm() > 4.0 {
			return i
		}
		Z = Z.Power(2).Add(C)
	}
	return n
}

func mandelbrot_traceOrbit2(Z Complex, n int) int {
	C := Z
	for i := 0; i < n; i++ {
		if Z.Norm() > 4.0 {
			return i
		}
		Z = Z.Mul(Z).Add(C)
	}
	return n
}

func mandelbrot_traceOrbit3(Z Complex, n int) int {

	var (
		x = Z.Re
		y = Z.Im

		Cx = x
		Cy = y
	)

	var xx, yy, xy float64

	for i := 0; i < n; i++ {

		xx = x * x
		yy = y * y

		if xx+yy > 4.0 {
			return i
		}

		xy = x * y
		y = xy + xy + Cy
		x = xx - yy + Cx
	}

	return n
}

func mandelbrotPow3_traceOrbit1(Z Complex, n int) int {
	C := Z
	for i := 0; i < n; i++ {
		if Z.Norm() > 4.0 {
			return i
		}
		Z = Z.Power(3).Add(C)
	}
	return n
}

func mandelbrotPow3_traceOrbit2(Z Complex, n int) int {
	C := Z
	for i := 0; i < n; i++ {
		if Z.Norm() > 4.0 {
			return i
		}
		Z = Z.Mul(Z).Mul(Z).Add(C)
	}
	return n
}

func mandelbrotPow3_traceOrbit3(Z Complex, n int) int {

	var (
		x = Z.Re
		y = Z.Im

		Cx = x
		Cy = y
	)

	var xx, yy float64

	for i := 0; i < n; i++ {

		xx = x * x
		yy = y * y

		if xx+yy > 4.0 {
			return i
		}

		x = x*(xx-3.0*yy) + Cx
		y = y*(3.0*xx-yy) + Cy
	}

	return n
}

func juliaSet_TraceOrbit1(Z, C Complex, n int) int {
	for i := 0; i < n; i++ {
		if Z.Norm() > 4.0 {
			return i
		}
		Z = Z.Mul(Z).Add(C)
	}
	return n
}

func juliaSet_TraceOrbit2(Z, C Complex, n int) int {

	var (
		x = Z.Re
		y = Z.Im

		Cx = C.Re
		Cy = C.Im
	)

	var xx, yy, xy float64

	for i := 0; i < n; i++ {

		xx = x * x
		yy = y * y

		if xx+yy > 4.0 {
			return i
		}

		xy = x * y
		y = xy + xy + Cy
		x = xx - yy + Cx
	}

	return n
}
