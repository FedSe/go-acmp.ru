package main
import . "fmt"
func main() {
	var (
		a, m [51][51]int
		n, i, l, j, v int
	)
	Scan(&n)

	
	for i < n {
	i++
		v = 0
		for v < n {
		v++
			Scan(&a[i][v])
		}
	}

	for l <= n {
		v = 0
		for v <= n {
			m[l][v] = -1e9
		v++
		}
	l++
	}

	m[0][0] = 0

	for j < n {
	j++
		l = 0
		for l <= n {
			h := &m[j][l]
			*h = m[j-1][l]
			i = 0
			for i < l {
			i++
				v = a[i][j] + m[j-1][l-i]
				if *h <  v {
					*h = v
				}
			}
		l++
		}
	}

	Print(m[n][n])
}