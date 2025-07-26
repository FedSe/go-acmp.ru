package main
import . "fmt"
func main() {
	var m, x, y int
	Scan(&m, &m, &y, &x)

	n := y & 1 *(2*x-1-m)

	Print(n+m*y-x)
}