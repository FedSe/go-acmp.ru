package main
import . "fmt"
func main() {
	var a, b, c int
	s:="YES"
	Scan(&a, &b, &c)

	c*=2
	if c > a || c > b {
		s = "NO"
	}

	Print(s)
}