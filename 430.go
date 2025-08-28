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
	)

	Scan(&r, &a, &c, &b, &d)

	a *= Pi / 180
	b *= Pi / 180
	c *= Pi / 180
	d *= Pi / 180

	x := r*C(a)*C(c) - r*C(b)*C(d)
	y := r*C(a)*S(c) - r*C(b)*S(d)
	z := r*S(a) - r*S(b)

	Print(r * 2 * Asin(Min(1., Sqrt(x*x+y*y+z*z)/2/r)))
}