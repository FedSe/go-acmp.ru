package main
import . "fmt"
func main() {
	var (
		a       [100][100]int
		n, i, j int
	)

	Scan(&n)
	for i < n {
		j = 0
		for j < n {
			Scan(&a[i][j])
			j++
		}
		i++
	}

	n--
	i = n
	for i >= 0 {
		j = n
		for j >= 0 {
			Println(a[j][i])
			j--
		}
		i--
	}
}