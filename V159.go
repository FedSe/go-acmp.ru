package main
import . "fmt"
func main() {
	var (
		c [20001]int
		n, i, j, m int
	)
	Scan(&n)

	for i < n {
		Scan(&m)
	i++
		c[m-1] = i
	}

	for j < n {
		Print(c[j], " ")
	j++
	}
}