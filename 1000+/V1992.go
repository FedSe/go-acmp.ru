package main
import . "fmt"
func main() {
	n := 0
	a := 2

	Scan(&n)
	if n < 7 {
		a = 1
	}
	if n > 5 {
		a++
	}

	Print(a)
}