package main
import . "fmt"
func main() {
	var (
		a, b, c [100][100]int
		n, m, p, i, l, j, w int
	)

	Scan(&n, &m, &p)
	for i < n {
		for
		k := 0
		k < m
		{
			Scan(&a[i][k])
		k++
		}
	i++
	}

	for l < m {
		for
		i = 0
		i < p
		{
			Scan(&b[l][i])
		i++
		}
	l++
	}

	for j < n {
		for
		i = 0
		i < p
		{
			c[j][i] = 0
			for
			v := 0
			v < m
			{
				c[j][i] += a[j][v] * b[v][i]
			v++
			}
		i++
		}
	j++
	}

	for w < n {
		for
		i = 0
		i < p
		{
			Print(c[w][i], " ")
		i++
		}
		Println()
	w++
	}
}