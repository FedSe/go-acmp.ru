package main
import (
	. "fmt"
	. "math/big"
)
func main() {
	var (
		x, y *Int
		b    = 1
		s    = ""
		c    = func(q **Int) {
			Scan(&s)
			i := 0
			for i < len(s) {
				h := s[i:] + s[:i]
				if h[0] > 48 || len(h) < 2 {
					u := new(Int)
					u.SetString(h, 10)
					if *q == nil || u.Cmp(*q)*b > 0 {
						*q = u
					}
				}
				i++
			}
			b = -b
		}
	)

	c(&x)
	c(&y)

	Print(y.Sub(x, y))
}