package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/urfave/cli"

	"github.com/gitchander/gofra"
)

func actionDefault(c *cli.Context) {

	configName := c.Parent().String("config")

	err := coreDefault(configName)
	checkError(err)

	renderIfNeed(c)
}

func renderIfNeed(c *cli.Context) {
	if c.Parent().Bool("render") {
		actionDraw(c)
	}
}

func actionDraw(c *cli.Context) {

	var (
		configName = c.Parent().String("config")
		imageName  = c.Parent().String("image")
	)

	begin := time.Now()

	err := coreRender(configName, imageName)
	checkError(err)

	fmt.Println("Work duration:", time.Since(begin))
}

func actionMove(c *cli.Context) {
	actionMove4(c)
}

func actionRotate(c *cli.Context) {

	configName := c.Parent().String("config")

	angleDeg, err := strconv.ParseInt(c.Args().First(), 10, 32)
	checkError(err)

	err = coreRotate(configName, int(angleDeg))
	checkError(err)

	renderIfNeed(c)
}

func actionIter(c *cli.Context) {

	configName := c.Parent().String("config")

	n, err := strconv.ParseInt(c.Args().First(), 10, 32)
	checkError(err)

	err = coreIter(configName, int(n))
	checkError(err)

	renderIfNeed(c)
}

func actionScale(c *cli.Context) {

	configName := c.Parent().String("config")
	scaleFactor, err := strconv.ParseFloat(c.Args().First(), 64)
	checkError(err)

	err = coreScale(configName, scaleFactor)
	checkError(err)

	renderIfNeed(c)
}

func actionPalette(c *cli.Context) {

	configName := c.Parent().String("config")

	args := c.Args()

	palPeriod, err := strconv.ParseFloat(args.Get(0), 64)
	checkError(err)

	palShift, err := strconv.ParseFloat(args.Get(1), 64)
	checkError(err)

	err = corePalette(configName, palPeriod, palShift)
	checkError(err)

	renderIfNeed(c)
}

func actionRandomPalette(c *cli.Context) {

	configName := c.Parent().String("config")

	err := coreRandomPalette(configName)
	checkError(err)

	renderIfNeed(c)
}

func actionAntiAliasing(c *cli.Context) {

	configName := c.Parent().String("config")

	aa, err := gofra.ParseAntiAliasing(c.Args().First())
	checkError(err)

	err = coreAntiAliasing(configName, aa)
	checkError(err)

	renderIfNeed(c)
}
