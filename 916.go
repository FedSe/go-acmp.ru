package main
import (
	. "fmt"
	. "sort"
)
func main() {
	var (
		a                [100]int
		n, k, p, i, j, s int
	)

	Scan(&n, &k)
	for j < n {
		Scan(&a[j])
		j++
	}

	Ints(a[:n])
	for i < n {
		if i%k < 1 {
			p++
		}
		i++
		s += a[n-i] * p
	}

	Print(s)
}