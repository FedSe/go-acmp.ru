package main
import . "fmt"
func main() {
	var a, b, c int
	s := "NO"

	Scan(&a, &b, &c)
	if a < b+c && b < a+c && c < b+a {
		s = "YES"
	}

	Print(s)
}