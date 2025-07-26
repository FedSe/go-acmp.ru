package main
import (
	. "fmt"
	. "math"
)
func main() {
	var ( x, y, l float64
	n, v int
)
	Scan(&n)
	for 0 < n {
		Scan(&v, &l)
		if v == 1 {
			y += l
		}
		if v == 3 {
			x += l
		}
		if v == 5 {
			y -= l
		}
		if v == 7 {
			x -= l
		}

		l /= Sqrt(2)
		if v == 2 {
			x += l
			y += l
		}
		if v == 4 {
			x += l
			y -= l
		}
		if v == 6 {
			x -= l
			y -= l
		}
		if v == 8 {
			x -= l
			y += l
		}
	n--
	}

	Printf("%.3f %.3f", x, y)
}