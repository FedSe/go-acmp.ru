package main
import (
	. "fmt"
	. "math"
)
func main() {
	x := 0.
	y := x
	m := 0
	n := 0
	c := map[float64]int{}

	Scan(&n)

	for 0 < n {
		Scan(&x, &y)
		a := Atan2(y, x)
		c[a]++
		if m < c[a] { m = c[a] }
	n--
	}

	Print(m)
}