package main
import . "fmt"
func main() {
	var (
		g          [16][]byte
		w          = [31][16][16]int{{{1}}}
		z          = []int{0, 0, 1, -1}
		N, K, i, s int
	)

	Scan(&N, &K)
	for i < N {
		s := ""
		Scan(&s)
		g[i] = []byte(s)
		i++
	}

	for s < K {
		i = 0
		for i < N {
			j := 0
			for j < N {
				v := w[s][i][j]
				if v > 0 && g[i][j] != 49 {
					d := 0
					for d < 4 {
						y := i + z[d]
						u := j - z[3-d]
						if y >= 0 && y < N && u >= 0 && u < N && g[y][u] == 48 {
							w[s+1][y][u] += v
						}
						d++
					}
				}
				j++
			}
			i++
		}
		s++
	}

	Print(w[K][N-1][N-1])
}