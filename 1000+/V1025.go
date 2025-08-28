package main
import . "fmt"
func main() {
	for {
		var (
			N, M, l, i, G int
			P             = Print
			q, z          [3e4]int
		)
		G = 1e9

		Scan(&N, &M)
		if N*M < 1 {
			break
		}
		g := make([][]int, N)

		for l < N {
			g[l] = make([]int, M)
			j := 0
			for j < M {
				Scan(&g[l][j])
				j++
			}
			l++
		}
		if N*M > 1 {
			l = 1
			for l < M {
				q[l] = 1e9
				z[l] = -1e9
				l++
			}
			l = g[0][0]
			q[0] = l
			z[0] = l

			for i < N {
				l = 0
				for l < M {
					if i+l > 0 {
						a := G
						b := -G
						o := q[l] + g[i][l]
						if i > 0 {
							if o < a {
								a = o
							}
							o = z[l] + g[i][l]
							if o > b {
								b = o
							}
						}
						if l > 0 {
							o = q[l-1] + g[i][l]
							if o < a {
								a = o
							}
							o = z[l-1] + g[i][l]
							if o > b {
								b = o
							}
						}
						q[l] = a
						z[l] = b
					}
					l++
				}
				i++
			}
		}
		M--
		if q[M] == z[M] {
			P("Y")
		} else {
			P("N")
		}
	}
}