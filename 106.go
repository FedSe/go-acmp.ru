package main
import . "fmt"
func main() {
	var n, m, j, i int

	Scan(&n)
	for i < n {
		Scan(&m)
		j += m
		i++
	}

	if n < j*2 {
		j = n - j
	}

	Print(j)
}