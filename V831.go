package main
import . "fmt"
func main() {
	var a, b, s int
	Scan(&a, &b)
	n := -1

	for a <= b {
		c := a
		u := 0
		for c > 0 {
			u += c % 10
			c /= 10
		}

		d := 2
		f := 0 > 1
		for d*d <= a {
			if a % d < 1 {
				f = 1 > 0
			}
		d++
		}

		if !(f || a < 2) && u >= s {
			s = u
			n = a
		}
	a++
	}

	Print(n)

}