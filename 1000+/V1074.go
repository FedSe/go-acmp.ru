package main
import (
	. "fmt"
	. "math"
)
func main() {
	var a, b, c, d, e, f, g, h float64
	Scan(&a, &b, &c, &d, &e, &f, &g, &h)
	Print((Hypot((a-e), (b-f)) - c - g) / (d + h))
}