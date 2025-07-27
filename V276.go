package main
import . "fmt"
func main() {
	var m, n, i, j int
	Scan(&n, &m)

	for i < m-n%m {
		Println(n / m)
		i++
	}

	for j < n%m {
		Println(n/m + 1)
		j++
	}
}