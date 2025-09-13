package main
import . "fmt"
func main() {
	var a, b, c int
	Scan(&a, &b, &c)
	Print(c / (2*b + 3*a))
}