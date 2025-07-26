package main
import . "fmt"
func main() {
	var (
		a [100][100]int
		n, i, l int
	)
	Scan(&n)
	for i < n {
		j := 0
		for j < n {
			Scan(&a[i][j])
		j++
		}
	i++
	}
	for l < n {
		i = 0
		for i < n {
			Println(a[i][l])
		i++
		}
	l++
	}
}