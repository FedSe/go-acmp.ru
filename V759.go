package main
import (
	. "fmt"
	. "sort"
)
func main() {
	var n, m, s, i, j int

	Scan(&n, &m)
	a := make([]int, n)

	for i < n {
		Scan(&a[i])
	i++
	}
	Ints(a)

	for j < m {
		n--
		if a[n] > 0 {
			s += a[n]
		}
	j++
	}
	Print(s)
}