package main
import . "fmt"
func main() {
	var (
		n, m, w, t, l, i int
		d, q, h          [4e4]int
		M                = 9 << 30
	)

	Scan(&n, &m)
	for t < m {
		Scan(&q[t], &h[t])
		if q[t] > w {
			w = q[t]
		}
		t++
	}

	w += n
	for l < w {
		l++
		d[l] = M
	}

	for i < w {
		j := 0
		for j < m {
			l = i + q[j]
			t = d[i] + h[j]
			if t < d[l] {
				d[l] = t
			}
			j++
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