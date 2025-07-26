package main
import . "fmt"
func main() {
	n:=0
	Scan(&n)

	m := n/2 - 2
	d := 4

	if n%2 > 0 {
		m = n / 2
		d = 1
	}
	if n%4 < 1 {
		m = n/2 - 1
		d = 2
	}

	Print(m, m+d)
}