package main
import (
	. "fmt"
	. "sort"
)
func main() {
	n := 0
	
	Scan(&n)
	a := make([]int, n)
	for i := range a {
		Scan(&a[i])
	}

	Sort(Reverse(IntSlice(a)))
	for i, v := range a {
		if i < 1 || v != a[i-1] {
			Println(v)
		}
	}
}