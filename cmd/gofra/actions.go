package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/urfave/cli"
)

func actionDefault(c *cli.Context) {

	configName := c.Parent().String("config")

	err := coreDefault(configName)
	if err != nil {
		log.Fatal(err)
	}
}

func actionRender(c *cli.Context) {

	configName := c.Parent().String("config")
	imageName := c.String("image")

	begin := time.Now()

	err := coreRender(configName, imageName)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Work duration:", time.Since(begin))
}

func actionIter(c *cli.Context) {

	configName := c.Parent().String("config")

	n, err := strconv.ParseInt(c.Args().First(), 10, 32)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = coreIter(configName, int(n))
	if err != nil {
		log.Fatal(err.Error())
	}
}

func actionScale(c *cli.Context) {

	configName := c.Parent().String("config")
	scaleFactor, err := strconv.ParseFloat(c.Args().First(), 64)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = coreScale(configName, scaleFactor)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func actionMove(c *cli.Context) {

	// ./mset move 5 5

	configName := c.Parent().String("config")

	args := c.Args()

	//	factorX, err := strconv.ParseInt(args[0], 10, 64)
	//	if err != nil {
	//		log.Fatal(err.Error())
	//	}

	//	factorY, err := strconv.ParseInt(args[1], 10, 64)
	//	if err != nil {
	//		log.Fatal(err.Error())
	//	}

	const w = 10

	factorX, err := parseMoveCoord(args[0], w/2)
	if err != nil {
		log.Fatal(err.Error())
	}

	factorY, err := parseMoveCoord(args[1], w/2)
	if err != nil {
		log.Fatal(err.Error())
	}

	var (
		x = cropRange(int(factorX), 0, w)
		y = cropRange(int(factorY), 0, w)
	)

	/*

		+------------+------------+
		|(0,0)       |(5,0)       |(10,0)
		|            |            |
		|            |            |
		+------------+------------+
		|(0,5)       |(5,5)       |(10,5)
		|            |            |
		|            |            |
		+------------+------------+
		 (0,10)       (5,10)       (10,10)

	*/

	var (
		xf = float64(2*x-w) / w
		yf = float64(2*(w-y)-w) / w
	)

	err = coreMove(configName, xf, yf)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func parseMoveCoord(s string, defaultVal int) (int, error) {
	if s == "-" {
		return defaultVal, nil
	}
	val, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return int(val), nil
}

func actionMovePrev(c *cli.Context) {

	// ./mset move -- -0.5 0.1

	configName := c.Parent().String("config")

	args := c.Args()

	factorX, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		log.Fatal(err.Error())
	}

	factorY, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = coreMovePrev(configName, factorX, factorY)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func actionPalette(c *cli.Context) {

	configName := c.Parent().String("config")

	args := c.Args()

	palPeriod, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		log.Fatal(err.Error())
	}

	palShift, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = corePalette(configName, palPeriod, palShift)
	if err != nil {
		log.Fatal(err.Error())
	}
}
