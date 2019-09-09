package main

import (
	"errors"
	"image"

	"github.com/gitchander/gofra"
)

func newParamsFromFile(fileName string) (*gofra.Parameters, error) {
	var p gofra.Parameters
	if err := p.LoadFromFile(fileName); err != nil {
		return nil, err
	}
	return &p, nil
}

func coreDefault(configName string) error {
	params := gofra.DefaultParameters
	return params.SaveToFile(configName)
}

func coreRender(configName, imageName string) error {

	params, err := newParamsFromFile(configName)
	if err != nil {
		return err
	}

	size := image.Point{
		X: params.ImageSize.Width,
		Y: params.ImageSize.Height,
	}
	m := gofra.NewImageRGBA(size)

	renderWithProgress(m, params)

	if err := gofra.ImageSaveToPNG(m, imageName); err != nil {
		return err
	}

	return nil
}

func coreIter(configName string, n int) error {

	params, err := newParamsFromFile(configName)
	if err != nil {
		return err
	}

	if n < 1 {
		return errors.New("iter number < 1")
	}

	params.Calculation.Iterations = n

	return params.SaveToFile(configName)
}

func coreScale(configName string, scale float64) error {

	params, err := newParamsFromFile(configName)
	if err != nil {
		return err
	}

	params.FractalInfo.Location.Radius /= scale

	return params.SaveToFile(configName)
}

func coreRotate(configName string, angle int) error {

	params, err := newParamsFromFile(configName)
	if err != nil {
		return err
	}

	anglePrev := params.FractalInfo.Location.AngleDeg
	params.FractalInfo.Location.AngleDeg = angleDegNorm(anglePrev + angle)

	return params.SaveToFile(configName)
}

func coreMove1(configName string, x, y float64) error {

	params, err := newParamsFromFile(configName)
	if err != nil {
		return err
	}

	x = cropFloat64(x, -1, +1)
	y = cropFloat64(y, -1, +1)

	params.MoveRelativeLocation(x, y)

	return params.SaveToFile(configName)
}

func coreMove2(configName string, x, y float64) error {

	params, err := newParamsFromFile(configName)
	if err != nil {
		return err
	}

	var (
		dX = params.ImageSize.Width
		dY = params.ImageSize.Height

		minXY = minInt(dX, dY)
	)

	x *= float64(dX) / float64(minXY)
	y *= float64(dY) / float64(minXY)

	params.MoveRelativeLocation(x, y)

	return params.SaveToFile(configName)
}

func corePalette(configName string, Period, Shift float64) error {

	params, err := newParamsFromFile(configName)
	if err != nil {
		return err
	}

	p := &(params.Palette)

	p.Period = Period
	p.Shift = Shift

	return params.SaveToFile(configName)
}
