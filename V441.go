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

type S struct {
	p    int
	m, c []int
}

func (s *S) H(w int) int {
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
	)

	Fscan(s, &n)

	F := func(n int) []T {
		a := make([]T, n)
		i := 0
		v := 0
		for i < n {
			Fscan(s, &v)
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
			Print(-1)
			return
		}
		c[j] = T{a[j].v, a[j].w, b[j].w}
		j++
	}
	Slice(c, func(i, j int) bool {
		return c[i].q < c[j].q
	})

	e := S{
		p: 400,
		m: make([]int, n),
		c: make([]int, (n+399)/400),
	}

	for _, t := range c {
		r += t.w + e.H(t.w+1) - t.q
		e.m[t.w] = 1
		e.c[t.w/e.p]++
	}

	Print(r)
}