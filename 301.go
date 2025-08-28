package main
import . "fmt"
func main() {
	var (
		n, m int
		y    = ""
		x    = y
		P    = Sprint
	)

	Scan(&m, &n)
	s := m
	m--
	for 0 < n {
		n--
		a := 9
		if s < 9 {
			a = s
		}
		s -= a
		x += P(a)

		a = 9
		if m < 9 {
			a = m
		}
		m -= a

		if n < 1 {
			a++
		}
		y = P(a) + y
	}

	Println(x, y)
}