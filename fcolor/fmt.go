package fcolor

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

const (
	formatDecRGB = "%d %d %d"
	formatHexRGB = "#%02x%02x%02x"
)

var (
	regexpDecRGB = regexp.MustCompile(patternDecRGB())
	regexpHexRGB = regexp.MustCompile(patternHexRGB())
)

func reFraming(s string) string {
	return "(" + s + ")"
}

func patternDecRGB() string {

	const whitespace = "\\x20"

	ds := []string{
		"25[0-5]",     // [250 ... 255]
		"2[0-4][0-9]", // [200 ... 249]
		"1[0-9]{2}",   // [100 ... 199]
		"[1-9][0-9]",  // [10 ... 99]
		"[0-9]",       // [0 ... 9]
	}

	for i := range ds {
		ds[i] = reFraming(ds[i])
	}

	digit := reFraming(strings.Join(ds, "|"))

	return "^" + reFraming(digit+whitespace) + "{2}" + digit + "$"
}

func patternHexRGB() string {
	return "^#(([0-9A-Fa-f]){6})$"
}

func rgb_to_str(r, g, b uint8) string {

	var colorFormat string

	//colorFormat= formatDecRGB
	colorFormat = formatHexRGB

	return fmt.Sprintf(colorFormat, r, g, b)
}

func str_to_rgb(s string, r, g, b *uint8) error {

	var colorFormat string

	switch {
	case regexpDecRGB.MatchString(s):
		colorFormat = formatDecRGB

	case regexpHexRGB.MatchString(s):
		colorFormat = formatHexRGB

	default:
		return errors.New(fmt.Sprintf("gofra: Color.UnmarshalJSON: wrong match pattern: \"%s\"", s))
	}

	n, err := fmt.Sscanf(s, colorFormat, r, g, b)
	if err != nil {
		return err
	}
	if n != 3 {
		return errors.New("Color.UnmarshalJSON: wrong parse params")
	}

	return nil
}
