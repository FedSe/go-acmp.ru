package main
import . "fmt"
func main() {
	n := 0

	Scan(&n)

	Print(n - n/3 - n/5 - n/2 - n/30 + n/15 + n/6 + n/10)
}