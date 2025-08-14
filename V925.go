package main
import . "fmt"
func main() {
	var k, n, a, b, c int
	Scan(&k, &n, &a, &b, &c)

	if b < a {
		a, b = b, a
	}
	if c < a {
		a, c = c, a
	}

	b += a + c - 2*n
	if b < 0 {
		b = 0
	}
	if k > 1 {
		b = a
	}

	Print(b)
}