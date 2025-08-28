package main
import (
	. "fmt"
	. "math/big"
)
func main() {
	var (
		n, i int64
		N    = NewInt
		r    = N(1)
	)

	Scan(&n)
	for i < n {
		i++
		r.Mul(r, N(i))
	}

	Print(r)
}