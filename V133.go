package main
import . "fmt"
type I struct{ n, d int }
func main() {
	var (
		c, d    [100]int
		N, M, l int
		p       []I
	)

	Scan(&N)
	for l < N {
		Scan(&c[l])
		l++
	}
	Scan(&M)

	g := make([][]I, N)
	for 0 < M {
		l = 0
		v := 0
		Scan(&l, &v)
		l--
		v--
		g[l] = append(g[l], I{v, c[l]})
		g[v] = append(g[v], I{l, c[v]})
		M--
	}

	l = 1
	for l < N {
		d[l] = 1e9
		l++
	}

	p = append(p, I{})
	for len(p) > 0 {
		n := len(p)
		M = p[n-1].n
		p = p[:n-1]
		for _, e := range g[M] {
			l = e.n
			e := d[M] + e.d
			if e < d[l] {
				d[l] = e
				p = append(p, I{l, e})
			}
		}
	}

	l = d[N-1]
	if l > 1e8 {
		l = -1
	}
	Print(l)
}