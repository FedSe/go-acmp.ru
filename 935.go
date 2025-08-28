package main
import . "fmt"
func main() {
	var a, b, c, d int
	Scan(&a, &b, &c, &d)
	Print([]any{"YES", "NO"}[(a+b+c+d)&1])
}