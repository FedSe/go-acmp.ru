package main
import . "fmt"
func main() {
	m := 0
	n := 0
	Scan(&m, &n)
	m--
	Print(m*n - m)
}