package main
import . "fmt"
func main() {
	var a, b, c, d int
	s := "NO"
	Scan(&a, &b, &c, &d)
	if b*d == 8 || b+1 == d && a == c && b > 1 {
		s = "YES"
	}

	Print(s)
}