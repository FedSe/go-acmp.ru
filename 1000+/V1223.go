package main
import (
	. "fmt"
	. "math"
)
func main() {
	var (
		x, y          [1e3]float64
		n, p, r, i, j int
		q             = 1
		s             = 1
		H             = Hypot
	)

	Scan(&n)
	for j < n {
		Scan(&x[j], &y[j])
		j++
	}

	z := H(x[0]-x[1], y[0]-y[1])
	w := z
	n--

	for i < n {
		j = i
		for j < n {
			j++
			c := H(x[i]-x[j], y[i]-y[j])
			if c < z {
				z = c
				p = i
				q = j
			}
			if c > w {
				w = c
				r = i
				s = j
			}
		}
		i++
	}

	Print(p+1, q+1, r+1, s+1)
}