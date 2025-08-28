package main
import . "fmt"
func main() {
	n := 0
	Scan(&n)

	m := n / 2
	d := 2

	if n%2 > 0 {
		m += 2
		d -= 3
	}
	if n%4 < 1 {
		m++
		d -= 2
	}

	Print(m-2, m+d)
}