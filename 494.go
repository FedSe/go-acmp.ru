package main
import (
	. "fmt"
	. "math"
)
func main() {
	var x, y, z, h, s, r float64
	v := 1e-9

	Scan(&x, &y, &z, &h)
	z -= x
	h -= y
	for r < 1500 {
		r++
		a := (z*z + h*h) * 2
		b := -2 * (x*z + y*h)
		q := Sqrt(b*b - 2*a*(x*x+y*y-r*r))
		k := (b + q) / a
		a = (b - q) / a
		if k > -v && k < 1+v {
			s++
		}
		if a > -v && a < 1+v && k-a > v {
			s++
		}
	}

	Print(s)
}