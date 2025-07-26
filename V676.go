package main
import . "fmt"
func main() {
	a := 2
	n := 0
	Scan(&n)
	if n % 3 > 0 {
		a = 1
	}

	Print(a)
}