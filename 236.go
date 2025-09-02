package main
import (
	. "fmt"
	. "strings"
)
func main() {
	var (
		s          = ""
		x, i, t, y int
		q          = 1
		S          = Sscan
	)

	Scan(&s, &x)
	s += "+"
	if s[0] == 45 {
		q = -1
		i = 1
	}
	for i < len(s) {
		o := i
		for s[o] != 43 && s[o] != 45 {
			o++
		}
		h := s[i:o]
		z := s[o]
		i = o + 1
		o = 0
		S(h, &y)
		u := Split(h, "x")
		if len(u) > 1 {
			h = u[0]
			o = len(h)
			if o > 0 {
				h = h[:o-1]
			}
			S(h, &y)
			if h == "" {
				y = 1
			}
			if h == "-" {
				y = -1
			}
			o = 1
			h = u[1]
			if len(h) > 0 && h[0] == 94 {
				S(h[1:], &o)
			}
		}
		b := 1
		p := 0
		for p < o {
			b *= x
			p++
		}
		t += q * y * b
		q = -1
		if z == 43 {
			q = 1
		}
	}

	Print(t)
}