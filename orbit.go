package gofra

import (
	"math"

	. "github.com/gitchander/gofra/complex"
)

// type Escaper interface {
// 	Escape() bool
// }

type Orbit interface {
	Next() bool
}

type OrbitFactory interface {
	NewOrbit(Complex) Orbit
}

// Escape time algorithm
func TraceOrbit(o Orbit, n int) int {
	for i := 0; i < n; i++ {
		if !o.Next() {
			return i
		}
	}
	return -1
}

//------------------------------------------------------------------------------
func newOrbitFactory(fi FractalInfo) (of OrbitFactory) {

	type OptComplex struct {
		Present bool
		Value   Complex
	}

	var C OptComplex
	if len(fi.Parameters) > 0 {
		C = OptComplex{
			Present: true,
			Value:   fi.Parameters[0],
		}
	}

	switch fi.Formula {
	case FM_MANDELBROT:
		of = FactoryMandelbrot{}
	case FM_MANDELBROT_POW3:
		of = FactoryMandelbrotPow3{}
	case FM_MANDELBROT_POW4:
		of = FactoryMandelbrotPow4{}
	case FM_MANDELBROT_POW5:
		of = FactoryMandelbrotPow5{}
	case FM_MANDELBROT_POW6:
		of = FactoryMandelbrotPow6{}
	case FM_JULIA_SET:
		if C.Present {
			of = FactoryJuliaSet{C: C.Value}
		} else {
			of = FactoryJuliaSet{C: JuliaSetDefaultConst}
		}
	case FM_PHOENIX:
		if C.Present {
			of = FactoryPhoenix{C: C.Value}
		} else {
			of = FactoryPhoenix{C: PhoenixDefaultConst}
		}
	case FM_BURNING_SHIP:
		of = FactoryBurningShip{}
	case FM_BURNING_SHIP_IM:
		of = FactoryBurningShipIm{}
	case FM_SPIDER:
		of = FactorySpider{}
	case FM_TRICORN:
		of = FactoryTricorn{}
	default:
		of = FactoryMandelbrot{}
	}

	return
}

//------------------------------------------------------------------------------
// Mandelbrot

// Z = Z^2 + C
type orbitMandelbrot struct {
	Z, C Complex
}

var _ Orbit = &orbitMandelbrot{}

func (p *orbitMandelbrot) Next() bool {
	Z := p.Z
	if Z.Norm() > 4.0 {
		return false
	}
	p.Z = Z.Mul(Z).Add(p.C)
	return true
}

type FactoryMandelbrot struct{}

var _ OrbitFactory = FactoryMandelbrot{}

func (FactoryMandelbrot) NewOrbit(Z Complex) Orbit {
	return &orbitMandelbrot{
		Z: Z,
		C: Z,
	}
}

//------------------------------------------------------------------------------
// MandelbrotPow3

// Z = Z^3 + C
type orbitMandelbrotPow3 struct {
	Z, C Complex
}

var _ Orbit = &orbitMandelbrotPow3{}

func (p *orbitMandelbrotPow3) Next() bool {
	Z := p.Z
	if Z.Norm() > 4.0 {
		return false
	}
	p.Z = Z.Mul(Z).Mul(Z).Add(p.C)
	return true
}

type FactoryMandelbrotPow3 struct{}

var _ OrbitFactory = FactoryMandelbrotPow3{}

func (FactoryMandelbrotPow3) NewOrbit(Z Complex) Orbit {
	return &orbitMandelbrotPow3{
		Z: Z,
		C: Z,
	}
}

//------------------------------------------------------------------------------
// MandelbrotPow4

// Z = Z^4 + C
type orbitMandelbrotPow4 struct {
	Z, C Complex
}

var _ Orbit = &orbitMandelbrotPow4{}

func (p *orbitMandelbrotPow4) Next() bool {
	Z := p.Z
	if Z.Norm() > 4.0 {
		return false
	}
	ZZ := Z.Mul(Z)
	p.Z = ZZ.Mul(ZZ).Add(p.C)
	return true
}

type FactoryMandelbrotPow4 struct{}

var _ OrbitFactory = FactoryMandelbrotPow4{}

func (FactoryMandelbrotPow4) NewOrbit(Z Complex) Orbit {
	return &orbitMandelbrotPow4{
		Z: Z,
		C: Z,
	}
}

//------------------------------------------------------------------------------
// MandelbrotPow5

// Z = Z^5 + C
type orbitMandelbrotPow5 struct {
	Z, C Complex
}

var _ Orbit = &orbitMandelbrotPow5{}

func (p *orbitMandelbrotPow5) Next() bool {
	Z := p.Z
	if Z.Norm() > 4.0 {
		return false
	}
	ZZ := Z.Mul(Z)
	p.Z = ZZ.Mul(ZZ).Mul(Z).Add(p.C)
	return true
}

type FactoryMandelbrotPow5 struct{}

var _ OrbitFactory = FactoryMandelbrotPow5{}

func (FactoryMandelbrotPow5) NewOrbit(Z Complex) Orbit {
	return &orbitMandelbrotPow5{
		Z: Z,
		C: Z,
	}
}

//------------------------------------------------------------------------------
// MandelbrotPow6

// Z = Z^6 + C
type orbitMandelbrotPow6 struct {
	Z, C Complex
}

var _ Orbit = &orbitMandelbrotPow6{}

