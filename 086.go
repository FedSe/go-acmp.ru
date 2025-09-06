package main
import . "fmt"
func main() {
	n := 0
	Scan(&n)
	n--
	Print(n*n - n)
}