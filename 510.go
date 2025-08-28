package main
import . "fmt"
func main() {
	b := 0
	a := 1
	n := 0

	Scan(&n)
	if n&1 < 1 {
		b = 3
		for 2 < n {
			b, a = 4*b-a, b
			n -= 2
		}
	}

	Print(b)
}