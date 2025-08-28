package main
import (
	. "fmt"
	. "math"
	. "sort"
)
func main() {
	var (
		e             []float64
		N, r, g       int
		R, X, Y, x, y float64
	)

	Scan(&X, &Y, &R, &N)
	for 0 < N {
		Scan(&x, &y)
		x -= X
		y -= Y
		if x*x+y*y <= R*R+1e-9 {
			e = append(e, Atan2(y, x))
		}
		N--
	}

	if len(e) > 0 {
		Float64s(e)
		for _, a := range e {
			e = append(e, a+2*Pi)
		}
		for r < len(e) {
			for e[r]-e[N] > Pi+1e-9 {
				N++
			}
			w := r - N + 1
			if w > g {
				g = w
			}
			r++
		}
	}
	Print(g)
}