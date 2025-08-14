package main
import (
	. "fmt"
	. "math"
)

var (
	a, b, c, d, x, y, r float64
	A                   = Acos
	L                   = Min
	G                   = Max
	Q                   = Sqrt
)

func F(x, y, m, l float64) float64 {
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