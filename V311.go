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
	)
	
	Scan(&n)
	a := N(1)

	for i < n {
		i++
		e.Add(&e, a.Mul(a, N(i)))
	}

	Print(&e)
}