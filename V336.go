package main
import . "fmt"
func main() {
	var a, m, l int
	s := ""

	Scan(&s)
	for _, r := range s {
		a--
		if r > 49 {
			a += 2
		}
		if a < m {
			m = a
		}
		if a > l {
			l = a
		}
	}

	Print(l - m + 1)
}