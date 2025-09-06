package main
import (
	. "fmt"
	. "math"
)
func main() {
	var a, b, z, c, s, n float64

	Scan(&n)
	for 0 < n {
		Scan(&z, &c)
		s += Hypot(a-z, b-c)
		a = z
		b = c
		n--
	}

	Print(s + Hypot(z, c))
}