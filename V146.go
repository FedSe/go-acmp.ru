package main
import (
	. "fmt"
	. "math/big"
)
func main() {
	var a Int

	Scan(&a)

	Print(a.Sqrt(&a))
}