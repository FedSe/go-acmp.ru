package main
import . "fmt"
func main() {
	a := 0
	b := 0

	Scan(&a, &b)
	for b > 0 {
		a, b = b, a%b
	}

	Print(a)
}