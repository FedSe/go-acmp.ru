package main
import (
	. "fmt"
	. "math/big"
)
var (
	n, a, b, j int
	N          = NewInt
	q          = N(1)
)
func main() {
	Scan(&n, &a, &b)

	for j < 2 {
		a += n
		k := n
		e := 1
		i := 0
		for i < k {
			e *= a - i
			i++
			e /= i
		}
		j++
		a = b
		q.Mul(q, N(int64(e)))
	}

	Print(q)
}