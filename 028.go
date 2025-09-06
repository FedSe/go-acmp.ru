package main
import . "fmt"
func main() {
	var a, c, b, x, y int

	Scan(&a, &c, &b, &x, &x, &y)
	c = 2*c - y
	if a == b {
		x = 2*a - x
		c = y
	}

	Print(x, c)
}