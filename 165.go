package main
import . "fmt"
func main() {
	var (
		a, d       [70][70]int
		n, m, i, j int
	)

	Scan(&n, &m)
	for j < n {
		l := 0
		for l < m {
			Scan(&a[j][l])
			l++
		}
		j++
	}
	d[0][0] = 1
	for i < n {
		j = 0
		for j < m {
			v := a[i][j]
			w := d[i][j]
			if w*v > 0 {
				if i+v < n {
					d[i+v][j] += w
				}
				if j+v < m {
					d[i][j+v] += w
				}
			}
			j++
		}
		i++
	}

	Print(d[n-1][m-1])
}