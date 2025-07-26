package main
import (
	. "fmt"
	. "math"
)

func F(x, y, m, l float64) float64 {
	x -= m
	y -= l
	return Sqrt(x*x + y*y)
}

func main() {
	var a, b, c, d, x, y, r float64
	Scan(&a, &b, &c, &d, &x, &y, &r)

	e := F(a, b, c, d)
	s := F(a, b, x, y)
	t := F(c, d, x, y)

	o := Acos(Max(-1, Min(1, (-e*e + s*s + t*t) / 2 / s / t))) - Acos(Min(1, r/s)) - Acos(Min(1, r/t))

	if o > 0 {
		e = Sqrt(Max(0, s*s-r*r)) + Sqrt(Max(0, t*t-r*r)) + o*r
	}

	Print(e)
}