package main
import (
	. "fmt"
	. "math"
)
func main() {
	n := 0
	s := "NO"

	Scan(&n)
	a := int(Sqrt(float64(n)))
	q := a + 1
	if a*a == n || q*q == n {
		s = "YES"
	}

	Print(s)
}