package main
import . "fmt"
func main() {
	var (
		n, i, l int
		m       = 1
		g, q    []int
		h       [15][15]int
		d, p, e [4e4]int
		P       = Println
	)

	Scan(&n)
	f := 1 << n
	for i < n {
		j := 0
		for j < n {
			Scan(&h[i][j])
			j++
		}
		i++
	}

	for m < f {
		c := 0
		i = 0
		for i < n {
			if 1<<i&m > 0 {
				c++
				j := i
				for j < n {
					if c > 5 || 1<<j&m > 0 && h[i][j] < 1 {
						goto A
					}
					j++
				}
			}
			i++
		}
		g = append(g, m)
		q = append(q, c)
	A:
		m++
		d[m]--
	}

	m = 1
	for m < f {
		for x, u := range g {
			if m&u == u {
				i = d[m^u]
				if i > -1 {
					if q[x] > i {
						i = q[x]
					}
					if i > d[m] {
						d[m] = i
						p[m] = u
					}
				}
			}
		}
		m++
	}

	m = 1
	f--
	for f > 0 {
		i = 0
		for i < n {
			if 1<<i&p[f] > 0 {
				e[i] = m
			}
			i++
		}
		m++
		f ^= p[f]
	}

	P(m - 1)
	for l < n {
		P(e[l])
		l++
	}
}