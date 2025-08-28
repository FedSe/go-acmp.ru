package main
import . "fmt"
func main() {
	var a, b, h, d, c int
	Scan(&h, &a, &b)

	for c < h-a {
		d++
		c += a - b
	}

	Print(d + 1)
}