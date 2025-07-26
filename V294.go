package main
import . "fmt"
func main() {
	var a, b, c, d, e, f int
	Scan(&a, &c, &e, &b, &d, &f)

	s := a*e + b*f
	a -= a*c/100
	b -= b*d/100

	c = a
	if a > b {
		a = b
		c = b
	}

	Print(s - a*e - c*f)
}