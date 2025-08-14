package main
import . "fmt"
func main() {
	n := 0
	m := 0

	Scan(&n, &m)
	for n != m {
		if n > m {
			n, m = m, n
		}
		m -= n
	}

	for 0 < n {
		Print(1)
		n--
	}
}