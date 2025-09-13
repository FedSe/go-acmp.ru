package main
import . "fmt"
func main() {
	n := 0
	m := 0
	q := 1

	Scan(&n, &m)
	m++
	if n > 0 {
		q = 2 * n
	}

	Print(q * m)
}