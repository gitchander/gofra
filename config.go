package gofra

import (
	. "github.com/gitchander/gofra/complex"

	"github.com/gitchander/gofra/math2d"
)

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

type Config struct {
	ImageSize   Size        `json:"image_size"`
	FractalInfo FractalInfo `json:"fractal_info"`
	Calculation Calculation `json:"calculation"`
	Palette     Palette     `json:"palette"`
}

var DefaultConfig = Config{
	ImageSize: Size{
		Width:  512,
		Height: 512,
	},
	FractalInfo: FractalInfo{
		Formula: FormulaMandelbrot,
		Location: Location{
			Center: Complex{
				Re: 0,
				Im: 0,
			},
			Radius:   2,
			AngleDeg: 0,
		},
		Parameters: nil,
	},
	Calculation: Calculation{
		Iterations:   100,
		AntiAliasing: AA_NONE,
	},
	Palette: DefaultPalette,
}

// relative value
// range: [-1 ... +1]
func (p *Config) MoveRelativeLocation(x, y float64) {

	loc := &(p.FractalInfo.Location)

	var m math2d.Matrix
	m.InitIdendity()
	m.Scale(loc.Radius, loc.Radius)
	m.Rotate(degToRad(float64(loc.AngleDeg)))
	m.Translate(loc.Center.Re, loc.Center.Im)

	loc.Center.Re, loc.Center.Im = m.TransformPoint(x, y)
}

func (p *Config) RotateDeg(angle int) {
	loc := &(p.FractalInfo.Location)
	loc.AngleDeg = angleDegNorm(loc.AngleDeg + angle)
}
