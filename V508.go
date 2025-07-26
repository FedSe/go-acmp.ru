package main
import . "fmt"
func main() {
	var n, s, o, r, f, h int
	p := ""
	Scan(&n)
	for 0 < n {
		Scan(&p, &o, &r, &f, &h)
		a := -1
		s -= r - f
		if p != "L" {
			a = s
		}
		s -= o - h
		Print(a, " ")
	n--
	}
}