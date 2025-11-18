package main
import . "fmt"
func main() {
	var n, m, a, b, c, u int

	Scan(&n, &m, &a, &b, &c)
	c *= n*m - 1
	if n > 1 && m > 1 {
		if m < n {
			n = m
		}
		if b < a {
			a = b
		}
		u = (2*n - 2) * a
	}

	Print(c + u)
}