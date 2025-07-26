package main
import . "fmt"
func main() {
	n := 0
	a := 3
	Scan(&n)

	for 0 < n-1 {
		a *= 2
		n--
	}

	Print(a)
}