package main
import (
   . "fmt"
   . "math/big"
)
func main() {
	var n, i int64
	Scan(&n)
	r := NewInt(1)
	for i < n {
		i++
		r.Mul(r, NewInt(i))
	}
	Print(r)
}