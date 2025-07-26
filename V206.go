package main
import . "fmt"
type H struct {
	b, e, d int
}

func main() {
	var (
		x, y [101]int
		n, u, m int
		F = 99999
	)
	Scan(&n, &u, &m)

	r := make([][]H, n+1)
	for 0 < m {
	m--
		var l, h, g int
		Scan(&l, &h, &g)
		for
		j := 1
		j < l
		{
			w := 0
			k := 0
			Scan(&k, &w)
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
		for
		i := 0
		i < n
		{
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
	if n == F { n = -1 }
	Print(n)
}