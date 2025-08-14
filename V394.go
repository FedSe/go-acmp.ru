package main
import . "fmt"
func main() {
	n := 0
	m := 0

	Scan(&n, &m)
	a := n
	for n*m > 0 {
		if n > m {
			n, m = m, n
		}
		m %= n
	}

	Print(a / (n+m))
}