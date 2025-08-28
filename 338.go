package main
import . "fmt"

var (
	n, m, l, r, u int
	g, v          [100][100]int
)

func F(i, j int) {
	if i < 0 || i >= n || j < 0 || j >= m || v[i][j]+g[i][j] > 0 {
		return
	}
	v[i][j] = 1
	F(i+1, j)
	F(i-1, j)
	F(i, j+1)
	F(i, j-1)
}

func main() {
	Scan(&n, &m)

	for l < n {
		j := 0
		for j < m {
			Scan(&g[l][j])
			j++
		}
		l++
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