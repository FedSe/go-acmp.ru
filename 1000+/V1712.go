package main
import (
	. "fmt"
	. "sort"
)
func main() {
	var a, b, c, d, i int

	Scan(&a, &b, &c)
	w := []int{a, b, c}
	Ints(w)
	for i < 3 {
		if []int{a, b, c}[i] == w[i] {
			d++
		}
		i++
	}

	if d < 1 {
		d = 2
	}
	if d > 2 {
		d = 0
	}

	Print(d)
}