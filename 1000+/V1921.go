package main
import . "fmt"
func main() {
	a := 0
	b := 0
	Scan(&a, &b)
	b = (a + 1) / (a - b)
	Print(a*(b-1), b)
}