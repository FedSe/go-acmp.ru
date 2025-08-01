package main
import . "fmt"
func main() {
	s := "second"
	a := s
	b := s

	Scan(&a, &b)

	if a == b {
		s = "draw"
	}

	switch a[0] - b[0] {
	case 254, 255, 3:
		s = "first"
	}

	Print(s)
}