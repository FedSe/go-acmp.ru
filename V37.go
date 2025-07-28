package main
import (
	. "fmt"
	. "math"
)
func main() {
	var (
		n, x, y, z, k int
		q             = .0
		s             = "Yes"
	)

	Scan(&n, &q)
	t := int(Round(q * 1000))

	for 0 < n {
		Scan(&x, &y, &z, &k)
		if (z*z+k*k)*1000000 > t*t*(x*x+y*y) {
			s = "No"
		}
		n--
	}

	Print(s)
}