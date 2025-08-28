package main
import . "fmt"
func main() {
	var (
		n, k, i, m int
		b [101][100]int
	)

	Scan(&n)
	for i <= n {
		for
		j := 0
		j < n
		{
			Scan(&b[i][j])
		j++
		}
	i++
	}

	for m < n {
		for
		i = m+1
		i < n
		{
			if b[m][i] > 0 && b[n][m] != b[n][i] {
				k++
			}
		i++
		}
	m++
	}

	Print(k)
}