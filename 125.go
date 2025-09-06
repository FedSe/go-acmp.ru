package main
import . "fmt"
func main() {
	var (
		n, k, i, m int
		b          [101][100]int
	)

	Scan(&n)
	for i <= n {
		j := 0
		for j < n {
			Scan(&b[i][j])
			j++
		}
		i++
	}

	for m < n {
		i = m + 1
		for i < n {
			if b[m][i] > 0 && b[n][m] != b[n][i] {
				k++
			}
			i++
		}
		m++
	}

	Print(k)
}