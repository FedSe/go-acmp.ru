package main
import . "fmt"
func main() {
	var a, b, c, d, e int
	s := "NO"

	Scan(&a, &b, &c, &d, &e)
	if a <= d && (b <= e || c <= e) ||
		b <= d && (a <= e || c <= e) ||
		c <= d && (a <= e || b <= e) {
		s = "YES"
	}

	Print(s)
}