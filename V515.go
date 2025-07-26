package main
import (
	. "fmt"
	. "math"
)

func main() {
	var a, b, z, c, s, n, i float64
	Scan(&n)

	for i < n {
		Scan(&z, &c)
		a -= z
		b -= c
		s += Sqrt(a*a + b*b)
		a = z
		b = c
	i++
	}
	
	s += Sqrt(z*z + c*c)

	Printf("%.3f", s)
}