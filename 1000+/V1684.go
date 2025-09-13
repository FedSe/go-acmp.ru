package main
import . "fmt"
func main() {
	a := 0
	b := 0

	Scan(&a, &b)
	for b > 0 {
		a, b = b, a%b
	}
	for a%2 < 1 {
		if b < 1 {
			b = a
		}
		a /= 2
	}

	Print(b, a)
}