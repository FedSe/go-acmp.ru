package main
import . "fmt"
func main() {
	var (
		n, m, u, v int
		e          = map[[2]int]int{}
		s          = "YES"
	)

	Scan(&n, &m)
	for 0 < m {
		Scan(&u, &v)
		if u > v {
			u, v = v, u
		}
		e[[2]int{u, v}] = 1
		m--
	}

	if len(e) < n {
		s = "NO"
	}
	Print(s)
}