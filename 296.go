package main
import . "fmt"
func main() {
	n := 0
	c := 0

	Scan(&n)
	for n%5 > 0 {
		n -= 3
		c++
	}

	Print(n/5, c)
}