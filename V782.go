package main
import . "fmt"
func main() {
	n := 0
	s := 0

	Scan(&n)
	if n > 3 {
		s++
	}
	for n > 1 {
		s += n/2
		n = n/2 + n%2
	}

	Print(s)
}