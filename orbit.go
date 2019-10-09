package gofra

import (
	"math"

	. "github.com/gitchander/gofra/complex"
)

// type Escaper interface {
// 	Escape() bool
// }

var complexZero = Complex{Re: 0, Im: 0}

type Orbit interface {
	Next() bool
}

type OrbitFactory interface {
	NewOrbit(Complex) Orbit
}

// Escape time algorithm
func TraceOrbit(t Orbit, n int) int {
	for i := 0; i < n; i++ {
		if !t.Next() {
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
	case FormulaMandelbrot:
		of = FactoryMandelbrot{}
	case FormulaMandelbrotPow3:
		of = FactoryMandelbrotPow3{}
	case FormulaMandelbrotPow4:
		of = FactoryMandelbrotPow4{}
	case FormulaMandelbrotPow5:
		of = FactoryMandelbrotPow5{}
	case FormulaMandelbrotPow6:
		of = FactoryMandelbrotPow6{}
	case FormulaJuliaSet:
		if C.Present {
			of = FactoryJuliaSet{C: C.Value}
		} else {
			of = FactoryJuliaSet{C: JuliaSetDefaultConst}
		}
	case FormulaPhoenix:
		if C.Present {
			of = FactoryPhoenix{C: C.Value}
		} else {
			of = FactoryPhoenix{C: PhoenixDefaultConst}
		}
	case FormulaBurningShip:
		of = FactoryBurningShip{}
	case FormulaBurningShipIm:
		of = FactoryBurningShipIm{}
	case FormulaSpider:
		of = FactorySpider{}
	case FormulaTricorn:
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

func newOrbitMandelbrot(C Complex) *orbitMandelbrot {
	return &orbitMandelbrot{
		Z: complexZero,
		C: C,
	}
}

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
	return newOrbitMandelbrot(Z)
}

//------------------------------------------------------------------------------
// MandelbrotPow3

// Z = Z^3 + C
type orbitMandelbrotPow3 struct {
	Z, C Complex
}

var _ Orbit = &orbitMandelbrotPow3{}

func newOrbitMandelbrotPow3(C Complex) *orbitMandelbrotPow3 {
	return &orbitMandelbrotPow3{
		Z: complexZero,
		C: C,
	}
}

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
	return newOrbitMandelbrotPow3(Z)
}

//------------------------------------------------------------------------------
// MandelbrotPow4

// Z = Z^4 + C
type orbitMandelbrotPow4 struct {
	Z, C Complex
}

var _ Orbit = &orbitMandelbrotPow4{}

func newOrbitMandelbrotPow4(C Complex) *orbitMandelbrotPow4 {
	return &orbitMandelbrotPow4{
		Z: complexZero,
		C: C,
	}
}

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
	return newOrbitMandelbrotPow4(Z)
}

//------------------------------------------------------------------------------
// MandelbrotPow5

// Z = Z^5 + C
type orbitMandelbrotPow5 struct {
	Z, C Complex
}

var _ Orbit = &orbitMandelbrotPow5{}

func newOrbitMandelbrotPow5(C Complex) *orbitMandelbrotPow5 {
	return &orbitMandelbrotPow5{
		Z: complexZero,
		C: C,
	}
}

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
	return newOrbitMandelbrotPow5(Z)
}

//------------------------------------------------------------------------------
// MandelbrotPow6

// Z = Z^6 + C
type orbitMandelbrotPow6 struct {
	Z, C Complex
}

var _ Orbit = &orbitMandelbrotPow6{}

func newOrbitMandelbrotPow6(C Complex) *orbitMandelbrotPow6 {
	return &orbitMandelbrotPow6{
		Z: complexZero,
		C: C,
	}
}

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
	return newOrbitMandelbrotPow6(Z)
}

//------------------------------------------------------------------------------
// JuliaSet

type orbitJuliaSet struct {
	Z, C Complex
}

var _ Orbit = &orbitJuliaSet{}

func newOrbitJuliaSet(Z, C Complex) *orbitJuliaSet {
	return &orbitJuliaSet{
		Z: Z,
		C: C,
	}
}

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

var JuliaSetDefaultConst = Complex{
	Re: -0.74543,
	Im: 0.11301,
}

var _ OrbitFactory = FactoryJuliaSet{}

func (f FactoryJuliaSet) NewOrbit(Z Complex) Orbit {
	return newOrbitJuliaSet(Z, f.C)
}

//------------------------------------------------------------------------------
// Phoenix

type orbitPhoenix struct {
	Z, C  Complex
	prevZ Complex
}

var _ Orbit = &orbitPhoenix{}

func newOrbitPhoenix(Z, C Complex) *orbitPhoenix {
	return &orbitPhoenix{
		Z:     Z,
		C:     C,
		prevZ: complexZero,
	}
}

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

var PhoenixDefaultConst = Complex{
	Re: 0.56667,
	Im: -0.5,
}

var _ OrbitFactory = FactoryPhoenix{}

func (f FactoryPhoenix) NewOrbit(Z Complex) Orbit {
	return newOrbitPhoenix(Z, f.C)
}

//------------------------------------------------------------------------------
// BurningShip

// https://en.wikipedia.org/wiki/Burning_Ship_fractal
// http://paulbourke.net/fractals/burnship/

type orbitBurningShip struct {
	Z, C Complex
}

var _ Orbit = &orbitBurningShip{}

func newOrbitBurningShip(C Complex) *orbitBurningShip {
	return &orbitBurningShip{
		Z: complexZero,
		C: C,
	}
}

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

func (FactoryBurningShip) NewOrbit(Z Complex) Orbit {
	return newOrbitBurningShip(Z)
}

//------------------------------------------------------------------------------
// BurningShipIm

type orbitBurningShipIm struct {
	Z, C Complex
}

var _ Orbit = &orbitBurningShipIm{}

func newOrbitBurningShipIm(C Complex) *orbitBurningShipIm {
	return &orbitBurningShipIm{
		Z: complexZero,
		C: C,
	}
}

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
	return newOrbitBurningShipIm(Z)
}

//------------------------------------------------------------------------------
// Spider

type orbitSpider struct {
	Z, C Complex
}

var _ Orbit = &orbitSpider{}

func newOrbitSpider(C Complex) *orbitSpider {
	return &orbitSpider{
		Z: complexZero,
		C: C,
	}
}

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

func (FactorySpider) NewOrbit(Z Complex) Orbit {
	return newOrbitSpider(Z)
}

//------------------------------------------------------------------------------
// Tricorn

type orbitTricorn struct {
	Z, C Complex
}

var _ Orbit = &orbitTricorn{}

func newOrbitTricorn(C Complex) *orbitTricorn {
	return &orbitTricorn{
		Z: complexZero,
		C: C,
	}
}

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
	return newOrbitTricorn(Z)
}
