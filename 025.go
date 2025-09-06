package main
import . "fmt"
func main() {
	a := 0
	b := 0
	s := "="

	Scan(&a, &b)
	if a > b {
		s = ">"
	}
	if a < b {
		s = "<"
	}

	Print(s)
}