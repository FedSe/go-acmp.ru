package main
import (
	. "fmt"
	. "sort"
)
func main() {
	var s, n, i, j int
	Scan(&n)

	a := make([]int, n)
	for i < n {
		Scan(&a[i])
	i++
	}
	Ints(a)

	for j <= n/2 {
		s += a[j]/2
	j++
	}

	Print(s+j)
}