package main
import (
	. "fmt"
	. "strings"
)
func main() {
	var (
		e, q [8]int
		p    []string
		s    = ""
		o    = 7
		i    = 1
		j    = 0
	)

	Scan(&s)
	for i < 8 {
		e[i] = -1
		i++
	}

	i = 2
	for i < 8 {
		h := 2
		for h < 5 && i-h >= 0 {
			z := i - h
			t := s[z:i]
			n := len(t)
			x := 0
			if n == 2 {
				if t[0] == t[1] {
					x = 2
				}
			}
			if n == 3 {
				a, b, c := t[0], t[1], t[2]
				if a == b || b == c || a == c {
					x = 2
				}
				if a == b && b == c {
					x = 3
				}
			}
			if n == 4 {
				a, b, c, d := t[0], t[1], t[2], t[3]
				if a == c && a != b && a != d && b != d ||
					b == d && a != b && b != c && a != c {
					x = 2
				}
				if a == b && c == d && a != c ||
					a == b && b == c && a != d ||
					b == c && c == d && a != b {
					x = 3
				}
				if a == d && b == c && a != b {
					x = 4
				}
				if a == c && b == d && a != b {
					x = 3
				}
				if a == b && b == c && c == d {
					x = 5
				}
			}
			if e[z]+x > e[i] {
				e[i] = e[z] + x
				q[i] = h
			}
			h++
		}
		i++
	}

	for o > 0 {
		l := q[o]
		p = append(p, s[o-l:o])
		o -= l
	}

	o = len(p)
	for j < o/2 {
		i = o - 1 - j
		p[j], p[i] = p[i], p[j]
		j++
	}

	Println(Join(p, "-"), e[7])
}