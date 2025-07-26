package main
import (
	. "fmt"
	. "math"
)
func main() {
	var a, b, c, d, e, f float64
	Scan(&a, &b, &c, &d, &e, &f)

	a = Min(a, b+c)
	b = Min(b, a+c)
	w := a+b

	Print(Min(Min(a/d + c/e + b/f, b/d + c/e + a/f), w/d + w/e))
}