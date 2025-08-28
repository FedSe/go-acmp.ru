package main
import (
	. "fmt"
	. "sort"
)
func main() {
	var n, d, l, i, s int

	Scan(&n, &d)
	p := make([]int, n)
	for l < n {
		Scan(&p[l])
		l++
	}

	n--
	Ints(p)
	for i != n {
		if p[i]+p[n] <= d {
			i++
		}
		s++
		n--
		if i > n {
			s--
			i = n
		}
	}

	Print(s+1)
}