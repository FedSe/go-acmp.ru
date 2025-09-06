package main
import . "fmt"
func main() {
	n := 0
	m := 0

	Scan(&n, &m)
	l := m
	if n*m < m {
		Print("Impossible")
		return
	}
	if n > m {
		l = n
	}
	if m > 0 {
		n += m - 1
	}

	Print(l, n)
}