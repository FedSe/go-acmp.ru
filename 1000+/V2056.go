package main
import (
	. "fmt"
	. "strings"
)
func main() {
	var (
		x, y, i int
		k       = "."
		a       = k
		m       = a
		S       = Sscan
		L       = Split
		P       = Print
	)

	Scan(&a, &m)
	for i < 4 {
		S(L(a, k)[i], &x)
		S(L(m, k)[i], &y)
		P(x & y)
		if i < 3 {
			P(k)
		}
		i++
	}
}