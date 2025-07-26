package main
import (
	. "fmt"
	. "math"
)
func main() {
	var x, y, z, c, r, s float64
	p := 2 * Asin(1)
	w := "NO"
	Scan(&x, &y, &z, &c, &r, &s)

	z -= x
	c -= y
	y = Sqrt(z*z + c*c)
	x = 2 * Acos( (y/2) / r )

	if z + c < 1 && p * r * r > s || ( y >= 2*r && 2 * p * r * r > s ) || r*r*(2*p - (x - Sin(x))) > s {
		w = "YES"
	}

	Print(w)
}