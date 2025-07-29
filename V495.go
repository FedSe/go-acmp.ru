package main
import (
	. "fmt"
	. "math"
)
func main() {
	var (
		x, y, w, v [100]float64
		n, m, z, j int
		s          = 0.
		S = Scan
	)

	S(&n)
	for z < n {
		S(&x[z], &y[z])
		z++
	}

	S(&m)
	for 0 < m {
		z = 0
		for z < n {
			k := 0
			if z < n-1 {
				k = z + 1
			}
			w[z] = (x[z] + x[k]) / 2
			v[z] = (y[z] + y[k]) / 2

			z++
		}
		z = 0
		for z < n {
			x[z] = w[z]
			y[z] = v[z]
			z++
		}
		m--
	}

	for j < n {
		z = 0
		if j < n-1 {
			z = j + 1
		}
		s += Sqrt((x[j]-x[z])*(x[j]-x[z]) + (y[j]-y[z])*(y[j]-y[z]))
		j++
	}

	Print(s)
}