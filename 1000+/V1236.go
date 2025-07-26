package main
import . "fmt"
func main() {
	var (
		a [100][100]int
		n, m, i, l int
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

	for l < n {
		for
		i = 0
		i < m/2
		{
			a[l][i], a[n-l-1][i] = a[n-l-1][i], a[l][i]
		i++
		}
	l++
	}

	for n > 0 {
	n--
		for
		i = 0
		i < m
		{
			Println(a[n][i])
		i++
		}
	}
}