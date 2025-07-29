package main
import (
	. "fmt"
	. "sort"
)
func main() {
	var (
		n, i, j int
		a       [3e4]int
		P       = Println
	)

	Scan(&n)
	for i < n {
		Scan(&a[i])
		i++
	}
	Ints(a[:n])

	for j < n {
		P(a[j])
		j += 2
	}

	n -= 1 + n%2
	for n > 0 {
		P(a[n])
		n -= 2
	}
}