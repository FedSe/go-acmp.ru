package main
import . "fmt"
func main() {
	var a, b, c int
	Scan(&a, &b, &c)
	b += c
	Print((a - 1 + b) / b)
}