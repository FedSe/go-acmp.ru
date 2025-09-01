package main
import (
	. "fmt"
	. "math"
)

var (
	A, B, C, L, p, o [8]float64
	n, l, i, j       int
	P                = Print
)

func F(i int) (a, b, c float64) {
	return A[i] / L[i], B[i] / L[i], C[i] / L[i]
}

func main() {
	Scan(&n)
	for l < n {
		Scan(&p[l], &o[l])
		l++
	}

	for i < n {
		l = (i + 1) % n
		k := p[l]
		u := o[l]
		A[i] = o[i] - u
		B[i] = k - p[i]
		C[i] = p[i]*u - k*o[i]
		L[i] = Hypot(A[i], B[i])
		i++
	}

	a, b, c := F(0)
	d, v, f := F(1)
	g, h, t := F(2)

	a -= d
	d -= g
	b -= v
	v -= h
	t -= f
	f -= c
	g = a*v - d*b
	b = (f*v - t*b) / g
	h = (a*t - d*f) / g

	for j < n {
		d = (A[j]*b + B[j]*h + C[j]) / L[j]
		if j < 1 {
			a = d
		}
		if Abs(d-a) > 1e-4 {
			P("NO")
			return
		}
		j++
	}

	P(`YES
`, b, h, a)
}