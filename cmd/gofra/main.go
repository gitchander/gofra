package main

import (
	"log"
	"os"
	"runtime"

	"github.com/urfave/cli/v2"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	app := cli.NewApp()

	app.Name = "gofra"
	app.Usage = "task list on the command line"
	app.Version = "0.0.1"

	app.Authors = []*cli.Author{
		&cli.Author{
			Name:  "chander",
			Email: "jpochander@gmail.com",
		},
		&cli.Author{
			Name:  "msetracer",
			Email: "msetracer@gmail.com",
		},
	}

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "config",
			Value: "fractal.json",
			Usage: "source configuration file name",
		},
		&cli.StringFlag{
			Name:  "image",
			Value: "fractal.jpeg",
			Usage: "destination image file name",
		},
		&cli.BoolFlag{
			Name:  "render",
			Usage: "render image after changed",
		},
	}

	app.Commands = []*cli.Command{
		{
			Name:    "default",
			Aliases: []string{"d"},
			Usage:   "reset config",
			Action:  actionDefault,
		},
		{
			Name:    "draw",
			Aliases: []string{"draw"},
			Usage:   "draw the fractal",
			Action:  actionDraw,
		},
		{
			Name:    "iter",
			Aliases: []string{"i"},
			Usage:   "set iters",
			Action:  actionIter,
		},
		{
			Name:    "scale",
			Aliases: []string{"s"},
			Usage:   "scale fractal",
			Action:  actionScale,
		},
		{
			Name:    "move",
			Aliases: []string{"m"},
			Usage:   "move center",
			Action:  actionMove,
		},
		{
			Name:    "rotate",
			Aliases: []string{"rotate"},
			Usage:   "rotate about center",
			Action:  actionRotate,
		},
		{
			Name:    "palette",
			Aliases: []string{"p"},
			Usage:   "set palette period and shift",
			Action:  actionPalette,
		},
		{
			Name:    "random_palette",
			Aliases: []string{"randpal"},
			Usage:   "random palette",
			Action:  actionRandomPalette,
		},
		{
			Name:    "anti_aliasing",
			Aliases: []string{"antial"},
			Usage:   "anti-aliasing",
			Action:  actionAntiAliasing,
		},
	}

	err := app.Run(os.Args)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
