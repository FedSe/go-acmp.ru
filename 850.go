package main
import . "fmt"
func main() {
	a:=1
	b:=1
	Scan(&a, &b)
	if a > b { a, b = b, a }
	Print(b/2+b%2, a)
}