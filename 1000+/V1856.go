package main
import . "fmt"
func main() {
	n := 0
	f := 0

	Scan(&n, &f)
	n -= 2
	Print([]int{n * n * n, 6 * n * n, 12 * n, 8}[f])
}