package main
import (
	. "fmt"
	. "math"
	. "sort"
)
func main() {
	var (
		v, x, y []float64
		n, i, j int
		s       = map[float64]int{}
		c       = .0
		k       = c
		P       = Println
	)

	Scan(&n)
	for i < n {
		Scan(&c, &k)
		x = append(x, c)
		y = append(y, k)
		i++
	}

	for j < n {
		i = j + 1
		for i < n {
			c = x[j] - x[i]
			k = y[j] - y[i]
			s[Sqrt(c*c+k*k)] = 1
			i++
		}
		j++
	}

	for d := range s {
		v = append(v, d)
	}
	Float64s(v)

	P(len(v))
	for _, d := range v {
		P(d)
	}
}