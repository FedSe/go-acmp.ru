package main
import . "fmt"
func main() {
	var a, b, c int

	Scan(&a, &b, &c)
	a += b - c
	if a < 0 {
		Print("Impossible")
		return
	}

	Print(a)
}