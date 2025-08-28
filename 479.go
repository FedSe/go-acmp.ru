package main

import (
	. "fmt"
	. "math"
)

type T = float64

var (
	a, b, c, d, x, y, r T
	A                   = Acos
	L                   = Min
	G                   = Max
	Q                   = Sqrt
)

func F(x, y, m, l T) T {
	x -= m
	y -= l
	return Q(x*x + y*y)
}

func main() {
	Scan(&a, &b, &c, &d, &x, &y, &r)

	e := F(a, b, c, d)
	s := F(a, b, x, y)
	t := F(c, d, x, y)
	o := A(G(-1, L(1, (-e*e+s*s+t*t)/2/s/t))) - A(L(1, r/s)) - A(L(1, r/t))
	if o > 0 {
		e = Q(G(0, s*s-r*r)) + Q(G(0, t*t-r*r)) + o*r
	}

	Print(e)
}