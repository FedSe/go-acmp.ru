package main
import (
	. "fmt"
	. "math/big"
)
func main() {
	var (
		a, b, c, i    int64
		s             = ""
		x, p, S, P, Q Int
		t             = NewInt(1)
	)

	Scan(&a, &b, &c, &s)
	b = a - b
	a -= c
	c = b + a - 1
	for i < a {
		e := c - (b + i)
		if e >= 0 {
			x.Add(&x, p.Lsh(p.Binomial(b-1+i, i), uint(e)))
		}
		i++
	}
	S.SetString(s, 10)
	Print(&P, Q.Sub(&S, P.Div(P.Mul(&S, &x), t.Lsh(t, uint(c)))))
}