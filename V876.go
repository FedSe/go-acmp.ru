package main
import (
	. "fmt"
	. "math"
)
func main() {
	var a, b, r, m, t, x, k float64
	Scan(&a, &b, &r)

	for x <= r {
		y := Sqrt(r*r - x*x)
		f := a*x + b*y
		if f > k {
			k = f
			m = x
			t = y
		}
		x += 0.000001
	}

	Println(k, m, t)
}