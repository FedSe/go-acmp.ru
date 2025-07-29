package main
import . "fmt"

type H struct {
	b, e, d int
}

func main() {
	var (
		x, y    [101]int
		n, u, m int
		F       = 99999
		S       = Scan
	)

	S(&n, &u, &m)
	r := make([][]H, n+1)

	for 0 < m {
		m--
		var l, h, g int
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

	for m <= n {
		y[m] = F
		m++
	}
	y[1] = 0

	for {
		v := -1
		m = F
		i := 0
		for i < n {
			i++
			if x[i] < 1 && y[i] < m {
				m = y[i]
				v = i
			}
		}
		if m == F {
			break
		}
		x[v] = 1
		for _, o := range r[v] {
			if o.b >= y[v] {
				if o.e < y[o.d] {
					y[o.d] = o.e
				}
			}
		}
	}

	n = y[u]
	if n == F {
		n = -1
	}
	Print(n)
}