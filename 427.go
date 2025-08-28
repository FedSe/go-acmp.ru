package main
import (
	. "fmt"
	. "sort"
)
func main() {
	var (
		u [1e4]int
		i = 0
		n = 0
		a = 1
	)

	Scan(&n)
	for i < n {
		Scan(&u[i])
		i++
	}

	Ints(u[:])
	for _, x := range u {
		if x > a {
			break
		}
		a += x
	}

	Print(a)
}