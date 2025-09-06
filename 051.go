package main
import . "fmt"
func main() {
	n := 0
	a := 1
	s := ""
	
	Scan(&n, &s)
	for n > 1 {
		a *= n
		n -= len(s)
	}

	Print(a)
}