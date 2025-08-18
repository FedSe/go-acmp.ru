package main
import (
	. "fmt"
	. "math/big"
)
func main() {
	var N Int
	Scan(&N)
	Print(N.Div(N.Mul(&N, &N), NewInt(4)))
}