package main
import . "fmt"
func main() {
	var n, p, s int
	Scan(&n, &p)

	if p < 2 {
		s = n
	} else {
		for p%2 > 0 {
			s += n / 2
			n -= n / 2
			p++
			p /= 2
		}
	}

	Print(s + p/2)
}