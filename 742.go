package main
import (
	. "fmt"
	. "math/big"
)
func main() {
	n := 0
	N := NewInt
	r := N(1)

	Scan(&n)
	n = n * (n - 1) / 2
	for 0 < n {
		r.Mul(r, N(3))
		n--
	}

	Print(r)
}