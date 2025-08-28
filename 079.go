package main
import . "fmt"
func main() {
	a := 0
	b := 0

	Scan(&a, &b)
	a %= 10
	c := a
	for b > 1 {
		a = a * c % 10
		b--
	}

	Print(a)
}