package main
import . "fmt"
func main() {
	var a, b, c, d int
	s := "NO"
	Scan(&a, &b, &c, &d)
	if a == c && b > 1 && (b == 2 && d == 4 || b+1 == d) {
		s = "YES"
	}

	Print(s)
}