package main
import (
	. "fmt"
	. "math"
)
func main() {
	var (
		x, y, d, c [1001]float64
		n, i int
		e = .0
	)

	Scan(&n)

	for i < n {
		Scan(&x[i], &y[i])
		d[i] = Inf(1)
	i++
	}

	d[0] = 0

	for {
		m := Inf(1)
		r := -1

		for
		i = 0
		i < n
		{
			if c[i] < 1 && d[i] < m {
				m = d[i]
				r = i
			}
		i++
		}

		if r == -1 {
			break
		}

		e = Max(e, d[r])
		c[r] = 1

		for
		i = 0
		i < n
		{
			if c[i] < 1 {
				a := x[i] - x[r]
				m = y[i] - y[r]
				d[i] = Min(a*a + m*m, d[i])
			}
		i++
		}
	}

	Printf("%.2f", Sqrt(e))
}