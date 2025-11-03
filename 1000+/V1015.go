package main
import . "fmt"
func main() {
	var i, a, x, n float64

	Scan(&n, &x, &i)
	x = (100 - x) / 100 * n
	for 0 < i {
		Scan(&a)
		n += a
		i--
	}

	Print(int(n), int((n-x)/n*100+.5+1e-9))
}