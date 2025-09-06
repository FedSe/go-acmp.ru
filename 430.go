package main
import (
	. "fmt"
	. "math"
)
func main() {
	var (
		r, a, b, c, d float64
		C             = Cos
		S             = Sin
		x             = 180.
	)

	Scan(&r, &a, &c, &b, &d)
	a *= Pi / x
	b *= Pi / x
	c *= Pi / x
	d *= Pi / x
	x = r*C(a)*C(c) - r*C(b)*C(d)
	y := r*C(a)*S(c) - r*C(b)*S(d)
	z := r*S(a) - r*S(b)

	Print(r * 2 * Asin(Min(1., Sqrt(x*x+y*y+z*z)/2/r)))
}