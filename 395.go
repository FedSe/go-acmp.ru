package main
import . "fmt"
func main() {
	var n, r, c int

	Scan(&n, &r)
	for n <= r {
		p := n
		s := 1
		for p > 0 {
			s *= p % 10
			p /= 10
		}
		if s > 0 && n%s < 1 {
			c++
		}
		n++
	}

	Print(c)
}