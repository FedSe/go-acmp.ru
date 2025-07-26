package main
import (
	. "fmt"
	. "sort"
)
func main() {
	var (
		a [100]int
		n, s, c, i, j int
	)
	Scan(&n)

	for i < n {
		Scan(&a[i])
	i++
	}
	Ints(a[:n])

	for j < n {
		i = a[j]
		c += i
		if j < n/2 {
			s += i
			c -= i
		}
	j++
	}

	Print(c - s)
}