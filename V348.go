package main
import . "fmt"
func main() {
	var a, b, c, d, e, f, g, h int
	s := "No"

	Scan(&a, &b, &c, &d, &e, &f, &g, &h)

	m := g - e
	n := h - f
	v := m * (b - f) - n * (a - e)
	w := m * (d - f) - n * (c - e)
	m = c - a
	n = d - b

	l := a*n-b*m
	t := v * w < 1 && (l+f*m-e*n)*(l-g*n+h*m) < 1

	if t {
		if b == f {
			b = a
			d = c
			f = e
			h = g
		}
		if b > d {
			b, d = d, b
		}
		if f > h {
			f, h = h, f
		}
		t = (f - d) * (h - b) < 1
	}

	if t {
		s = "Yes"
	}

	Print(s)
}