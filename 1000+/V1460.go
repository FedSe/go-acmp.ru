package main
import . "fmt"
func main() {
	var k, x, y, a int
	Scan(&k, &x, &y)
	if y % x == 0 {
		a = y
	} else if y / x <= (y-1)/(x+k-1) {
		a = (y/x + 1) * x
	} else {
		a = y
	}
	Print(a)
}