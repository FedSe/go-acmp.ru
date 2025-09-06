package main
import . "fmt"
func main() {
	var a, b, c, d int

	Scan(&a, &b, &c, &d)
	a -= c
	b -= d
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	for a > 0 {
		a, b = b%a, a
	}

	Print(b + 1)
}