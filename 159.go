package main
import . "fmt"
func main() {
	var (
		c          [2e4]int
		n, i, j, m int
	)

	Scan(&n)
	for i < n {
		Scan(&m)
		i++
		c[m-1] = i
	}

	for j < n {
		Println(c[j])
		j++
	}
}