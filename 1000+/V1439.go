package main
import (
	. "fmt"
	. "sort"
)
func main() {
	n := 0
	i := 0

	Scan(&n)
	a := make([]int, n)
	for i < n {
		Scan(&a[i])
		i++
	}
	Ints(a)

	Print(a[n-1] * a[1])
}