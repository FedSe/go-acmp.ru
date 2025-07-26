package main
import . "fmt"
func main() {
	var n, b, c int
	Scan(&n, &b, &c)

	for 1 < n {
		a := c - b
        	c = b
		b = a
	n--
	}

	Print(b, c)
}