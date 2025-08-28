package main
import (
	. "fmt"
	. "math"
)
func main() {
	var (
		x, y, l float64
		n, v    int
	)
	Scan(&n)
	for 0 < n {
		Scan(&v, &l)
		q := l / Sqrt(2)
		h := []float64{0, q, l, q, 0, -q, -l, -q, 0, q}
		x += h[v-1]
		y += h[v+1]
		n--
	}

	Printf("%.3f %.3f", x, y)
}