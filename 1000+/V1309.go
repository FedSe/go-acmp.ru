package main
import . "fmt"
func main() {
	var a, b, c int
	d := 1

	Scan(&a, &b, &c)
	a %= 3
	b %= 3
	c %= 3
	if a == b && b != c {
		d = 3
	}
	if a == c && c != b {
		d = 2
	}

	Print(d)
}