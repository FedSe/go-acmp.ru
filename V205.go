package main
import (
	. "fmt"
	. "strconv"
	. "strings"
)
func main() {
	var (
		q, w, e, h, m int
		s             = ""
		A             = Atoi
		P             = Printf
		f             = `%02d:%02d:%02d`
		k             = 3600
		u             = 86400
	)

	Scanf(f+`
%s`, &q, &w, &e, &s)

	p := Split(s, ":")

	c, _ := A(p[0])
	if len(p) > 1 {
		m = c
		c, _ = A(p[1])
	}
	if len(p) > 2 {
		h = m
		m = c
		c, _ = A(p[2])
	}

	c += q*k + w*60 + e + h*k + m*60
	q = c % u
	c /= u

	P(f, q/k, (q%k)/60, q%60)
	if c > 0 {
		P("+%d days", c)
	}
}