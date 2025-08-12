package main
import . "fmt"
func main() {
	var n, m, c int

	Scan(&n, &m)
	if m > n {
		n = 0
	}
	if m == n || m < 1 {
		c = 1
		n = -1
	}
	if m == 1 {
		c = n
		n = 0
	}

	d := 1
	m--
	for m*d < n {
		c += n - m*d
		d++
	}
	Print(c)
}