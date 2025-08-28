package main
import (
	. "fmt"
	. "math"
)
func main() {
	r := 0.
	n := r

	Scan(&r, &n)
	c := 0
	m := int(r / n)
	i := -m - 1
	for i <= m {
		b := float64(i) * n
		a := Abs(b)
		b += n
		if b > a {
			a = b
		}
		a = r*r - a*a
		if a > 0 {
			s := int(Floor(Sqrt(a) / n))
			if s >= 0 {
				c += s + s
			}
		}
		i++
	}

	Print(c)
}