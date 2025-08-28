package main
import (
	. "fmt"
	. "sort"
)
func main() {
	var a [6]int
	i := 0

	for i < 6 {
		Scan(&a[i])
		i++
	}

	Ints(a[:3])
	Ints(a[3:])

	Print(a[0]*a[3] + a[1]*a[4] + a[2]*a[5])
}