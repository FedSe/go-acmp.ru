package main
import . "fmt"
var a, b, c, d, e, f, g, h int

func F(a, b, c, d, e, f int) bool {
	c -= a
	d -= b
	return (c*(f-b)-d*(e-a))*(c*(h-b)-d*(g-a)) >= 0
}

func main() {
	s := "Out"

	Scan(&a, &b, &c, &d, &e, &f, &g, &h)

	if F(a, b, c, d, e, f) && F(c, d, e, f, a, b) && F(e, f, a, b, c, d) {
		s = "In"
	}

	Print(s)
}