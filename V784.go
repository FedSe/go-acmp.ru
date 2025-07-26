package main
import . "fmt"
func main() {
	a := 0
	b := 0

	Scan(&a, &a, &b)
	for a != b {
		if a > b {
			a /= 2
		}
		if b > a {
			b /= 2
		}
	}

	Print(a)
}