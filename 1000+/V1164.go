package main
import . "fmt"
func main() {
	var (
		n, m, a int
		M       = 9 << 39
		x       = -M
		y       = M
		P       = Print
	)

	Scan(&n, &m)
	for 0 < n {
		n--
		Scan(&a)
		if a%m == 0 {
			if a > x {
				x = a
			}
			if a < y {
				y = a
			}
		}
	}

	if y == M {
		P("NO")
	} else {
		P(y, x)
	}
}