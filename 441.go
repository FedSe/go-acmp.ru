package main
import (
	d "bufio"
	. "fmt"
	. "os"
	. "sort"
)

type T struct{ v, w, q int }
type H struct {
	p    int
	m, c []int
}

var (
	n, r, j int
	s       = d.NewReader(Stdin)
	P       = Print
)

func (s *H) H(w int) int {
	d := (w + s.p - 1) / s.p * s.p
	y := 0
	for w < len(s.m) && w < d {
		if s.m[w] > 0 {
			y++
		}
		w++
	}
	d /= s.p
	for d < len(s.c) {
		y += s.c[d]
		d++
	}
	return y
}

func F(n int) []T {
	a := make([]T, n)
	v := 0
	for 0 < n {
		n--
		Fscan(s, &v)
		a[n] = T{v, n, 0}
	}
	SliceStable(a, func(i, j int) bool {
		return a[i].v < a[j].v
	})
	return a
}

func main() {
	Scan(&n)
	a := F(n)
	b := F(n)
	c := make([]T, n)
	for j < n {
		if a[j].v != b[j].v {
			P(-1)
			return
		}
		c[j] = T{a[j].v, a[j].w, b[j].w}
		j++
	}
	Slice(c, func(i, j int) bool {
		return c[i].q < c[j].q
	})

	e := H{400, make([]int, n), make([]int, (n+399)/400)}

	for _, t := range c {
		r += t.w + e.H(t.w+1) - t.q
		e.m[t.w] = 1
		e.c[t.w/e.p]++
	}

	P(r)
}