func (p *orbitMandelbrotPow6) Next() bool {
	Z := p.Z
	if Z.Norm() > 4.0 {
		return false
	}
	ZZZ := Z.Mul(Z).Mul(Z)
	p.Z = ZZZ.Mul(ZZZ).Add(p.C)
	return true
}

type FactoryMandelbrotPow6 struct{}

var _ OrbitFactory = FactoryMandelbrotPow6{}

func (FactoryMandelbrotPow6) NewOrbit(Z Complex) Orbit {
	return &orbitMandelbrotPow6{
		Z: Z,
		C: Z,
	}
}

//------------------------------------------------------------------------------
// JuliaSet

type orbitJuliaSet struct {
	Z, C Complex
}

var _ Orbit = &orbitJuliaSet{}

func (p *orbitJuliaSet) Next() bool {
	Z := p.Z
	if Z.Norm() > 4.0 {
		return false
	}
	p.Z = Z.Mul(Z).Add(p.C)
	return true
}

type FactoryJuliaSet struct {
	C Complex
}

var JuliaSetDefaultConst = Complex{Re: -0.74543, Im: 0.11301}

//---------------------------------------
// "parameters": [
// 	{
// 		"re": -0.74543,
// 		"im": 0.11301
// 	}
// ]
//---------------------------------------

var _ OrbitFactory = FactoryJuliaSet{}

func (f FactoryJuliaSet) NewOrbit(Z Complex) Orbit {
	return &orbitJuliaSet{
		Z: Z,
		C: f.C,
	}
}

//------------------------------------------------------------------------------
// Phoenix

type orbitPhoenix struct {
	Z, C  Complex
	prevZ Complex
}

var _ Orbit = &orbitPhoenix{}

func (p *orbitPhoenix) Next() bool {
	Z := p.Z
	if Z.Norm() > 4.0 {
		return false
	}
	p.Z = Z.Mul(Z).AddScalar(p.C.Re).Add(p.prevZ.MulScalar(p.C.Im))
	p.prevZ = Z
	return true
}

type FactoryPhoenix struct {
	C Complex
}

var PhoenixDefaultConst = Complex{Re: 0.56667, Im: -0.5}

//---------------------------------------
// "parameters": [
// 	{
// 		"re": 0.56667,
// 		"im": -0.5
// 	}
// ]
//---------------------------------------

var _ OrbitFactory = FactoryPhoenix{}

func (f FactoryPhoenix) NewOrbit(Z Complex) Orbit {
	return &orbitPhoenix{
		Z:     Z,
		C:     f.C,
		prevZ: Complex{Re: 0, Im: 0},
	}
}

//------------------------------------------------------------------------------
// BurningShip

// https://en.wikipedia.org/wiki/Burning_Ship_fractal
// http://paulbourke.net/fractals/burnship/

type orbitBurningShip struct {
	Z, C Complex
}

var _ Orbit = &orbitBurningShip{}

func (p *orbitBurningShip) Next() bool {
	Z := p.Z
	if Z.Norm() > 4.0 {
		return false
	}
	Z.Re = math.Abs(Z.Re)
	Z.Im = math.Abs(Z.Im)
	p.Z = Z.Mul(Z).Add(p.C)
	return true
}

type FactoryBurningShip struct{}

var _ OrbitFactory = FactoryBurningShip{}

func (f FactoryBurningShip) NewOrbit(Z Complex) Orbit {
	return &orbitBurningShip{
		Z: Z,
		C: Z,
	}
}

//------------------------------------------------------------------------------
// BurningShipIm

type orbitBurningShipIm struct {
	Z, C Complex
}

var _ Orbit = &orbitBurningShipIm{}

func (p *orbitBurningShipIm) Next() bool {
	Z := p.Z
	if Z.Norm() > 4.0 {
		return false
	}
	Z.Im = math.Abs(Z.Im)
	p.Z = Z.Mul(Z).Add(p.C)
	return true
}

type FactoryBurningShipIm struct{}

var _ OrbitFactory = FactoryBurningShipIm{}

func (f FactoryBurningShipIm) NewOrbit(Z Complex) Orbit {
	return &orbitBurningShipIm{
		Z: Z,
		C: Z,
	}
}

//------------------------------------------------------------------------------
// Spider

type orbitSpider struct {
	Z, C Complex
}

var _ Orbit = &orbitSpider{}

func (p *orbitSpider) Next() bool {
	Z := p.Z
	if Z.Norm() > 4.0 {
		return false
	}
	p.Z = Z.Mul(Z).Add(p.C)
	p.C = p.C.MulScalar(0.5).Add(p.Z)
	return true
}

type FactorySpider struct{}

var _ OrbitFactory = FactorySpider{}

func (f FactorySpider) NewOrbit(Z Complex) Orbit {
	return &orbitSpider{
		Z: Z,
		C: Z,
	}
}

//------------------------------------------------------------------------------
// Tricorn

type orbitTricorn struct {
	Z, C Complex
}

var _ Orbit = &orbitTricorn{}

func (p *orbitTricorn) Next() bool {
	Z := p.Z
	if Z.Norm() > 4.0 {
		return false
	}
	Z = Z.Conjugate()
	p.Z = Z.Mul(Z).Add(p.C)
	return true
}

type FactoryTricorn struct{}

var _ OrbitFactory = FactoryTricorn{}

func (f FactoryTricorn) NewOrbit(Z Complex) Orbit {
	return &orbitTricorn{
		Z: Z,
		C: Z,
	}
}
