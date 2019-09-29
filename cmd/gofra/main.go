package main

import (
	"os"
	"runtime"

	"github.com/urfave/cli"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	app := cli.NewApp()
	app.Name = "gofra"
	app.Usage = "task list on the command line"
	app.Version = "0.0.1"
	app.Author = "Chander"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "chander",
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
			Usage: "source configuration file name",
		},
		cli.StringFlag{
			Name:  "image",
			Value: "fractal.png",
			Usage: "destination image file name",
		},
		cli.BoolFlag{
			Name:  "render",
			Usage: "render image after changed",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:      "default",
			ShortName: "d",
			Usage:     "reset config",
			Action:    actionDefault,
		},
		{
			Name:      "draw",
			ShortName: "draw",
			Usage:     "draw the fractal",
			Action:    actionDraw,
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
			Name:      "rotate",
			ShortName: "rotate",
			Usage:     "rotate about center",
			Action:    actionRotate,
		},
		{
			Name:      "palette",
			ShortName: "p",
			Usage:     "set palette period and shift",
			Action:    actionPalette,
		},
		{
			Name:      "random_palette",
			ShortName: "randpal",
			Usage:     "random palette",
			Action:    actionRandomPalette,
		},
		{
			Name:      "anti_aliasing",
			ShortName: "antial",
			Usage:     "anti-aliasing",
			Action:    actionAntiAliasing,
		},
	}

	app.Run(os.Args)
}
