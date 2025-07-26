package main
import . "fmt"
func main() {
	var (
	n = 0
	v = .0
	c = v
	s = ""
	l = 30.
	r = 4000.
)
	Scan(&n, &v)

	for 1 < n {
		Scan(&c, &s)
		m := (v + c) / 2
		a := s[0] > 99

		if ((a && c > v) || (!a && c < v)) && r > m {
			r = m
		}
		if ((a && c < v) || (!a && c > v)) && m > l {
			l = m
		}

		v = c
	n--
	}

	Print(l, r)

}