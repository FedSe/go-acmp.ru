package main
import . "fmt"
func main() {
	a := 0
	b := 0

	Scan(&a, &b)
	n := a - b
	if n < 0 {
		n = -n
	}
	n *= 2
	if n >= a && n >= b {
		Print(n)
		return
	}

	Print("No")
}