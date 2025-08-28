package main
import . "fmt"
func main() {
	n := 0
	s := ""

	Scan(&n)
	for n > 0 {
		l := n % 3
		if l < 1 {
			l = 3
		}
		s = Sprint(l) + s
		n -= l
		n /= 3
	}

	Print(s)
}