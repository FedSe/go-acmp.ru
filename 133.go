package main
import . "fmt"
type I struct{ n, d int }
func main() {
	var (
		c, d       [100]int
		n, m, l, w int
		g          [100][900]I
		p          = []I{{}}
		S          = Scan
	)

	S(&n)
	for l < n {
		S(&c[l])
		l++
	}
	S(&m)

	for 0 < m {
		l = 0
		v := 0
		S(&l, &v)
		l--
		v--
		g[l][w] = I{v, c[l]}
		g[v][w] = I{l, c[v]}
		w++
		m--
	}

	l = 1
	for l < n {
		d[l] = 1e9
		l++
	}

	w = len(p)
	for w > 0 {
		w--
		m = p[w].n
		p = p[:w]
		for _, e := range g[m] {
			l = e.n
			e := d[m] + e.d
			if e < d[l] {
				d[l] = e
				p = append(p, I{l, e})
				w++
			}
		}
	}

	l = d[n-1]
	if l > 1e8 {
		l = -1
	}

	Print(l)
}