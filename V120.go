package main
import . "fmt"
func main() {
	var (
		n, m, i int
		a [20][20]int
	)
	Scan(&n, &m)

	for i < n {
		for
		j := 0
		j < m
		{
			Scan(&a[i][j])
		j++
		}
	i++
	}

	i = 1
	for i < 20 {
		a[0][i] += a[0][i-1]
		a[i][0] += a[i-1][0]
	i++
	}

	i = 1
	for i < n {
		for
		j := 1
		j < m
		{
			r := a[i-1][j]
			e := a[i][j-1]
			a[i][j] = a[i][j] + e
			if e > r {
				a[i][j] += r - e
			}
		j++
		}
	i++
	}

	Print(a[n-1][m-1])
}