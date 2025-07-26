package main
import . "fmt"
func main() {
	a := 0
	n := 0
	s := ""

	Scan(&n)
	for 0 < n {
		Scan(&a)
		s = Sprint(a, " ") + s
	n--
	}
	
	Print(s)
}