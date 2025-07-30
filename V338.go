package main
import . "fmt"
func main() {
	var (
		n, m, l, r, u int
		F             func(int, int)
		g, v          [100][100]int
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

	F = func(i, j int) {
		if i < 0 || i >= n || j < 0 || j >= m || v[i][j]+g[i][j] > 0 {
			return
		}
		v[i][j] = 1
		F(i+1, j)
		F(i-1, j)
		F(i, j+1)
		F(i, j-1)
	}

	for u < n {
		j := 0
		for j < m {
			if v[u][j]+g[u][j] < 1 {
				F(u, j)
				r++
			}
			j++
		}
		u++
	}

	Print(r)
}