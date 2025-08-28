package main
import . "fmt"
func main() {
	var a, b, c int
	Scan(&a, &b, &c)
	Print((c*(a+b) + 7) / 8)
}