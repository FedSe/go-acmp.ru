package main
import . "fmt"
func main() {
	var s, n, a int
	p := 3

	Scan(&n)
	for 0 < n {
		Scan(&a)

		if a > 0 {
			s += p
			p += 4
		}

		p -= 3

		if p < 4 {
			p = 3
		}
	n--
	}

	Print(s)
}