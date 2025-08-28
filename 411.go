package main
import (
	. "fmt"
	. "math"
)
func main() {
	var (
		a = 0.
		b = a
		c = a
		w = 1
		P = Println
	)

	Scan(&a, &b, &c)

	if a == 0 {
		if b == 0 {
			w--
			if c == 0 {
				w--
			}
		}

		P(w)
		if w > 0 {
			P(-c / b)
		}
		return
	}

	c = b*b - 4*a*c

	if c < 0 {
		w = 0
	}

	if c > 0 {
		w = 2
	}

	P(w)
	c = Sqrt(c)
	for 0 < w {
		P((-b - c) / (a * 2))
		c = -c
		w--
	}
}