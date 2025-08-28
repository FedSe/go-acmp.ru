package main
import (
	. "fmt"
	. "math/big"
)
func main() {
	var a, b Int
	N := NewInt
	Scan(&b)

	Print(b.Div(b.Mul(&b, a.Add(&b, N(1))), N(2)))
}