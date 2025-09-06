package main
import (
	. "fmt"
	. "math/big"
)
func main() {
	var a, b Int
	Scan(&a, &b)
	Print(a.Add(&a, &b))
}