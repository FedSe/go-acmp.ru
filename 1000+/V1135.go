package main
import . "fmt"
func main() {
	n := -1
	a := 1
	for a > 0 {
		Scan(&a)
		n += 1 - a&1
	}
	Print(n)
}