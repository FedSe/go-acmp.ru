package main
import (
	. "fmt"
	. "math"
)
func main() {
	n := 0.
	Scan(&n)
	Printf("%.f", 2*n+Ceil(2*Sqrt(n)))
}