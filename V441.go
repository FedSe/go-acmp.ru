package main
import (
	d "bufio"
	. "fmt"
	. "os"
	. "strconv"
	. "sort"
)

type T struct {
	v, w, q int
}

type S struct {
	p int
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
	var n, v, i, r, l, j int
	o := 400
	s := d.NewScanner(Stdin)
	s.Split(d.ScanWords)
	Scan(&n)

	a := make([]T, n)
	for l < n && s.Scan() {
		v, _ = Atoi(s.Text())
		a[l] = T{v: v, w: l}
	l++
	}
	SliceStable(a, func(i, j int) bool {
		return a[i].v < a[j].v
	})

	b := make([]T, n)
	for s.Scan() {
		v, _ = Atoi(s.Text())
		b[i] = T{v: v, w: i}
	i++
	}
	SliceStable(b, func(i, j int) bool {
		return b[i].v < b[j].v
	})

	c := make([]T, n)
	for j < n {
		if a[j].v != b[j].v {
			Print(-1)
			return
		}
		c[j] = T{v: a[j].v, w: a[j].w, q: b[j].w}
	j++
	}
	Slice(c, func(i, j int) bool {
		return c[i].q < c[j].q
	})

	var e S
	e.p = o
	e.m = make([]int, n)
	e.c = make([]int, (n+o-1)/o)

	for _, t := range c {
		r += t.w + e.H(t.w+1) - t.q
		e.m[t.w] = 1
		e.c[t.w/e.p]++
	}

	Print(r)
}