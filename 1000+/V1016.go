package main
import (
	. "fmt"
	. "sort"
)
type T struct{ a, c, b string }
func main() {
	var (
		a, b, c, d string
		n, x, y    int
		S          = Sscan
	)

	Scan(&n)
	q := make([]T, n)
	for 0 < n {
		n--
		Scan(&a, &b, &c, &d)
		q[n] = T{a, c, b + " " + d}
	}

	Slice(q, func(i, j int) bool {
		a = q[i].c
		b = q[j].c
		n = len(a) - 1
		w := len(b) - 1
		S(a[:n], &x)
		S(b[:w], &y)
		if x != y {
			return x < y
		}
		if a[n] != b[w] {
			return a[n] < b[w]
		}
		return q[i].a < q[j].a
	})

	for _, s := range q {
		Println(s.c, s.a, s.b)
	}
}