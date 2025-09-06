package main
import . "fmt"
func main() {
	var (
		n, s, i, m int
		p          [100]int
	)

	Scan(&n)
	for i < n {
		Scan(&p[i])
		i++
	}
	for i > 0 {
		i--
		if p[i] > m {
			m = p[i]
		}
		s += m
	}

	Print(s)
}