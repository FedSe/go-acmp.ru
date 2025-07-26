package main
import . "fmt"
func main() {
	var (
		n, m, i, f, k int
		g, h [30]string
	)

	Scan(&n)
	for k < n {
		Scan(&g[k])
	k++
	}

	Scan(&m)
	for f < m {
		Scan(&h[f])
	f++
	}

	for i < n {
		k = 0
		for
		j := 0
		j < m
		j++ {
			l := len(g[i])
			if l != len(h[j]) {
				continue
			}
			f = 0

			for
			u := 0
			u < l
			{
				if g[i][u] != h[j][u] {
					f++
				}
			u++
			}

			if f == 1 {
				k++
			}
		}
		Print(k, " ")
	i++
	}
}