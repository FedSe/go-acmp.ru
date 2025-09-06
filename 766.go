package main
import . "fmt"
func main() {
	var a, b, c int
	s := "YES"

	Scan(&a, &b, &c)
	if a*b < c {
		s = "NO"
	}

	Print(s)
}