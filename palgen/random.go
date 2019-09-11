package palgen

import (
	"math/rand"
	"time"
)

func NewRandSeed(seed int64) *rand.Rand {
	return rand.New(rand.NewSource(seed))
}

func NewRandTime(t time.Time) *rand.Rand {
	return NewRandSeed(t.UTC().UnixNano())
}

func NewRandNow() *rand.Rand {
	return NewRandTime(time.Now())
}

func RandParams(r *rand.Rand, p *Params) {

	p.A = Vec3{0.5, 0.5, 0.5}
	p.B = Vec3{0.5, 0.5, 0.5}

	if true {
		const nC = 5
		p.C = Vec3{
			0: float64(r.Intn(nC)),
			1: float64(r.Intn(nC)),
			2: float64(r.Intn(nC)),
		}
	} else {
		const nC = 5
		p.C = Vec3{
			0: r.Float64() * nC,
			1: r.Float64() * nC,
			2: r.Float64() * nC,
		}
	}

	// random value in range [0..1]
	p.D = Vec3{
		0: r.Float64(),
		1: r.Float64(),
		2: r.Float64(),
	}
}
