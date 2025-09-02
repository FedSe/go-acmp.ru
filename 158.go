package main
import (
	. "fmt"
	. "math/big"
)
func main() {
	var (
		x, k int64
		a    Int
	)
	Scan(&x, &k)
	Print(a.Binomial(x/5+k, x/5))
}