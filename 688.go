package main
import (
	. "fmt"
	. "math"
)
func main() {
	var s, S, d, D, x, y, n, i float64

	Scan(&s, &S, &d, &D, &n)
	for i < n {
		i++
		Scan(&x, &y)
		a := 2 * Hypot(s-x, S-y)
		x = d - x
		y = D - y
		if a <= Hypot(x, y) {
			Print(i)
			return
		}
	}

	Print("NO")
}