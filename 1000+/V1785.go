package main
import . "fmt"
func main() {
	n := 0
	Scan(&n)
	Print(n * (n*n + 6*n + 8) / 3)
}