package main
import . "fmt"
func main() {
	var a, b, c int
	Scan(&a, &b)
	d := b % a
	if d > 0 {
		c = a - d
	}
	Print(b/a, d, c)
}