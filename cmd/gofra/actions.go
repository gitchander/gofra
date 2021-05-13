package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/urfave/cli/v2"

	"github.com/gitchander/gofra"
)

func actionDefault(c *cli.Context) error {

	configName := c.String("config")

	err := coreDefault(configName)
	if err != nil {
		return err
	}

	renderIfNeed(c)

	return nil
}

func renderIfNeed(c *cli.Context) {
	if c.Bool("render") {
		actionDraw(c)
	}
}

func actionDraw(c *cli.Context) error {

	var (
		configName = c.String("config")
		imageName  = c.String("image")
	)

	begin := time.Now()

	err := coreRender(configName, imageName)
	if err != nil {
		return err
	}

	fmt.Println("Work duration:", time.Since(begin))

	return nil
}

func actionMove(c *cli.Context) error {
	return actionMove4(c)
}

func actionRotate(c *cli.Context) error {

	configName := c.String("config")

	angleDeg, err := strconv.ParseInt(c.Args().First(), 10, 32)
	if err != nil {
		return err
	}

	err = coreRotate(configName, int(angleDeg))
	if err != nil {
		return err
	}

	renderIfNeed(c)

	return nil
}

func actionIter(c *cli.Context) error {

	configName := c.String("config")

	n, err := strconv.ParseInt(c.Args().First(), 10, 32)
	if err != nil {
		return err
	}

	err = coreIter(configName, int(n))
	if err != nil {
		return err
	}

	renderIfNeed(c)

	return nil
}

func actionScale(c *cli.Context) error {

	configName := c.String("config")
	scaleFactor, err := strconv.ParseFloat(c.Args().First(), 64)
	if err != nil {
		return err
	}

	err = coreScale(configName, scaleFactor)
	if err != nil {
		return err
	}

	renderIfNeed(c)

	return nil
}

func actionPalette(c *cli.Context) error {

	configName := c.String("config")

	args := c.Args()

	palPeriod, err := strconv.ParseFloat(args.Get(0), 64)
	if err != nil {
		return err
	}

	palShift, err := strconv.ParseFloat(args.Get(1), 64)
	if err != nil {
		return err
	}

	err = corePalette(configName, palPeriod, palShift)
	if err != nil {
		return err
	}

	renderIfNeed(c)

	return nil
}

func actionRandomPalette(c *cli.Context) error {

	configName := c.String("config")

	err := coreRandomPalette(configName)
	if err != nil {
		return err
	}

	renderIfNeed(c)

	return nil
}

func actionAntiAliasing(c *cli.Context) error {

	configName := c.String("config")

	aa, err := gofra.ParseAntiAliasing(c.Args().First())
	if err != nil {
		return err
	}

	err = coreAntiAliasing(configName, aa)
	if err != nil {
		return err
	}

	renderIfNeed(c)

	return nil
}
