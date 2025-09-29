package main
import . "fmt"
type T [2e3]int
func main() {
	var (
		q, w, z, g    T
		p             [101]T
		a, t, h, i, j int
	)

	Scan(&a, &t)
	for i < a {
		i++
		Scan(&q[i], &w[i], &z[i], &g[i])
		if w[i] > h {
			h = w[i]
		}
	}

	h += t - 1
	for i := range p {
		for j := range p[i] {
			p[i][j] = 2e9
		}
	}

	p[0][0] = 0
	for j < a {
		j++
		i := 0
		for i < h {
			k := 0
			for k <= i && k <= g[j] {
				o := z[j]
				if k < w[j] {
					o = q[j]
				}
				v := p[j-1][i-k] + o*k
				if v < p[j][i] {
					p[j][i] = v
				}
				k++
			}
			i++
		}
	}

	j = p[a][t]
	for t < h {
		t++
		if p[a][t] < j {
			j = p[a][t]
		}
	}

	if j > 1e9 {
		j = -1
	}

	Print(j)
}