package gofra

import (
	"math"

	. "github.com/gitchander/gofra/complex"
)

// ETA - Escape time algorithm
// trace orbit

type OrbitTracer interface {
	Init(Z Complex)
	Next(Z Complex) Complex
	Clone() OrbitTracer
}

func TraceOrbit(ot OrbitTracer, Z Complex, n int) int {
	ot.Init(Z)
	for i := 0; i < n; i++ {
		if Z.Norm() > 4.0 {
			return i
		}
		Z = ot.Next(Z)
	}
	return n
}

/*
// Escape time algorithm

type Orbita interface {
	//Done() bool
	Escape() bool

	Next()
}

func Trace(p Orbita, n int) int {
	for i := 0; i < n; i++ {
		if p.Escape() {
			return i
		}
		p.Next()
	}
	return -1
}
*/

func newOrbitTracer(fi FractalInfo) (t OrbitTracer) {

	var C Complex
	if len(fi.Parameters) > 0 {
		C = fi.Parameters[0]
	}

	switch fi.Formula {
	case FM_MANDELBROT:
		t = &Mandelbrot{}
	case FM_MANDELBROT_POW3:
		t = &MandelbrotPow3{}
	case FM_MANDELBROT_POW4:
		t = &MandelbrotPow4{}
	case FM_MANDELBROT_POW5:
		t = &MandelbrotPow5{}
	case FM_MANDELBROT_POW6:
		t = &MandelbrotPow6{}
	case FM_JULIA_SET:
		t = &JuliaSet{C: C}
	case FM_PHOENIX:
		t = &Phoenix{C: C}
	case FM_BURNING_SHIP:
		t = &BurningShip{}
	case FM_BURNING_SHIP_IM:
		t = &BurningShipIm{}
	case FM_SPIDER:
		t = &Spider{}
	case FM_TRICORN:
		t = &Tricorn{}
	default:
		t = &Mandelbrot{}
	}

	return
}

type Mandelbrot struct {
	C Complex
}

func (f *Mandelbrot) Init(Z Complex) {
	f.C = Z
}

func (f *Mandelbrot) Next(Z Complex) Complex {
	return Z.Mul(Z).Add(f.C)
}

func (f *Mandelbrot) Clone() OrbitTracer {
	cl := *f
	return &cl
}

type MandelbrotPow3 struct {
	C Complex
}

func (f *MandelbrotPow3) Init(Z Complex) {
	f.C = Z
}

func (f *MandelbrotPow3) Next(Z Complex) Complex {
	return Z.Mul(Z).Mul(Z).Add(f.C)
}

func (f *MandelbrotPow3) Clone() OrbitTracer {
	cl := *f
	return &cl
}

type MandelbrotPow4 struct {
	C Complex
}

func (f *MandelbrotPow4) Init(Z Complex) {
	f.C = Z
}

func (f *MandelbrotPow4) Next(Z Complex) Complex {
	ZZ := Z.Mul(Z)
	return ZZ.Mul(ZZ).Add(f.C)
}

func (f *MandelbrotPow4) Clone() OrbitTracer {
	cl := *f
	return &cl
}

type MandelbrotPow5 struct {
	C Complex
}

func (f *MandelbrotPow5) Init(Z Complex) {
	f.C = Z
}

func (f *MandelbrotPow5) Next(Z Complex) Complex {
	ZZ := Z.Mul(Z)
	return ZZ.Mul(ZZ).Mul(Z).Add(f.C)
}

func (f *MandelbrotPow5) Clone() OrbitTracer {
	cl := *f
	return &cl
}

type MandelbrotPow6 struct {
	C Complex
}

func (f *MandelbrotPow6) Init(Z Complex) {
	f.C = Z
}

func (f *MandelbrotPow6) Next(Z Complex) Complex {
	ZZZ := Z.Mul(Z).Mul(Z)
	return ZZZ.Mul(ZZZ).Add(f.C)
}

func (f *MandelbrotPow6) Clone() OrbitTracer {
	cl := *f
	return &cl
}

type JuliaSet struct {
	C Complex
}

// C := Complex{Re: -0.74543, Im: 0.11301}

func (f *JuliaSet) Init(Z Complex) {}

func (f *JuliaSet) Next(Z Complex) Complex {
	return Z.Mul(Z).Add(f.C)
}

func (f *JuliaSet) Clone() OrbitTracer {
	cl := *f
	return &cl
}

type Phoenix struct {
	C     Complex
	prevZ Complex
}

var PhoenixDefault = Phoenix{
	C: Complex{
		Re: 0.56667,
		Im: -0.5,
	},
}

func (f *Phoenix) Init(Z Complex) {
	f.prevZ = Complex{
		Re: 0,
		Im: 0,
	}
}

func (f *Phoenix) Next(Z Complex) Complex {
	tempZ := Z
	Z = Z.Mul(Z).AddScalar(f.C.Re).Add(f.prevZ.MulScalar(f.C.Im))
	f.prevZ = tempZ
	return Z
}

func (f *Phoenix) Clone() OrbitTracer {
	cl := *f
	return &cl
}

type BurningShip struct {
	C Complex
}

func (f *BurningShip) Init(Z Complex) {
	f.C = Z
}

func (f *BurningShip) Next(Z Complex) Complex {
	Z.Re = math.Abs(Z.Re)
	Z.Im = math.Abs(Z.Im)
	return Z.Mul(Z).Add(f.C)
}

func (f *BurningShip) Clone() OrbitTracer {
	cl := *f
	return &cl
}

type BurningShipIm struct {
	C Complex
}

func (f *BurningShipIm) Init(Z Complex) {
	f.C = Z
}

func (f *BurningShipIm) Next(Z Complex) Complex {
	Z.Im = math.Abs(Z.Im)
	return Z.Mul(Z).Add(f.C)
}

func (f *BurningShipIm) Clone() OrbitTracer {
	cl := *f
	return &cl
}

type Spider struct {
	C Complex
}

func (f *Spider) Init(Z Complex) {
	f.C = Z
}

func (f *Spider) Next(Z Complex) Complex {
	Z = Z.Mul(Z).Add(f.C)
	f.C = f.C.MulScalar(0.5).Add(Z)
	return Z
}

func (f *Spider) Clone() OrbitTracer {
	cl := *f
	return &cl
}

type Tricorn struct {
	C Complex
}

func (f *Tricorn) Init(Z Complex) {
	f.C = Z
}

func (f *Tricorn) Next(Z Complex) Complex {
	Z = Z.Conjugate()
	return Z.Mul(Z).Add(f.C)
}

func (f *Tricorn) Clone() OrbitTracer {
	cl := *f
	return &cl
}
