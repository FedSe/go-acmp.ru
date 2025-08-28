package main
import . "fmt"
func main() {
	var a, b, c, d, e, f, m, i int
	Scan(&c, &m, &d, &e, &f)

	h := -1
	v := h
	d--
	for i <= 1e6 {
		i++
		p := m * i
		x := d / p
		y := d - x*p
		if x+1 == e && y/i+1 == f {
			x = (c - 1) / p
			y = c - x*p - 1
			z := y/i + 1
			if v < 0 {
				v = z
				b = 1
			}
			x++
			if h < 0 {
				h = x
				a = 1
			}
			if a > 0 && x != h {
				a++
			}
			if b > 0 && z != v {
				b++
			}
		}
	}

	if a != 1 {
		h = 0
		if a < 1 {
			h--
		}
	}

	if b != 1 {
		v = 0
		if b < 1 {
			v--
		}
	}

	Print(h, v)
}