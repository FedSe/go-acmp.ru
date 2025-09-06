package main
import . "fmt"
func main() {
	var s, p, c int

	Scan(&s, &p)
	for c*s < p {
		s--
		c++
	}
	if s < c {
		s, c = c, s
	}

	Print(c, s)
}