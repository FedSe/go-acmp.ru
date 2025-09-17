package main
import . "fmt"
func main() {
	n := 0
	Scan(&n)
	Print(n * (n + 2) * (2*n + 1) / 8)
}