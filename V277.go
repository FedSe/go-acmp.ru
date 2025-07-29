package main
import . "fmt"
func main() {
	var a, b, c, d, i int
	P := Print
	Scan(&a, &b, &c)

	if a != 0 {
		P(a)
		d++
	}

	s := "x"
	for i < 2 {
		if b != 0 {
			if d*b > 0 {
				P("+")
			}
			if b*b != 1 {
				P(b)
			}
			if b == -1 {
				s = "-" + s
			}
			P(s)
			d++
		}
		b = c
		s = "y"
		i++
	}

	if d < 1 {
		P(0)
	}
}