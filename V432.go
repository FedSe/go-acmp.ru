package main
import . "fmt"
func main() {
	var (
		v             [200][200]bool
		g             [200]string
		d             = []int{-1, 1, 0, 0, -1, 1}
		N, M, u, j, i int
		F             func(r, c int)
	)

	Scan(&N, &M)
	for j < N {
		Scan(&g[j])
		j++
	}

	F = func(r, c int) {
		if r < 0 || r >= N || c < 0 || c >= M || v[r][c] || g[r][c] == 46 {
			return
		}
		v[r][c] = 1 > 0
		l := 0
		for l < 4 {
			F(r+d[l], c+d[l+2])
			l++
		}
	}

	for i < N {
		j = 0
		for j < M {
			if g[i][j] == 35 && !v[i][j] {
				u++
				F(i, j)
			}
			j++
		}
		i++
	}

	Print(u)
}