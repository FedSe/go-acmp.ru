package main
import . "fmt"
func main() {
	n := 0

	Scan(&n)
	if n & 1 < 1 {
		n = -2
	}

	Print(n & 1 * 2, n+2)
}