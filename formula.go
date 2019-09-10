package gofra

import (
	"encoding/json"
	"fmt"
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
	name_MANDELBROT      = "mandelbrot"
	name_MANDELBROT_POW3 = "mandelbrot-pow3"
	name_MANDELBROT_POW4 = "mandelbrot-pow4"
	name_MANDELBROT_POW5 = "mandelbrot-pow5"
	name_MANDELBROT_POW6 = "mandelbrot-pow6"
	name_JULIA_SET       = "julia-set"
	name_PHOENIX         = "phoenix"
	name_BURNING_SHIP    = "burning-ship"
	name_BURNING_SHIP_IM = "burning-ship-im"
	name_SPIDER          = "spider"
	name_TRICORN         = "tricorn"
)

var fmNames = map[Formula]string{
	FM_MANDELBROT:      name_MANDELBROT,
	FM_MANDELBROT_POW3: name_MANDELBROT_POW3,
	FM_MANDELBROT_POW4: name_MANDELBROT_POW4,
	FM_MANDELBROT_POW5: name_MANDELBROT_POW5,
	FM_MANDELBROT_POW6: name_MANDELBROT_POW6,
	FM_JULIA_SET:       name_JULIA_SET,
	FM_PHOENIX:         name_PHOENIX,
	FM_BURNING_SHIP:    name_BURNING_SHIP,
	FM_BURNING_SHIP_IM: name_BURNING_SHIP_IM,
	FM_SPIDER:          name_SPIDER,
	FM_TRICORN:         name_TRICORN,
}

var fmValues = map[string]Formula{
	name_MANDELBROT:      FM_MANDELBROT,
	name_MANDELBROT_POW3: FM_MANDELBROT_POW3,
	name_MANDELBROT_POW4: FM_MANDELBROT_POW4,
	name_MANDELBROT_POW5: FM_MANDELBROT_POW5,
	name_MANDELBROT_POW6: FM_MANDELBROT_POW6,
	name_JULIA_SET:       FM_JULIA_SET,
	name_PHOENIX:         FM_PHOENIX,
	name_BURNING_SHIP:    FM_BURNING_SHIP,
	name_BURNING_SHIP_IM: FM_BURNING_SHIP_IM,
	name_SPIDER:          FM_SPIDER,
	name_TRICORN:         FM_TRICORN,
}

func (f Formula) MarshalJSON() ([]byte, error) {
	value := f
	name, ok := fmNames[value]
	if !ok {
		return nil, fmt.Errorf("gofra.Formula.MarshalJSON: undefined value %d", value)
	}
	return json.Marshal(name)
}

func (f *Formula) UnmarshalJSON(data []byte) error {
	var name string
	err := json.Unmarshal(data, &name)
	if err != nil {
		return err
	}
	value, ok := fmValues[name]
	if !ok {
		return fmt.Errorf("gofra.Formula.UnmarshalJSON: undefined name %q", name)
	}
	*f = value
	return nil
}
