package main
import (
	b "bufio"
	. "fmt"
	. "os"
	. "sort"
)

type F = float64
type H struct{ x, y, z F }

func main() {
	var (
		N, a, i, j, l int
		c, d, e, z    F
		r             = b.NewReader(Stdin)
	)

	Scan(&N, &N, &N)
	h := make([]H, N)
	for i < N {
		Fscan(r, &c, &d, &e, &z)
		h[i] = H{c, d, e}
		i++
	}
	Slice(h, func(i, j int) bool {
		m := h[i].x
		o := h[i].y
		v := h[j].x
		t := h[j].y
		if m < 0 && v >= 0 {
			return 1 > 0
		}
		if v < 0 && m >= 0 {
			return 0 > 1
		}
		if m <= 0 && v <= 0 {
			return o > t
		}
		return o < t
	})
	g := make([]F, N)
	for j < N {
		g[j] = h[j].z
		j++
	}
	u := make([]F, N)
	copy(u, g)
	Float64s(u)
	q := u[:0]
	for i, x := range u {
		if i < 1 || x-u[i-1] > 1e-8 {
			q = append(q, x)
		}
	}
	k := make([]int, N)
	for i, x := range g {
		k[i] = SearchFloat64s(q, x-1e-8/2) + 1
	}
	i = len(q)
	t := make([]int, i+1)
	for l < N {
		j = k[l]
		s := 0
		for j > 0 {
			s += t[j]
			j -= j & -j
		}
		a += l - s
		j = k[l]
		for j <= i {
			t[j]++
			j += j & -j
		}
		l++
	}

	Print(a)
}