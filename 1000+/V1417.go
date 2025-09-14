package main
import . "fmt"
func main() {
	var a, b, c, d int
	s := "No"

	Scan(&a, &b, &c, &d)
	if 2*a+b+c+d > 2 {
		s = "Yes"
	}

	Print(s)
}