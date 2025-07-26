package main
import (
	. "fmt"
	. "math/big"
)
func main() {
	var (
		e Int
		i, n int64
	)
	Scan(&n)
	a := NewInt(1)

	for i < n {
	i++
		e.Add(&e, a.Mul(a, NewInt(i)))
	}

	Print(&e)
}