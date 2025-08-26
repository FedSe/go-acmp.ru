package main

import . "fmt"

type H struct{ b, e, d int }

func main() {
	var (
		x, y             [101]int
		r                [101][]H
		n, u, m, l, h, g int
		z                = 1
		F                = 99999
		S                = Scan
	)

	S(&n, &u, &m)
	for 0 < m {
		m--
		S(&l, &h, &g)
		j := 1
		for j < l {
			w := 0
			k := 0
			S(&k, &w)
			r[h] = append(r[h], H{g, w, k})
			h = k
			g = w
			j++
		}
	}

	for z < n {
		z++
		y[z] = F
	}

	for m != F {
		z = 0
		m = F
		i := 0
		for i < n {
			i++
			if x[i] < 1 && y[i] < m {
				m = y[i]
				z = i
			}
		}

		x[z] = 1
		for _, o := range r[z] {
			if o.b >= y[z] && o.e < y[o.d] {
				y[o.d] = o.e
			}
		}
	}

	n = y[u]
	if n == F {
		n = -1
	}

	Print(n)
}
