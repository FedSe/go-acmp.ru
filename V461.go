package main
import (
	. "fmt"
	. "sort"
)
func main() {
	var (
		s, n, i, j int
		a          [2e3]int
	)

	Scan(&n)
	for i < n {
		Scan(&a[i])
		i++
	}
	Ints(a[:n])

	for j <= n/2 {
		s += a[j] / 2
		j++
	}

	Print(s + j)
}