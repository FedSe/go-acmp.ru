package main
import . "fmt"
func main() {
	e := ""
	m := 0

	Scan(&m)
	n := m / 5 * 2
	m %= 5
	if n > 0 {
		e = "88"
	}
	if n&1 > 0 {
		e = "8"
	}

	for 0 < m {
		m--
		e += "1"
	}

	Print(e)
}