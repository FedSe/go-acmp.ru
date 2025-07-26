package main
import . "fmt"
func main() {
	n := .0
	p := n
	Scan(&n, &p)
	Print(100*n/(100 - p + n*p))
}