package main
import . "fmt"
func main() {
	a := 0
	Scan(&a)
	Print(a / (2 - a&1) * (1 - 1/a))
}