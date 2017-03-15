package main

import (
	"log"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/urfave/cli"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	app := cli.NewApp()
	app.Name = "fractus"
	app.Usage = "task list on the command line"
	app.Version = "0.0.1"
	app.Author = "Chander"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Chander",
			Email: "jpochander@gmail.com",
		},
		cli.Author{
			Name:  "msetracer",
			Email: "msetracer@gmail.com",
		},
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config",
			Value: "fractal.json",
			Usage: "configuration file",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:      "default",
			ShortName: "d",
			Usage:     "reset params",
			Action:    actionDefault,
		},
		{
			Name:      "render",
			ShortName: "r",
			Usage:     "render the fractal",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "image",
					Value: "fractal.png",
					Usage: "destination image file name",
				},
			},
			Action: actionRender,
		},
		{
			Name:      "iter",
			ShortName: "i",
			Usage:     "set iters",
			Action:    actionIter,
		},
		{
			Name:      "scale",
			ShortName: "s",
			Usage:     "scale fractal",
			Action:    actionScale,
		},
		{
			Name:      "move",
			ShortName: "m",
			Usage:     "move center",
			Action:    actionMove,
		},
		{
			Name:      "palette",
			ShortName: "p",
			Usage:     "set palette period and shift",
			Action:    actionPalette,
		},
	}

	app.Run(os.Args)
}

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

	log.Println("work duration:", time.Since(begin))
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

	err = coreMove(configName, factorX, factorY)
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
