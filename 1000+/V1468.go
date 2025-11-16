package main
import . "fmt"
func main() {
	x := 0
	S := Sprint

	Scan(&x)
	x--
	y := x + 2
	a := len(S(x))
	b := len(S(y))
	if a < b || a == b && x < y {
		y = x
	}

	Print(y)
}