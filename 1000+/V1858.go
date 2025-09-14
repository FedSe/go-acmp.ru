package main
import . "fmt"
func main() {
	n := 0
	m := 0

	Scan(&n, &m)
	n--
	d := (m - 1) / n
	Print(d*m - n*d*(d+1)/2)
}