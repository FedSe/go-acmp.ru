package main
import (
	. "fmt"
	. "math"
)
func main() {
	n := 0
	i := 0

	Scan(&n)
	p := make([][2]float64, n)
	for i < n {
		Scan(&p[i][0], &p[i][1])
	i++
	}

	F := func (g float64) float64 {
		u := 0.
		for _, o := range p {
			u += Sqrt(Pow(o[0]-g, 2) + Pow(o[1], 2))
		}
		return u
	}

	l := p[0][0]
	r := p[1][0]
	for Abs(r-l) > 1 {
		m := (2*l + r) / 3
		h := (l + 2*r) / 3

		if F(m) < F(h) {
			r = h
		} else {
			l = m
		}
	}

	Print((l+r)/2)
}