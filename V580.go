package main
import (
	. "fmt"
	. "math"
)
func main() {
	var a, b, c, r float64
	s := "YES"

	Scan(&a, &b, &c, &r)

	p := (a + b + c) / 2
	if r > Sqrt((p-a)*(p-b)*(p-c)/p) {
		s = "NO"
	}

	Print(s)
}