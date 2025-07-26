package main
import . "fmt"
func main() {
	s := 0.
	n := -1.
	a := 1.

	for a > 0 {
		Scan(&a)
		s += a
		n++
	}

	if n != 0 {
		n = s / n
	}

	Print(n)
}