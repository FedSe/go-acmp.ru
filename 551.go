package main
import (
	. "fmt"
	. "math"
)
func main() {
	var a, c, h, b float64
	Scan(&a, &c, &h, &b)
	s := "YES"

	if h - Sqrt(a*a - c*c) + c > b {
		s = "NO"
	}

	Print(s)
}