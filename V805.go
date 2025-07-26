package main
import . "fmt"
func main() {
	var (
		n, x, y, l, i int
		a = -100000
		b = a
		c = a*-3
	)

	Scan(&n)
	for i < n {
		Scan(&x, &y, &l)
		if a < x { a = x }
		if b < y { b = y }
		x += y+l
		if c > x { c = x }
	i++
	}
	c -= a + b
	s := float64(c*c) / 2.
	if c < 0 {
		s = 0
	}

	Print(s)
}