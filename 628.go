package main

import (
	. "fmt"
	. "math"
)

type A = float64

var (
	p, z [2e4]A
	n, i int
)

func F(g A) A {
	u := 0.
	for I, o := range p[:n] {
		u += Sqrt(Pow(o-g, 2) + Pow(z[I], 2))
	}
	return u
}

func main() {
	Scan(&n)
	for i < n {
		Scan(&p[i], &z[i])
		i++
	}

	l := p[0]
	r := p[1]
	for Abs(r-l) > 1 {
		m := (2*l + r) / 3
		h := (l + 2*r) / 3
		if F(m) < F(h) {
			r = h
			m = l
		}
		l = m
	}

	Print((l + r) / 2)
}