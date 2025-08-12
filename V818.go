package main
import . "fmt"
func main() {
	n := 0
	b := 0

	Scan(&n)
	a := 1 - n
	for 0 < n {
		Scan(&b)
		a += b
		n--
	}

	Print(a)
}