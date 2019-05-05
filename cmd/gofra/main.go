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
			Action:    actionMove3,
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
