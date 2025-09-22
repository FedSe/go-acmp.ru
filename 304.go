package main
import (
	. "fmt"
	. "math/big"
)

var (
	N          = NewInt
	r          = N(1)
	f          = 24
	a, x, u, i int
)

func C(n, k int) *Int {
	c := N(1)
	i := k
	for i < n {
		i++
		c.Mul(c, N(int64(i))).Div(c, N(int64(i-k)))
	}
	return c
}

func main() {
	Scan(&u)
	for i < u {
		Scanf("%d:%d", &a, &x)
		if a < x {
			a, x = x, a
		}
		if i > 3 {
			f = 14
		}
		if a == f+1 {
			r.Mul(r, C(a+x-1, x))
		} else {
			r.Mul(r, C(2*f, f)).Mul(r, N(0).Lsh(N(1), uint(x-f)))
		}
		i++
	}

	Print(r)
}