package main
import . "fmt"

func F(a, b int) int {
	if a < b {
		b = a
	}
	return b
}

func H(a, b int) int {
	if a > b {
		b = a
	}
	return b
}

func main() {
	var a, b, c, d, e, f, i, z, x, y, n, k, l, w, r, t, u, o, p int

	for i < 2 {
		w = z
		r = x
		t = y
		u = n
		o = k
		p = l
		Scan(&a, &b, &c, &d, &e, &f)
		z = F(a, d)
		x = F(b, e)
		y = F(c, f)
		n = H(a, d)
		k = H(b, e)
		l = H(c, f)
		i++
	}

	i = (u-w)*(o-r)*(p-t) + (n-z)*(k-x)*(l-y)

	a = H(w, z)
	b = F(u, n)
	c = H(r, x)
	d = F(o, k)
	e = H(t, y)
	f = F(p, l)

	if a < b && c < d && e < f {
		i -= (b - a) * (d - c) * (f - e)
	}

	Print(i)
}