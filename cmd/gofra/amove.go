package main

import (
	"errors"
	"log"
	"strconv"

	"github.com/urfave/cli"
)

func actionMove1(c *cli.Context) {

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

	err = coreMove1(configName, factorX, factorY)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func actionMove2(c *cli.Context) {

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
		x = cropInt(int(factorX), 0, w)
		y = cropInt(int(factorY), 0, w)
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

	err = coreMove2(configName, xf, yf)
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

// ./mset move ..
// ./mset move N_

// Cardinal direction
// https://en.wikipedia.org/wiki/Cardinal_direction
// directions:
// n - north
// e - east
// s - south
// w - west

/*

	(NW)---(Nw)---(N.)---(Ne)---(NE)
	 ||     ||     ||     ||     ||
	 ||     ||     ||     ||     ||
	(nW)---(nw)---(n.)---(ne)---(nE)
	 ||     ||     ||     ||     ||
	 ||     ||     ||     ||     ||
	(.W)---(.w)---(..)---(.e)---(.E)
	 ||     ||     ||     ||     ||
	 ||     ||     ||     ||     ||
	(sW)---(sw)---(s.)---(se)---(sE)
	 ||     ||     ||     ||     ||
	 ||     ||     ||     ||     ||
	(SW)---(Sw)---(S.)---(Se)---(SE)

*/

func actionMove3(c *cli.Context) {

	// ./mset move ph mh

	configName := c.Parent().String("config")

	args := c.Args()

	if len(args) < 2 {
		log.Fatal("move action need two parameters")
	}

	factorX, err := parseCoef(args[0])
	if err != nil {
		log.Fatal(err.Error())
	}

	factorY, err := parseCoef(args[1])
	if err != nil {
		log.Fatal(err.Error())
	}

	err = coreMove1(configName, factorX, factorY)
	if err != nil {
		log.Fatal(err.Error())
	}
}

// w (whole - целый)
// h (half - половина)
// q (quarter - четверть)
// e (eighth - восьмая часть)

var errInsuffDataLen = errors.New("insufficient data length")

func parseCoef(s string) (float64, error) {
	bs := []byte(s)

	if len(bs) == 0 {
		return 0, errInsuffDataLen
	}

	var negative bool
	switch b := bs[0]; b {
	case 'm':
		negative = true
		fallthrough
	case 'p':
		bs = bs[1:]
	}

	if len(bs) == 0 {
		return 0, errInsuffDataLen
	}

	var v float64
	for _, b := range bs {
		switch b {
		case 'w':
			v += 1
		case 'h':
			v += 0.5 // 1 / 2
		case 'q':
			v += 0.25 // 1 / 4
		case 'e':
			v += 0.125 // 1 / 8
		case 'z':
			v += 0
		default:
			return 0, errors.New("invalid coefficient")
		}
	}

	if negative {
		v = -v
	}

	return v, nil
}
