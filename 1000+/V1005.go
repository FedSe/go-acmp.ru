package main
import (
	. "fmt"
	. "sort"
)
func main() {
	var a, b, c, d, e, f, r, s int
	Scan(&a, &b, &c, &d, &e, &f, &r)

	q := [][]int{
		{a, d},
		{b, e},
		{c, f}}

	Slice(q, func(i, j int) bool {
		return q[i][0] < q[j][0]
	})

	for _, v := range q {
		a = v[0]
		b = v[1]
		if a < 1 {
			s += b
		} else if r >= a {
			if r/a < b {
				b = r / a
			}
			s += b
			r -= b * a
		}
	}

	Print(s)
}