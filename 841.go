package main
import . "fmt"
func main() {
	n := 0

	Scan(&n)
	a := n*(n-1) + 2
	if n < 1 {
		a = 1
	}

	Print(a)
}