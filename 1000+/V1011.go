package main
import (
	. "fmt"
	. "math"
)

type T = float64

var (
	x, y, z [3]T
	s       = ""
	i, l    int
	H       = Hypot
	P       = Print
)

func F(a, b, c, d T) (T, T) {
	e := H(a, b)
	f := H(c, d)
	return a/e + c/f, b/e + d/f
}

func G(a, b, c, d, e, f, g, h T) (T, T) {
	t := ((e-a)*h - (f-b)*g) / (c*h - d*g)
	return a + t*c, b + t*d
}

func main() {
	for i < 3 {
		Scan(&x[i], &y[i])
		i++
	}

	for l < 3 {
		z[l] = H(x[(l+1)%3]-x[l], y[(l+1)%3]-y[l])
		l++
	}

	Scan(&s)

	m := x[0]
	n := x[1]
	k := x[2]
	u := y[0]
	p := y[1]
	j := y[2]
	a, b := F(n-m, p-u, k-m, j-u)
	c, d := F(m-n, u-p, k-n, j-p)
	g, h := G(m, u, a, b, n, p, c, d)
	if s == "In" {
		a = z[0]
		b = z[1]
		c = z[2]
		P(g, h, Sqrt((a+b-c)*(b+c-a)*(c+a-b)/(a+b+c))/2)
		return
	}

	a = p - u
	b = m - n
	c = j - p
	d = n - k
	g, h = G((m+n)/2, (u+p)/2, a, b, (n+k)/2, (p+j)/2, c, d)

	P(g, h, H(g-m, h-u))
}