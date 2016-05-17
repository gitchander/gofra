package gofra

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"strings"

	. "github.com/gitchander/gofra/complex"
	"github.com/gitchander/gofra/fcolor"
	"github.com/gitchander/gofra/mth2d"
)

type Formula int

const (
	FM_MANDELBROT Formula = iota
	FM_MANDELBROT_POW3
	FM_MANDELBROT_POW4
	FM_MANDELBROT_POW5
	FM_MANDELBROT_POW6
	FM_JULIA_SET
	FM_PHOENIX
	FM_BURNING_SHIP
	FM_BURNING_SHIP_IM
	FM_SPIDER
	FM_TRICORN
)

const (
	str_FM_MANDELBROT      = "Mandelbrot"
	str_FM_MANDELBROT_POW3 = "Mandelbrot^3"
	str_FM_MANDELBROT_POW4 = "Mandelbrot^4"
	str_FM_MANDELBROT_POW5 = "Mandelbrot^5"
	str_FM_MANDELBROT_POW6 = "Mandelbrot^6"
	str_FM_JULIA_SET       = "Julia Set"
	str_FM_PHOENIX         = "Phoenix"
	str_FM_BURNING_SHIP    = "Burning Ship"
	str_FM_BURNING_SHIP_IM = "Burning Ship Im"
	str_FM_SPIDER          = "Spider"
	str_FM_TRICORN         = "Tricorn"
)

var name_Formula = map[Formula]string{
	FM_MANDELBROT:      str_FM_MANDELBROT,
	FM_MANDELBROT_POW3: str_FM_MANDELBROT_POW3,
	FM_MANDELBROT_POW4: str_FM_MANDELBROT_POW4,
	FM_MANDELBROT_POW5: str_FM_MANDELBROT_POW5,
	FM_MANDELBROT_POW6: str_FM_MANDELBROT_POW6,
	FM_JULIA_SET:       str_FM_JULIA_SET,
	FM_PHOENIX:         str_FM_PHOENIX,
	FM_BURNING_SHIP:    str_FM_BURNING_SHIP,
	FM_BURNING_SHIP_IM: str_FM_BURNING_SHIP_IM,
	FM_SPIDER:          str_FM_SPIDER,
	FM_TRICORN:         str_FM_TRICORN,
}

var value_Formula = map[string]Formula{
	str_FM_MANDELBROT:      FM_MANDELBROT,
	str_FM_MANDELBROT_POW3: FM_MANDELBROT_POW3,
	str_FM_MANDELBROT_POW4: FM_MANDELBROT_POW4,
	str_FM_MANDELBROT_POW5: FM_MANDELBROT_POW5,
	str_FM_MANDELBROT_POW6: FM_MANDELBROT_POW6,
	str_FM_JULIA_SET:       FM_JULIA_SET,
	str_FM_PHOENIX:         FM_PHOENIX,
	str_FM_BURNING_SHIP:    FM_BURNING_SHIP,
	str_FM_BURNING_SHIP_IM: FM_BURNING_SHIP_IM,
	str_FM_SPIDER:          FM_SPIDER,
	str_FM_TRICORN:         FM_TRICORN,
}

func (f *Formula) MarshalJSON() ([]byte, error) {

	s, ok := name_Formula[*f]
	if !ok {
		return nil, errors.New("Forula.MarshalJSON")
	}

	return json.Marshal(s)
}

func (f *Formula) UnmarshalJSON(data []byte) error {

	var s string

	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	v, ok := value_Formula[s]
	if !ok {
		return errors.New("Formula.UnmarshalJSON")
	}

	*f = v

	return nil
}

type AntiAliasing int

const (
	AA_NONE AntiAliasing = iota
	AA_4X
	AA_9X
	AA_16X
)

var key_AntiAliasing = map[AntiAliasing]string{
	AA_NONE: "NONE",
	AA_4X:   "4X",
	AA_9X:   "9X",
	AA_16X:  "16X",
}

var val_AntiAliasing = map[string]AntiAliasing{
	"NONE": AA_NONE,
	"4X":   AA_4X,
	"9X":   AA_9X,
	"16X":  AA_16X,
}

func (aa *AntiAliasing) MarshalJSON() ([]byte, error) {

	s, ok := key_AntiAliasing[*aa]
	if !ok {
		return nil, errors.New("AntiAliasing.MarshalJSON")
	}

	return json.Marshal(s)
}

func (aa *AntiAliasing) UnmarshalJSON(data []byte) error {

	var s string

	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	s = strings.ToUpper(s)
	v, ok := val_AntiAliasing[s]
	if !ok {
		return errors.New("AntiAliasing.UnmarshalJSON")
	}

	*aa = v

	return nil
}

type Location struct {
	Center Complex
	Radius float64
	Angle  float64
}

type FractalInfo struct {
	Formula    Formula
	Location   Location
	Parameters []Complex
}

type Calculation struct {
	Iterations   int
	AntiAliasing AntiAliasing
}

type Size struct {
	Width  int
	Height int
}

type Parameters struct {
	ImageSize   Size
	FractalInfo FractalInfo
	Calculation Calculation
	Palette     Palette
}

var DefaultParameters = Parameters{
	ImageSize: Size{
		Width:  512,
		Height: 512,
	},
	FractalInfo: FractalInfo{
		Formula:    FM_MANDELBROT,
		Parameters: nil,
		Location: Location{
			Center: Complex{0, 0},
			Radius: 2,
			Angle:  0,
		},
	},
	Calculation: Calculation{
		Iterations:   100,
		AntiAliasing: AA_NONE,
	},
	Palette: Palette{
		Colors: []fcolor.RGB{
			fcolor.RGB{1, 1, 1},
			fcolor.RGB{0, 0, 0},
		},
		SpaceColor: fcolor.RGB{0, 0, 0},
		Period:     30,
	},
}

func (p *Parameters) LoadFromFile(fileName string) error {

	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	r := bufio.NewReader(file)
	d := json.NewDecoder(r)

	if err = d.Decode(p); err != nil {
		return err
	}

	return nil
}

func (p *Parameters) SaveToFile(fileName string) error {

	data, err := json.MarshalIndent(p, "", "\t")
	if err != nil {
		return err
	}

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err = file.Write(data); err != nil {
		return err
	}

	return nil
}

// relative value
func (p *Parameters) MoveRelativeLocation(x, y float64) {

	normRelativeValue := func(x float64) float64 {
		if x < -1 {
			x = -1
		}
		if x > 1 {
			x = 1
		}
		return x
	}

	x = normRelativeValue(x)
	y = normRelativeValue(y)

	loc := &(p.FractalInfo.Location)

	var m mth2d.Matrix
	m.InitIdendity()
	m.Scale(loc.Radius, loc.Radius)
	m.Rotate(degToRad(loc.Angle))
	m.Translate(loc.Center.Re, loc.Center.Im)

	loc.Center.Re, loc.Center.Im = m.TransformPoint(x, y)
}
