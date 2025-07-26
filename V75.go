package main
import (
	. "fmt"
	. "math/big"
)
func main() {
	n:=0
	Scan(&n)
	e := NewInt(1)

	for 0 < n {
		e.Mul(e, NewInt(45))
	n--
	}

	Print(e)
}