package main
import (. "fmt" 
. "math"
)
func main() {
	var a, b, c, d, e, f float64
	s:="YES"
	Scan(&c, &d, &a, &e, &f, &b)
	c-=e
	d-=f
	e = Sqrt(c*c + d*d)

	if a > e + b || b > e + a || e > a + b {
		s="NO"
	}

	Print(s)

}