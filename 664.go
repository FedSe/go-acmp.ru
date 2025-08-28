package main
import (
	. "fmt"
	. "math"
)
func main() {
	k := 0.
	m := k
	n := k

	Scan(&k, &m, &n)

	if n > k {
		m *= Ceil((2 * n) / k)/2
	}

	Print(m*2)
}