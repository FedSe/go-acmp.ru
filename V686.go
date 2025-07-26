package main
import (
	. "fmt"
	. "sort"
)
func main() {
	var n, i, j int
	Scan(&n)

	a := make([]int, n)
	for i < n {
		Scan(&a[i])
	i++
	}
	Ints(a)

	for j < n {
		Print(a[j], " ")
	j += 2
	}

	n -= 1 + n%2
	for n > 0 {
		Print(a[n], " ")
	n -= 2
	}
}