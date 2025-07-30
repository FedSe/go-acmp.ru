package main
import (
	. "fmt"
	. "math/big"
)
func main() {
	var (
		k, n, i int
		d       [301]*Int
		N       = NewInt
	)
	d[0] = N(1)

	Scan(&k, &n)
	for i < n {
		i++
		d[i] = N(0)
		j := 1
		for j <= k && j <= i {
			d[i].Add(d[i], d[i-j])
			j++
		}
	}

	Print(d[n])
}