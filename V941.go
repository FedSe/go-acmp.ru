package main
import . "fmt"
func main() {
	var a, b, i, m, c, d uint
	Scan(&a, &b)

	for i < 2 {
		c = 0
		d = 1
		for a > 0 {
			c += a % 10 * d
			d *= 3
			a /= 10
		}
		m += c
		a = b
	i++
	}

	Print(m)
}