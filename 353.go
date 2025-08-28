package main
import (
	. "fmt"
	. "math"
	. "sort"
)
func main() {
	var (
		d    [1e5]float64
		n, i int
		m    = 0.
	)

	Scan(&n)
	for i < n {
		Scan(&d[i])
		i++
	}

	Float64s(d[:n])
	i = 2
	for i < n {
		a, b, c := d[i-2], d[i-1], d[i]
		if a+b > c {
			s := (a + b + c) / 2
			e := Sqrt(s * (s - a) * (s - b) * (s - c))
			if e > m {
				m = e
			}
		}
		i++
	}

	if m > 0 {
		Printf("%.3f", m)
	} else {
		Print(0)
	}
}