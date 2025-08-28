package main
import (
	. "fmt"
	. "math/big"
)
func main() {
	var (
		k, n, i int
		d       [301]Int
	)
	d[0] = *NewInt(1)

	Scan(&k, &n)
	for i < n {
		i++
		j := 1
		for j <= k && j <= i {
			d[i].Add(&d[i], &d[i-j])
			j++
		}
	}

	Print(&d[n])
}