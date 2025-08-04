package main
import (
	. "container/heap"
	. "fmt"
)

type I struct{ n, d, i int }
type P []*I

func (p P) Len() int           { return len(p) }
func (p P) Less(i, j int) bool { return 1 < 0 }
func (p P) Swap(i, j int)      {}
func (p *P) Push(x any) {
	m := x.(*I)
	m.i = len(*p)
	*p = append(*p, m)
}
func (p *P) Pop() any {
	o := *p
	n := len(o)
	t := o[n-1]
	*p = o[:n-1]
	return t
}

func main() {
	var (
		c, d    [100]int
		N, M, l int
		p       = &P{}
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
		g[l] = append(g[l], I{v, c[l], 0})
		g[v] = append(g[v], I{l, c[v], 0})
		M--
	}

	l = 1
	for l < N {
		d[l] = 1e9
		l++
	}

	Push(p, &I{})
	for p.Len() > 0 {
		M = Pop(p).(*I).n
		for _, e := range g[M] {
			l = e.n
			e := d[M] + e.d
			if e < d[l] {
				d[l] = e
				Push(p, &I{l, e, 0})
			}
		}
	}

	l = d[N-1]
	if l > 1e8 {
		l = -1
	}
	Print(l)
}