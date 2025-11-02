package main
import . "fmt"
func main() {
	var a, b, c, d int
	Scan(&a, &b, &c, &d)
	Print(2*c-a, 2*d-b)
}