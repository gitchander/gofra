package interval

type Interval struct {
	Min, Max int // [Min, Max)
}

func (i Interval) Width() int {
	return i.Max - i.Min
}

func (a Interval) Split(count int) []Interval {

	if a.Min > a.Max {
		a.Min, a.Max = a.Max, a.Min
	}

	width := a.Max - a.Min

	if count > width {
		count = width
	}

	if count <= 0 {
		return nil
	}

	quo, rem := quoRem(width, count)

	ins := make([]Interval, count)

	var b Interval
	b.Min = a.Min
	for i := 0; i < count; i++ {

		b.Max = b.Min + quo
		if rem > 0 {
			b.Max++
			rem--
		}

		ins[i] = b
		b.Min = b.Max
	}

	return ins
}

func quoRem(a, b int) (quo, rem int) {
	quo = a / b
	rem = a - quo*b
	return
}
