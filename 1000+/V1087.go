package main
import . "fmt"
func main() {
	s := "second"
	a := s
	b := s

	Scan(&a, &b)

	if (a[0] == 112 && b[0] == 114) || (a[0] == 114 && b[0] == 115) || (a[0] == 115 && b[0] == 112) {
		s = "first"
	}

	if a == b {
		s = "draw"
	}

	Print(s)
}