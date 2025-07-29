package main
import . "fmt"
func main() {
	var (
		n, m, i, f, k int
		g, h          [30]string
		S             = Scan
	)

	S(&n)
	for k < n {
		S(&g[k])
		k++
	}

	S(&m)
	for f < m {
		S(&h[f])
		f++
	}

	for i < n {
		k = 0
		j := 0
		for j < m {
			l := len(g[i])
			if l == len(h[j]) {
				f = 0
				u := 0
				for u < l {
					if g[i][u] != h[j][u] {
						f++
					}
					u++
				}
				if f == 1 {
					k++
				}
			}
			j++
		}
		Println(k)
		i++
	}
}