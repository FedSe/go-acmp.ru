package main
import (
	. "fmt"
	. "sort"
)
func main() {
	var (
		a          [100]int
		n, s, i, k int
	)

	Scan(&n, &s)
	for i < n {
		Scan(&a[i])
		i++
	}
	Ints(a[:n])
	for k < n && s >= a[k] {
		s -= a[k]
		k++
	}

	Print(k)
}