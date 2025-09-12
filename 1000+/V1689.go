package main
import . "fmt"
func main() {
	var n, m, a, b int
	Scan(&n, &m, &a, &b)
	n--
	Print(n*a, b*(m-1+n))
}