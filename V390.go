package main
import (
	. "fmt"
	. "math"
)
var x, y, a, d, b, e, c, f float64

func F(a, d, b, e float64) float64 {
	b-=a
	e-=d
	return Abs(b*(y-d) - (x-a)*e) / Sqrt(b*b + e*e)
}

func main() {
	Scan(&a, &d, &b, &e, &c, &f, &x, &y)
	Print(Min(F(a, d, c, f), Min(F(a, d, b, e), F(b, e, c, f))))
}