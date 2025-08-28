package main
import (
	. "fmt"
	. "math"
)
func main() {
	n := 0
	m := 0
	Scan(&m, &n)

	m--
	n--
	if m > n {
		m, n = n, m
	}

	a := 1
	if m == int(Floor(float64(n-m)*(1+Sqrt(5))/2)) {
		a = 2
	}
	Print(a)
}