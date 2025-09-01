package main
import (
	. "fmt"
	. "sort"
)
type A struct{ L, R int }
func main() {
	var (
		n, b, c, d, i int
		q             [2e4]A
	)

	Scan(&n)
	for i < n {
		Scan(&q[i].L, &q[i].R)
		i++
	}

	Slice(q[:n], func(i, j int) bool {
		return q[i].L < q[j].L
	})

	for _, v := range q {
		if i > 0 {
			c = v.L
			d = v.R
			i = 0
		}
		if v.L > d {
			b += d - c
			c = v.L
			d = v.R
		} else if v.R > d {
			d = v.R
		}
	}

	Print(b + d - c)
}