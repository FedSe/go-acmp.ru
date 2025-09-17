package main
import . "fmt"
func main() {
	var a, b, c int
	Scan(&a, &b, &c)
	t := a%2 + b%2 + c%2
	Print(a/2+b/2+c/2+t/2, ".", t%2*5)
}