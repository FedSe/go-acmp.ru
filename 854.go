package main
import . "fmt"
func main() {
	a := 0
	b := 0
	s := ""
	Scan(&a, &b, &s)

	if s[2] < 98 && a < b || s[0] < 98 || (len(s) > 5 && a > b) {
		a = b
	}

	Print(a)
}