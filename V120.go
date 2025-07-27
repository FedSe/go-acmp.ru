package main
import . "fmt"
func main() {
	var (
		n, m, j, i int
		a          [20][20]int
	)
	Scan(&n, &m)

	for j < n {
		v := 0
		for v < m {
			Scan(&a[j][v])
			v++
		}
		j++
	}

	for i < n {
		j = 0
		for j < m {
			v := 0
			if i > 0 {
				v = a[i-1][j]
			}
			if j > 0 {
				v = a[i][j-1]
				if i*j > 0 {
					w := a[i-1][j]
					if w < v {
						v = w
					}
				}
			}

			a[i][j] += v
			j++
		}
		i++
	}

	Print(a[n-1][m-1])
}