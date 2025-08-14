package main
import . "fmt"
func main() {
	var a, b, c, d, e, f, g, h int
	s := "Out"
    
	Scan(&a, &b, &c, &d, &e, &f, &g, &h)
    
	F := func(a, b, c, d, e, f int) bool {
		c -= a
		d -= b
		return (c*(f-b)-d*(e-a))*(c*(h-b)-d*(g-a)) >= 0
	}

	if F(a, b, c, d, e, f) && F(c, d, e, f, a, b) && F(e, f, a, b, c, d) {
		s = "In"
	}

	Print(s)
}