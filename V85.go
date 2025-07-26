package main
import . "fmt"
func main() {
	var n, m, i int
	Scan(&n, &m)

	for n != m {
		if n > m {
			n -= m
		} else {
			m -= n
		}
	}

	for i < n {
		Print(1)
	i++
	}

}