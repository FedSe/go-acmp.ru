package main
import . "fmt"
func main() {
	n := 0
	a := 1
	for a > 0 {
		Scan(&a)
		if a > n {
			n = a
		}
	}
	Print(n)
}