package main
import . "fmt"
func main() {
	var (
		a       [100][100]int
		n, m, i int
	)

	Scan(&n, &m)
	for i < n {
		j := 0
		for j < m {
			Scan(&a[i][j])
			j++
		}
		i++
	}

	for n > 0 {
		n--
		i = 0
		for i < m {
			Println(a[n][i])
			i++
		}
	}
}