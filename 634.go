package main
import . "fmt"
func main() {
	var (
		g, a, d                   [128][7]int
		p                         []int
		t                         [7]int
		n, k, q, w, i, l, j, z, x int
		e                         = 1
		N                         = 1 << 20
		h                         = N
		P                         = Println
		S                         = Scan
	)

	S(&n, &k)
	for i < n {
		j := 0
		for j < n {
			S(&a[i][j])
			j++
		}
		i++
	}

	for j < n {
		S(&t[j])
		j++
	}

	for i := range d {
		j = 0
		for j < n {
			d[i][j] = N
			g[i][j] = -1
			j++
		}
	}

	for l < n {
		d[1<<l][l] = t[l]
		l++
	}

	for e < 1<<n {
		i = 0
		for i < n {
			j = 0
			for j < n {
				s := e | 1<<j
				l = d[e][i] + a[i][j] + t[j]
				if l < d[s][j] {
					d[s][j] = l
					g[s][j] = i
				}
				j++
			}
			i++
		}
		e++
	}

	for x < 1<<n {
		j = x
		l = 0
		for j > 0 {
			l += j & 1
			j >>= 1
		}
		if l == k {
			i = 0
			for i < n {
				if d[x][i] < h {
					h = d[x][i]
					q = x
					w = i
				}
				i++
			}
		}
		x++
	}

	for {
		p = append(p, w)
		i = g[q][w]
		if i < 0 {
			break
		}
		q ^= 1 << w
		w = i
	}

	e = len(p)
	for z < e/2 {
		i = e - 1 - z
		p[z], p[i] = p[i], p[z]
		z++
	}

	P(h)
	for _, v := range p {
		P(v + 1)
	}
}