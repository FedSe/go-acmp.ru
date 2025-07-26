package main
import . "fmt"
func main() {
	var x, y, a int
	Scan(&x, &y)

	if x+y > 2 {
		a = 1
		if x == y {
			a = 2
		}
	}

	Print(a)
}