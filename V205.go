package main
import (
	. "fmt"
	. "strings"
)
func main() {
	var (
		q, w, e, h, m, c int
		s                = ""
		S                = Sscan
		P                = Printf
		f                = `%02d:%02d:%02d`
		k                = 3600
		u                = k * 24
	)

	Scanf(f+`
%s`, &q, &w, &e, &s)

	p := Split(s, ":")

	S(p[0], &c)
	if len(p) > 1 {
		m = c
		S(p[1], &c)
	}
	if len(p) > 2 {
		h = m * k
		m = c
		S(p[2], &c)
	}

	c += q*k + w*60 + e + h + m*60
	q = c % u
	c /= u

	P(f, q/k, (q%k)/60, q%60)
	if c > 0 {
		P("+%d days", c)
	}
}