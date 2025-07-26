package main
import (
   . "fmt"
   . "math/big"
)
func main() {
	n := 0
	r := NewInt(1)

	Scan(&n)
	n = n*(n-1)/2
	for 0 < n {
		r.Mul(r, NewInt(3))
	n--
	}

	Print(r)
}