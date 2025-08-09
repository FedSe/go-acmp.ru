package main
import (
	. "fmt"
	. "math/big"
)
func main() {
	n := 0
	N := NewInt
	Scan(&n)

	a := N(1)
	b := N(2)
	for 0 < n {
		s := N(1).Add(a, b)
		a.Set(b)
		b.Set(s)
		n--
	}

	Print(a)
}