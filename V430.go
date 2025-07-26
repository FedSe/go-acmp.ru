package main
import (
	. "fmt"
	. "math"
)
func main() {
	var r, a, b, c, d float64
	Scan(&r, &a, &c, &b, &d)

	a = a / 180 * Pi
	b = b / 180 * Pi
	c = c / 180 * Pi
	d = d / 180 * Pi

	x := r * Cos(a) * Cos(c) - r * Cos(b) * Cos(d)
	y := r * Cos(a) * Sin(c) - r * Cos(b) * Sin(d)
	z := r * Sin(a) - r * Sin(b)

	Print(r * 2 * Asin(Min(1., Sqrt(x*x + y*y + z*z)/2/r)))
}