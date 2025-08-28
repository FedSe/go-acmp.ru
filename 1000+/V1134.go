package main
import . "fmt"
func main() {
	s := 0.
	a := 1.
	n := -a

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