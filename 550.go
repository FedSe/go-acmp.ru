package main
import . "fmt"
func main() {
	n := 0
	s := 3
	
	Scan(&n)
	if n%400 < 1 || n%4 < 1 && n%100 > 0 {
		s--
	}

	Printf("1%d/09/%04d", s, n)
}