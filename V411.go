package main
import (
	. "fmt"
	. "math"
)
func main() {
	a:=0.
	b:=a
	c:=a
	w := 1
	Scan(&a, &b, &c)

	if a == 0 {
		if b == 0 {
			w--
			if c == 0 {
				w--
			}
		}

		Println(w)
		if w > 0 {
			Print(-c / b)
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

	Println(w)
	c = Sqrt(c)
	for 0 < w {
		Println((-b - c) / (a * 2))
		c = -c
	w--
	}

}