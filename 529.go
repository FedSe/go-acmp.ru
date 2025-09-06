package main
import (
	. "fmt"
	. "math"
)
func main() {
	var a, b, c, d float64
	Scan(&a, &b, &c, &d)
	Print(Hypot(c-a, d-b))
}