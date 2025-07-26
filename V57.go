package main
import (
	. "fmt"
	. "math"
)
func main() {
	var (
		n, i, f int
		c = 0.
		p = c
		e = c
		s = "NO"
		a [1001][2]float64
	)
	Scan(&n, &c, &p)

	for f <= n {
		Scan(&a[f][0], &a[f][1])
	f++
	}

	for i < n {
		f = 0
		t := 0.
		for f <= n {
			if f != i {
				x := a[i][0]-a[f][0]
				y := a[i][1]-a[f][1]
				t += Sqrt(x*x + y*y)
			}
		f++
		}

		if t < e || i < 1 {
			e = t
		}
	i++
	}

	if e*c <= p {
		s = "YES"
	}

	Print(s)
}