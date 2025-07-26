package main
import . "fmt"
func main() {
	var (
		a [100][100]int
		n, m, i, c, l, f int
	)

	Scan(&n, &m)
	for c < n {
		g := 0
		j := 0
		for j < m {
			Scan(&a[c][j])
			g += a[c][j]
		j++
		}
		Println(g)
	c++
	}
	for l < m {
		c = 0
		i = 0
		for i < n {
			c += a[i][l]
		i++
		}
		Println(c)
	l++
	}
	for f < n {
		i = 0
		for i < m {
			Println(a[f][i])
		i++
		}
	f++
	}
}