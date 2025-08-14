package main
import . "fmt"
func main() {
	n := 0
	a := 0
	b := 1

	Scan(&n)
	for n > 1 {
		if n&1 < 1 {
			b += a
		} else {
			a += b
		}
		n >>= 1
	}

	Print(a + b)
}