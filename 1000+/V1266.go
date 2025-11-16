package main
import . "fmt"
func main() {
	a := 0
	b := 0

	Scan(&a, &b)
	if a > b {
		a, b = b, a
	}

	a--
	b--
	b = b/2 - a/2
	if b < 1 {
		b = 1
	}

	Print(b - 1)
}