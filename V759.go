package main
import (
	. "fmt"
	. "sort"
)
func main() {
	var (
		n, m, s, i, j int
		a             [1e3]int
	)

	Scan(&n, &m)
	for i < n {
		Scan(&a[i])
		i++
	}
    
	Ints(a[:n])
	for j < m {
		n--
		if a[n] > 0 {
			s += a[n]
		}
		j++
	}
	Print(s)
}