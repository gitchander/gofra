package interval

import (
	"math/rand"
	"testing"
	"time"
)

func TestQuoRem(t *testing.T) {

	quoRemSample := func(a, b int) (quo, rem int) {
		quo = a / b
		rem = a % b
		return
	}

	randInt := func(r *rand.Rand) int {
		u := r.Uint32()
		u = u >> uint(r.Intn(31)+1)
		return int(u)
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var quo, rem [2]int

	for i := 0; i < 1000000; i++ {

		a, b := randInt(r), randInt(r)
		if b == 0 {
			b = 1
		}

		quo[0], rem[0] = quoRem(a, b)
		quo[1], rem[1] = quoRemSample(a, b)

		if quo[0] != quo[1] {
			t.Fatalf("quorem(%d, %d): wrong quo", a, b)
		}
		if rem[0] != rem[1] {
			t.Fatalf("quorem(%d, %d): wrong rem", a, b)
		}
	}
}

func TestInterval(t *testing.T) {

	randVal := func(r *rand.Rand) int {
		i := r.Intn(65536)
		if (r.Int() & 1) == 0 {
			i = -i
		}
		return i
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 100000; i++ {

		min, max := randVal(r), randVal(r)
		if min > max {
			min, max = max, min
		}

		count := r.Intn(100)

		a := Interval{min, max}
		as := a.Split(count)

		for j := 0; j < len(as); j++ {

			b := as[j]

			if b.Min >= b.Max {
				t.Fatalf("in[%d]: min(%d) >= max(%d)", j, b.Min, b.Max)
			}
		}

		for j := 0; j < len(as)-1; j++ {

			b0 := as[j]
			b1 := as[j+1]

			if b0.Max != b1.Min {
				t.Fatalf("in[%d].Max(%d) != in[%d].Min(%d)", j, b0.Max, j+1, b1.Min)
			}
		}

		if len(as) > 1 {

			b0 := as[0]
			b1 := as[len(as)-1]

			if a.Min != b0.Min {
				t.Fatalf("i.Min(%d) != in[%d].Min(%d)", a.Min, 0, b0.Min)
			}
			if a.Max != b1.Max {
				t.Fatalf("i.Max(%d) != in[%d].Max(%d)", a.Max, len(as)-1, b1.Max)
			}
		}
	}
}
