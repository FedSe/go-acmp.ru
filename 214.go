package main
import . "fmt"
func main() {
	a := 0
	b := 0

	Scan(&a, &b)
	a |= b
	Print(2 - a&1)
}