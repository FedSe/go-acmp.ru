package main
import (
	. "fmt"
	. "math/big"
)
func main() {
	var a, b Int

	Scan(&b)

	Print(b.Div(b.Mul(&b, a.Add(&b, NewInt(1))), NewInt(2)))
}