package main
import . "fmt"
type I = int
type T map[any]I
func main() {
	var (
		n, m, p, b, i, u I
		f                [20]T
		s                [20]I
		k                [20][]I
		t                [101][]I
		w, z, q          []I
		d                = T{}
		S                = Scan
	)

	S(&n, &m)
	for i < m {
		S(&p)
		e := 0
		h := T{}
		j := 0
		for j < p {
			S(&e)
			h[e] = 1
			t[e] = append(t[e], i)
			j++
		}
		f[i] = h
		i++
	}

	S(&p, &b)
	for u < m {
		if f[u][p] > 0 {
			w = append(w, u)
		}
		if f[u][b] > 0 {
			z = append(z, u)
		}
		u++
	}

	for 0 < m {
		m--
		s[m]--
	}

	for m < n {
		m++
		g := t[m]
		i = 0
		x := len(g)
		for i < x {
			p = i + 1
			for p < x {
				b = g[i]
				u = g[p]
				y := [2]I{b, u}
				c := [2]I{u, b}
				if d[y]+d[c] < 1 {
					k[b] = append(k[b], u)
					k[u] = append(k[u], b)
					d[y] = 1
					d[c] = 1
				}
				p++
			}
			i++
		}
	}

	for _, v := range w {
		s[v] = 0
		q = append(q, v)
	}

	b = -1
	for len(q) > 0 {
		i = q[0]
		q = q[1:]
		for _, v := range z {
			if v == i {
				b = s[i]
				goto A
			}
		}
		for _, v := range k[i] {
			if s[v] < 0 {
				s[v] = s[i] + 1
				q = append(q, v)
			}
		}
	}
A:
	Print(b)
}