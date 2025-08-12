package main
import (
	. "fmt"
	. "math/big"
)
func main() {
	var (
		e    Int
		i, n int64
		N    = NewInt
		a    = N(1)
	)

	Scan(&n)
	for i < n {
		i++
		e.Add(&e, a.Mul(a, N(i)))
	}

	Print(&e)
}