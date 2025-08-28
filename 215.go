package main
import . "fmt"
type T [100]int
func main() {
	var (
		d, g          [100]T
		n, m, c, j, i int
	)

	Scan(&n, &m)
	for j < n {
		l := 0
		for l < m {
			Scan(&g[j][l])
			l++
		}
		j++
	}

	for i < n {
		j = 0
		for j < m {
			if d[i][j] < 1 {
				s := []T{{i, j}}
				h := g[i][j]
				o := 0
				d[i][j] = 1
				for len(s) > 0 {
					k := len(s) - 1
					x := s[k][0]
					y := s[k][1]
					s = s[:k]
					k = 0
					for k < 4 {
						e := x + k%2 - k/3*2
						b := y + 1 - k + k/3*2
						a := 10001
						if e >= 0 && e < n && b >= 0 && b < m {
							a = g[e][b]
							if d[e][b] < 1 && a == h {
								d[e][b] = 1
								s = append(s, T{e, b})
							}
						}
						if a < h {
							o = 1
						}
						k++
					}
				}
				if o < 1 {
					c++
				}
			}
			j++
		}
		i++
	}

	Print(c)
}