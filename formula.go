package gofra

import (
	"encoding"
	"fmt"
)

type Formula int

const (
	FormulaMandelbrot Formula = iota
	FormulaMandelbrotPow3
	FormulaMandelbrotPow4
	FormulaMandelbrotPow5
	FormulaMandelbrotPow6
	FormulaJuliaSet
	FormulaPhoenix
	FormulaBurningShip
	FormulaBurningShipIm
	FormulaSpider
	FormulaTricorn
)

var namesFormula = map[Formula]string{
	FormulaMandelbrot:     "mandelbrot",
	FormulaMandelbrotPow3: "mandelbrot-pow3",
	FormulaMandelbrotPow4: "mandelbrot-pow4",
	FormulaMandelbrotPow5: "mandelbrot-pow5",
	FormulaMandelbrotPow6: "mandelbrot-pow6",
	FormulaJuliaSet:       "julia-set",
	FormulaPhoenix:        "phoenix",
	FormulaBurningShip:    "burning-ship",
	FormulaBurningShipIm:  "burning-ship-im",
	FormulaSpider:         "spider",
	FormulaTricorn:        "tricorn",
}

var valuesFormula = map[string]Formula{
	"mandelbrot":      FormulaMandelbrot,
	"mandelbrot-pow3": FormulaMandelbrotPow3,
	"mandelbrot-pow4": FormulaMandelbrotPow4,
	"mandelbrot-pow5": FormulaMandelbrotPow5,
	"mandelbrot-pow6": FormulaMandelbrotPow6,
	"julia-set":       FormulaJuliaSet,
	"phoenix":         FormulaPhoenix,
	"burning-ship":    FormulaBurningShip,
	"burning-ship-im": FormulaBurningShipIm,
	"spider":          FormulaSpider,
	"tricorn":         FormulaTricorn,
}

func _() {
	var f Formula
	var (
		_ encoding.TextMarshaler   = f
		_ encoding.TextUnmarshaler = &f
	)
}

func (f Formula) String() string {
	value := f
	name, ok := namesFormula[value]
	if ok {
		return name
	}
	return fmt.Sprintf("Formula(%d)", value)
}

func ParseFormula(s string) (Formula, error) {
	name := s
	value, ok := valuesFormula[name]
	if !ok {
		return 0, fmt.Errorf("gofra.ParseFormula: undefined name %q", name)
	}
	return value, nil
}

func (f Formula) MarshalText() (text []byte, err error) {
	value := f
	name, ok := namesFormula[value]
	if !ok {
		return nil, fmt.Errorf("gofra.Formula.MarshalText: undefined value %d", value)
	}
	text = []byte(name)
	return text, nil
}

func (f *Formula) UnmarshalText(text []byte) error {
	name := string(text)
	value, ok := valuesFormula[name]
	if !ok {
		return fmt.Errorf("gofra.Formula.UnmarshalText: undefined name %q", name)
	}
	*f = value
	return nil
}
