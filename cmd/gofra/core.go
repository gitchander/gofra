package main

import (
	"errors"
	"fmt"
	"image"
	"path/filepath"

	"github.com/gitchander/gofra"
	"github.com/gitchander/gofra/palgen"
	"github.com/gitchander/gofra/utils/random"
)

func LoadConfigFromFile(fileName string) (*gofra.Config, error) {
	var c gofra.Config
	err := gofra.LoadFromJsonFile(fileName, &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func coreDefault(configName string) error {
	c := gofra.DefaultConfig
	return gofra.SaveToJsonFile(configName, c)
}

func coreRender(configName, imageName string) error {

	config, err := LoadConfigFromFile(configName)
	if err != nil {
		return err
	}

	size := image.Point{
		X: config.ImageSize.Width,
		Y: config.ImageSize.Height,
	}
	m := gofra.NewImageRGBA(size)

	renderWithProgress(m, config)

	ext := filepath.Ext(imageName)
	switch ext {
	case ".png":
		return gofra.SaveImagePNG(imageName, m)
	case ".jpeg":
		return gofra.SaveImageJPEG(imageName, m)
	default:
		return fmt.Errorf("invalid image file ext %q", ext)
	}
}

func coreIter(configName string, n int) error {

	config, err := LoadConfigFromFile(configName)
	if err != nil {
		return err
	}

	if n < 1 {
		return errors.New("iter number < 1")
	}

	config.Calculation.Iterations = n

	return gofra.SaveToJsonFile(configName, config)
}

func coreScale(configName string, scale float64) error {

	config, err := LoadConfigFromFile(configName)
	if err != nil {
		return err
	}

	config.FractalInfo.Location.Radius /= scale

	return gofra.SaveToJsonFile(configName, config)
}

func coreRotate(configName string, angle int) error {

	config, err := LoadConfigFromFile(configName)
	if err != nil {
		return err
	}

	config.RotateDeg(angle)

	return gofra.SaveToJsonFile(configName, config)
}

func coreMove1(configName string, x, y float64) error {

	config, err := LoadConfigFromFile(configName)
	if err != nil {
		return err
	}

	x = cropFloat64(x, -1, +1)
	y = cropFloat64(y, -1, +1)

	config.MoveRelativeLocation(x, y)

	return gofra.SaveToJsonFile(configName, config)
}

func coreMove2(configName string, x, y float64) error {

	config, err := LoadConfigFromFile(configName)
	if err != nil {
		return err
	}

	var (
		dX = config.ImageSize.Width
		dY = config.ImageSize.Height

		minXY = minInt(dX, dY)
	)

	x *= float64(dX) / float64(minXY)
	y *= float64(dY) / float64(minXY)

	config.MoveRelativeLocation(x, y)

	return gofra.SaveToJsonFile(configName, config)
}

func corePalette(configName string, Period, Shift float64) error {

	config, err := LoadConfigFromFile(configName)
	if err != nil {
		return err
	}

	p := &(config.Palette)

	p.Period = Period
	p.Shift = Shift

	return gofra.SaveToJsonFile(configName, config)
}

func coreRandomPalette(configName string) error {

	config, err := LoadConfigFromFile(configName)
	if err != nil {
		return err
	}

	p := &(config.Palette)

	r := random.NewRandNow()
	palgen.RandParams(r, &(p.Params))

	return gofra.SaveToJsonFile(configName, config)
}

func coreAntiAliasing(configName string, aa gofra.AntiAliasing) error {

	config, err := LoadConfigFromFile(configName)
	if err != nil {
		return err
	}

	config.Calculation.AntiAliasing = aa

	return gofra.SaveToJsonFile(configName, config)
}
