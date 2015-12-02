package gofra

var (
	tableSP_4x = []float64{
		-1.0 / 4.0,
		+1.0 / 4.0,
	}

	tableSP_9x = []float64{
		-1.0 / 3.0,
		0.0,
		+1.0 / 3.0,
	}

	tableSP_16x = []float64{
		-3.0 / 8.0,
		-1.0 / 8.0,
		+1.0 / 8.0,
		+3.0 / 8.0,
	}
)

type spPoint struct {
	X float64
	Y float64
}

func makeSpTable(table []float64) []spPoint {

	n := len(table)
	cs := make([]spPoint, n*n)

	for y := 0; y < n; y++ {
		for x := 0; x < n; x++ {
			cs[y*n+x] = spPoint{
				X: table[x],
				Y: table[y],
			}
		}
	}
	return cs
}

func makeSubpixelTable(aa AntiAliasing) (t []spPoint) {

	switch aa {
	case AA_4X:
		t = makeSpTable(tableSP_4x)
	case AA_9X:
		t = makeSpTable(tableSP_9x)
	case AA_16X:
		t = makeSpTable(tableSP_16x)
	}
	return
}
