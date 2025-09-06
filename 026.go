package main
import (
	. "fmt"
	. "math"
)
func main() {
	var a, b, c, d, e, f float64
	s := "YES"

	Scan(&c, &d, &a, &e, &f, &b)
	c = Hypot(c-e, d-f)
	if a > c+b || b > c+a || c > a+b {
		s = "NO"
	}

	Print(s)
}