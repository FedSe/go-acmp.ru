package main
import (
	. "fmt"
	. "math"
)
func main() {
	n := .0
	a := n
	s := "NO"

	Scan(&n, &a)
	a /= 2
	if a/Sin(Pi/n) < a/Tan(Pi/n)+1 {
		s = "YES"
	}

	Print(s)
}