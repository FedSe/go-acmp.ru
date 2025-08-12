package main
import . "fmt"
func main() {
	var a, b, s int
	n := -1

	Scan(&a, &b)
	for a <= b {
		c := a
		u := 0
		for c > 0 {
			u += c % 10
			c /= 10
		}
		d := 2
		c = 0
		for d*d <= a {
			if a%d < 1 {
				c = 1
			}
			d++
		}
		if c < 1 && a > 1 && u >= s {
			s = u
			n = a
		}
		a++
	}

	Print(n)
}