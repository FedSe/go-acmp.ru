package main
import (
	. "fmt"
	. "math"
)
func main() {
	var (
		x, y, r [10]float64
		n, c, i int
		v       = 2e2
		t       = v
	)

	Scan(&n, &t)
	for c < n {
		Scan(&x[c], &y[c], &r[c])
		c++
	}
	if n > 1 {
		for i < n-1 {
			c = i + 1
			for c < n {
				w := Sqrt(Pow(x[i]-x[c], 2) + Pow(y[i]-y[c], 2))
				w -= r[i] + r[c]
				if w <= 0 {
					t = 0
					goto A
				}
				if w < v {
					v = w
				}
				c++
			}
			i++
		}
		if v/2 < t {
			t = v / 2
		}
	}
A:
	Print(t)
}