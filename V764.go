package main
import (
	. "fmt"
	. "math"
)
func main() {
	var (
		x    = 0.
		y    = x
		m, n int
		c    = map[any]int{}
	)

	Scan(&n)
	for 0 < n {
		Scan(&x, &y)
		a := Atan2(y, x)
		c[a]++
		if m < c[a] {
			m = c[a]
		}
		n--
	}

	Print(m)
}