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
	)

	Scan(&n)
	for j < n {
		Scan(&x[j], &y[j])
		j++
	}

	a := x[0] - x[1]
	b := y[0] - y[1]
	z := Sqrt(a*a + b*b)
	w := z
	n--

	for i < n {
		j = i
		for j < n {
			j++
			a = x[i] - x[j]
			b = y[i] - y[j]
			c := Sqrt(a*a + b*b)
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