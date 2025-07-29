package main
import (
	. "fmt"
	. "math"
)
func main() {
	var a, b, c, d, e, f float64
	M := Min
	Scan(&a, &b, &c, &d, &e, &f)

	a = M(a, b+c)
	b = M(b, a+c)
	w := a + b

	Print(M(M(a/d+c/e+b/f, b/d+c/e+a/f), w/d+w/e))
}