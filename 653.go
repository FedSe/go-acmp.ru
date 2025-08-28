package main
import . "fmt"
func main() {
	var (
		n, m, l, i, k, u int
		g                [250][250]int
		s                = ""
		P                = Sprint
	)

	Scan(&n, &m)
	for l < n {
		j := 0
		for j < m {
			Scan(&g[l][j])
			j++
		}
		l++
	}

	for i < n {
		q := i*m + 1
		l = i
		for l < n {
			z := g[l][0]
			c := 1
			for c < m {
				if g[l][c] < z {
					z = g[l][c]
				}
				c++
			}
			if z == q {
				if l != i {
					g[i], g[l] = g[l], g[i]
					s += P(`
R `, i+1, l+1)
					u++
				}
				i++
				l = n
			}
			l++
		}

	}

	for k < m {
		if g[0][k] != k+1 {
			i = k + 1
			for i < m {
				if g[0][i] == k+1 {
					l = i
					i = 0
					for i < n {
						g[i][k], g[i][l] = g[i][l], g[i][k]
						i++
					}
					s += P(`
C `, k+1, l+1)
					u++
					i = m
				}
				i++
			}
		}
		k++
	}

	Print(u, s)
}