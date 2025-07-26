package main
import . "fmt"
func main() {
	n := 0
	Scan(&n)
	Print(n * (n + 1) * (n + 2) / 2)
}