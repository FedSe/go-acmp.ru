package main
import . "fmt"
func main() {
	s := 0
	b := 0
	a := 1
	Scan(&b)
	for a+b > 0 {
		s += b
		a = b
		Scan(&b)
	}
	Print(s)
}