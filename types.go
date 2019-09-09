package gofra

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"

	. "github.com/gitchander/gofra/complex"

	"github.com/gitchander/gofra/fcolor"
	"github.com/gitchander/gofra/math2d"
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

func (f Formula) MarshalJSON() ([]byte, error) {

	s, ok := name_Formula[f]
	if !ok {
		return nil, errors.New("Formula.MarshalJSON")
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

type Location struct {
	Center   Complex `json:"center"`
	Radius   float64 `json:"radius"`
	AngleDeg int     `json:"angle_deg"`
}

type FractalInfo struct {
	Formula    Formula   `json:"formula"`
	Location   Location  `json:"location"`
	Parameters []Complex `json:"parameters"`
}

type Calculation struct {
	Iterations   int          `json:"iterations"`
	AntiAliasing AntiAliasing `json:"anti_aliasing"`
}

type Size struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type Parameters struct {
	ImageSize   Size        `json:"image_size"`
	FractalInfo FractalInfo `json:"fractal_info"`
	Calculation Calculation `json:"calculation"`
	Palette     Palette     `json:"palette"`
}

var DefaultParameters = Parameters{
	ImageSize: Size{
		Width:  512,
		Height: 512,
	},
	FractalInfo: FractalInfo{
		Formula: FM_MANDELBROT,
		Location: Location{
			Center:   Complex{0, 0},
			Radius:   2,
			AngleDeg: 0,
		},
		Parameters: nil,
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
// range: [-1 ... +1]
func (p *Parameters) MoveRelativeLocation(x, y float64) {

	loc := &(p.FractalInfo.Location)

	var m math2d.Matrix
	m.InitIdendity()
	m.Scale(loc.Radius, loc.Radius)
	m.Rotate(degToRad(float64(loc.AngleDeg)))
	m.Translate(loc.Center.Re, loc.Center.Im)

	loc.Center.Re, loc.Center.Im = m.TransformPoint(x, y)
}
