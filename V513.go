package main
import . "fmt"
func main() {
	n := 0
	Scan(&n)
	Print(1<<n - n - 1)
}