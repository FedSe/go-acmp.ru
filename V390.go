package main

import (
	. "fmt"
	. "math"
)

type A = float64

var x, y, a, d, b, e, c, f A

func F(a, d, b, e A) A {
	b -= a
	e -= d
	return Abs(b*(y-d)-(x-a)*e) / Sqrt(b*b+e*e)
}

func main() {
	Scan(&a, &d, &b, &e, &c, &f, &x, &y)
	Print(Min(F(a, d, c, f), Min(F(a, d, b, e), F(b, e, c, f))))
}