package main
import . "fmt"
func main() {
	n := 0
	t := 0
	Scan(&n, &t)
	n++
	Print(t * 10 * n)
}