package main
import . "fmt"
func main() {
	n := 0
	m := 0
	Scan(&n, &m)
	n--
	m--
	Print(n*m)
}