package main
import (
	. "fmt"
	. "math"
)
func main() {
	var (
		m = .396850263
		z = m*m - Sqrt(m)
		l = 0.
		c = l
		r = m
		Q = 1e-9
		S = "%.6f "
		P = Printf
	)

	F := func() {
		m := 0.
		t := 0
		for t < 200 {
			t++
			m = (l + r) / 2
			if m*m-Sqrt(m) > c {
				l = m
			} else {
				r = m
			}
		}
	}

	Scan(&c)
	if c < Q && c > z+Q {
		F()
		P(S, r)
	}
	if c > z-Q {
		l = 1e6
		r = m
		F()
		P(S, r)
	}
	if c < z-Q {
		P("No solution")
	}
}