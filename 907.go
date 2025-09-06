package main
import . "fmt"
func main() {
	var a, b, c int
	s := "YES"

	Scan(&a, &b, &c)
	if c*2 > a || c*2 > b {
		s = "NO"
	}

	Print(s)
}