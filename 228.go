package main
import (
	. "fmt"
	. "math"
)
func main() {
	var (
		d, e    [5e3]float64
		n, i, l int
		m       = 100.
	)

	Scan(&n)
	for i < n {
		Scan(&d[i], &e[i])
		i++
	}
	for l < n-1 {
		m = Max(m, Max(m/d[l]*d[l+1], m/e[l]*e[l+1]))
		l++
	}

	Printf("%.2f", m)
}