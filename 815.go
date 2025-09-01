package main
import (
	. "fmt"
	. "sort"
)

type A map[int][]int
type B map[int]int

func main() {
	var (
		y                []int
		g                [2e3]A
		p                [1e3]int
		n, i, k, s, v, e int
		S                = Scan
		P                = Println
	)

	S(&n, &i, &k)
	for v <= n {
		g[v] = A{}
		v++
	}

	for 0 < i {
		S(&k, &v, &n)
		g[k][n] = append(g[k][n], v)
		i--
	}

	S(&v)
	for i < v {
		S(&p[i])
		i++
	}

	S(&s)
	o := B{s: 1}
	for e < v {
		k = p[e]
		q := B{}
		for h := range o {
			for _, a := range g[h][k] {
				q[a] = 1
			}
		}
		if len(q) < 1 {
			P("Hangs")
			return
		}
		o = q
		e++
	}

	for v := range o {
		y = append(y, v)
	}
	Ints(y)

	P("OK", len(y))
	for _, v := range y {
		P(v)
	}
}