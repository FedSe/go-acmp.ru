package main
import (
	. "fmt"
	. "math/big"
)
func main() {
	var n int64
	N := NewInt
	Scan(&n)
	a := N(n)
	Print(a.Mul(a, N(n-1)).Div(a, N(2)))
}