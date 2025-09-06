package main
import . "fmt"
func main() {
	x := 0
	s := "No"
	a := s

	Scan(&a, &x)
	if x&1 > 0 || a[0] < 73 {
		s = "Yes"
	}

	Print(s)
}