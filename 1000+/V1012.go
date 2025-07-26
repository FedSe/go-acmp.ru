package main
import . "fmt"
func main() {
	n := 0
	m := 0
	Scan(&n, &m)
	a := m-1
	Print(n*m/a, n/a)
}