package main
import . "fmt"
func main() {
	a := 0
	b := 0
	Scan(&a, &b)

	if a%b < 1 || b%a < 1 {
		a = 1
	}

	Print(a)
}