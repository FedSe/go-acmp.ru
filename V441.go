package main

import (
	d "bufio"
	. "fmt"
	. "os"
	. "sort"
)

type T struct {
	v, w, q int
}

type H struct {
	p    int
	m, c []int
}

func (s *H) H(w int) int {
	d := (w + s.p - 1) / s.p * s.p
	y := 0
	for w < len(s.m) && w < d {
		if s.m[w] > 0 {
			y++
		}
		w++
	}
	i := d / s.p
	for i < len(s.c) {
		y += s.c[i]
		i++
	}
	return y
}

func main() {
	var (
		n, r, j int
		s       = d.NewReader(Stdin)
		S       = Fscan
		P       = Print
	)

	S(s, &n)

	F := func(n int) []T {
		a := make([]T, n)
		i := 0
		v := 0
		for i < n {
			S(s, &v)
			a[i] = T{v, i, 0}
			i++
		}
		SliceStable(a, func(i, j int) bool {
			return a[i].v < a[j].v
		})
		return a
	}

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