package main
import (
	. "fmt"
	. "math/big"
)
func main() {
	var a, n Int

	Scan(&a, &n)

	Print(n.Exp(&a, &n, nil))
}