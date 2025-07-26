package main
import . "fmt"
func main() {
	var (
		a, b, c, d, e, f int
		s = "Boxes are incomparable"
		w = "The first box is "
		v = "er than the second one"
	)
	Scan(&a, &b, &c, &d, &e, &f)

	if a < b {
		a, b = b, a
	}
	if b < c {
		b, c = c, b
	}
	if a < b {
		a, b = b, a
	}
	if d < e {
		d, e = e, d
	}
	if e < f {
		e, f = f, e
	}
	if d < e {
		d, e = e, d
	}

	if a >= d && b >= e && c >= f {
		s = w + "larg" + v
	}

	if a <= d && b <= e && c <= f {
		s = w + "small" + v
	}

	if a ^ d ^ b ^ e ^ c == f {
		s = "Boxes are equal"
	}

	Print(s)
}