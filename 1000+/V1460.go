package main
import . "fmt"
func main() {
	var k, x, y int

	Scan(&k, &x, &y)
	a := y
	if y/x <= (y-1)/(x+k-1) {
		a = y/x*x + x
	}

	Print(a)
}