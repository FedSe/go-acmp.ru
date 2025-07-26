package main
import . "fmt"
func main() {
	var k, n, m, d int
	Scan(&k, &n, &m, &d)

	a := -1
	e := n*m*k - n*m - n*k - k*m

	if e > 0 {
		d = (d * n * m * k) / e
		if !(d%k != 0 || d%n != 0 || d%m != 0) {
			a = d
		}
	}

	Print(a)

}