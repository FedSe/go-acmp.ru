package main
import (
	. "fmt"
	. "math/big"
)
func main() {
	var n Int

	Scan(&n)

	Print(n.Exp(NewInt(2), &n, nil))
}