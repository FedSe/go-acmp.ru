package main
import (
	. "fmt"
	. "math/big"
)
func main() {
	n := 0
	N := NewInt
	e := N(1)

	Scan(&n)
	for 0 < n {
		e.Mul(e, N(45))
		n--
	}
	Print(e)
}