package main
import . "fmt"
type A [20]int
func main() {
	var (
		v          [20]A
		a          [20]string
		N, o, i, j int
		q          = A{0, 1, 0, -1}
	)

	Scan(&N)
	for j < N {
		Scan(&a[j])
		j++
	}
	for i < N {
		j = 0
		for j < N {
			if a[i][j] == 66 && v[i][j] < 1 {
				z := []A{{i, j}}
				e := map[A]int{}
				v[i][j] = 1
				for len(z) > 0 {
					d := len(z) - 1
					x := z[d][0]
					y := z[d][1]
					z = z[:d]
					d = 0
					for d < 4 {
						h := x + q[d]
						f := y + q[d+1]
						if h >= 0 && f >= 0 && h < N && f < N {
							w := a[h][f]
							p := A{h, f}
							if w == 66 && v[h][f] < 1 {
								v[h][f] = 1
								z = append(z, p)
							}
							if w == 46 {
								e[p] = 1
							}
						}
						d++
					}
				}
				if len(e) == 1 {
					o++
				}
			}
			j++
		}
		i++
	}

	Print(o)
}