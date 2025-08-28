package main
import . "fmt"
func main() {
	var n, m, f, l int
	Scan(&n, &m, &f, &l)
	n = m + f + l - n
	Print(n, m-n, f-n)
}