package main
import . "fmt"
func main() {
	n := 0
	c := 0
	a := 1
	Scan(&n)

	for a*2 < n {
		a *= 2
	}

	for n > 1 {
		if n > a {
			n -= a
			c++
		}
		a /= 2
	}

	Print(c % 3)
}