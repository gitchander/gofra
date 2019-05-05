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
