package main
import . "fmt"
func main() {
	var (
		n, m, w, t, l, i int
		q                [100]struct{ a, b int }
		d                [9e4]int
		M                = 9 << 30
	)

	Scan(&n, &m)
	for t < m {
		Scan(&q[t].a, &q[t].b)
		if q[t].a > w {
			w = q[t].a
		}
		t++
	}

	w += n
	for l < w {
		l++
		d[l] = M
	}

	for i < w {
		if d[i] < M {
			for _, p := range q[:m] {
				l = i + p.a
				if l < w && d[i]+p.b < d[l] {
					d[l] = d[i] + p.b
				}
			}
		}
		i++
	}

	for n < w {
		if d[n] < M {
			M = d[n]
		}
		n++
	}

	Print(M)
}