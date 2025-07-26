package main
import . "fmt"
func main() {
	var n, a, b int
	Scan(&n, &a, &b)
	if a < b { a, b = b, a}
	a-=b
	b = n - a
	if a > b { a = b }
	Print(a-1)
}