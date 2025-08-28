package main
import . "fmt"

var (
	v             [200][200]bool
	g             [200]string
	d             = [4]int{-1, 1}
	N, M, u, j, i int
)

func F(r, c int) {
	if r < 0 || r >= N || c < 0 || c >= M || v[r][c] || g[r][c] == 46 {
		return
	}
	v[r][c] = 1 > 0
	l := 0
	for l < 4 {
		F(r+d[l], c-d[3-l])
		l++
	}
}

func main() {
	Scan(&N, &M)
	for j < N {
		Scan(&g[j])
		j++
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