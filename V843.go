package main
import . "fmt"
func main() {
	var n, m, k, f, l int
	Scan(&n, &m, &f, &l)
	k = m+f+l-n

	Print(k, m-k, f-k)
}