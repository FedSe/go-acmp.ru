package main
import . "fmt"
func main() {
	n := 0

	Scan(&n)
	if n > 1 {
		a := n % 3
		if a == 1 {
			n -= 4
		}
		i := n / 3
		n = 1
		for i > 0 {
			n *= 3
			i--
		}
		if a > 0 {
			n *= 4 / a
		}
	}

	Print(n)
}