package main
import . "fmt"
func main() {
	var n, m, d, l int
	Scan(&n, &m, &d, &l)
	Print((n+m)*d*l-n*m*d*d)
}