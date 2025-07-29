package main
import . "fmt"
func main() {
	var (
		a, b, c             [100][100]int
		n, m, p, i, l, j, w int
		S = Scan
	)

	S(&n, &m, &p)
	for i < n {
		k := 0
		for k < m {
			S(&a[i][k])
			k++
		}
		i++
	}

	for l < m {
		i = 0
		for i < p {
			S(&b[l][i])
			i++
		}
		l++
	}

	for j < n {
		i = 0
		for i < p {
			c[j][i] = 0
			l = 0
			for l < m {
				c[j][i] += a[j][l] * b[l][i]
				l++
			}
			i++
		}
		j++
	}

	for w < n {
		i = 0
		for i < p {
			Println(c[w][i])
			i++
		}
		w++
	}
}