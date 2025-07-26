package main
import . "fmt"
func main() {
	m := 0
	n := 0
	Scan(&m, &n)
	y := ""
	x := y
	s := m

	m--
	for 0 < n {
	n--
		a := 9
		if s < 9 { a = s }
		s -= a
		x += Sprint(a)
		
		a = 9
		if m < 9 { a = m }
		m -= a

		if n < 1 { a++ }
		y = Sprint(a) + y
	}

	Print(x, " ", y)
}