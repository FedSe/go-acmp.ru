package main
import . "fmt"
func main() {
	var a, b, c, d, e, f, r, w int
	s := "NO"

	Scan(&a, &e, &b, &f, &c, &r, &d, &w)

	d -= c
	w -= r
	x := b - a
	y := f - e

	if d*x != -w*y || b*b+f*f-a*a-e*e == 2*(c*x+r*y) {
		s = "YES"
	}

	Print(s)
}