package main
import . "fmt"
func main() {
	n := 0
	s := ""
	w := s

	Scan(&n)
	for 0 < n {
		Scan(&w)
		s = w + " " + s
		n--
	}

	Print(s)
}