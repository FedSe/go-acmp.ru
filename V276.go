package main
import . "fmt"
func main() {
	var m, n, i, j int
	P := Println
	Scan(&n, &m)

	for i < m-n%m {
		P(n / m)
		i++
	}

	for j < n%m {
		P(n/m + 1)
		j++
	}
}