package main
import . "fmt"
func main() {
	var a, b, x int

	Scan(&a, &b)
	a *= b
	for x*x < a {
		x++
	}
	if x*x > a {
		x = 0
	}

	Print(x)
}