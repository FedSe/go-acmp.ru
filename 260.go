package main
import (
	. "fmt"
	. "sort"
)
type A struct {
	a, b int
}
func main() {
	var (
		w             []A
		g             [2e3][]A
		d, h          [2e3]int
		n, m, k, c, i int
		S             = Scan
	)

	S(&n, &m, &k, &c)
	for i < k {
		S(&h[i])
		i++
	}

	for 0 < n {
		d[n] = 1 << 30
		n--
	}

	d[c] = 0
	p := []A{{c, 0}}
	for 0 < m {
		S(&i, &n, &c)
		g[i] = append(g[i], A{n, c})
		g[n] = append(g[n], A{i, c})
		m--
	}

	for len(p) > 0 {
		m = len(p) - 1
		y := p[m]
		p = p[:m]
		for _, v := range g[y.a] {
			n = v.a
			a := d[y.a] + v.b
			if a < d[n] {
				d[n] = a
				p = append(p, A{n, a})
			}
		}
	}

	for _, v := range h[:k] {
		w = append(w, A{v, d[v]})
	}

	Slice(w, func(i, j int) bool {
		a := w[i].b
		b := w[j].b
		u := a < b
		if a == b {
			u = w[i].a < w[j].a
		}
		return u
	})

	for _, v := range w {
		Println(v.a, v.b)
	}
}