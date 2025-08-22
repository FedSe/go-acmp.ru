package main
import (
	. "fmt"
	. "math/big"
	. "strings"
)
func main() {
	var (
		D          []byte
		u, o, g, q string
		n          int64
		S          = map[any]int{}
		A, R, r, a Rat
		X, v, W, w Int
		p          = -1
		N          = NewInt
		I          = Index
		j          = N(1)
		i          = N(1)
		F          = N(10)
	)

	Scan(&q, &n)
	h := I(q, "(")
	if h > 0 {
		g = q[h+1 : I(q, ")")]
		q = q[:h]
	}
	h = I(q, ".")
	u = q
	if h != -1 {
		u = q[:h]
		o = q[h+1:]
	}

	A.SetString(u)
	X.SetString(o, 10)
	v.SetString(g, 10)
	j.Exp(F, N(int64(len(o))), nil)
	r = *A.Mul(A.Add(&A, R.Add(r.SetFrac(&X, j), a.SetFrac(&v, w.Mul(j, i.Sub(i.Exp(F, N(int64(len(g))), nil), N(1)))))), NewRat(n, 1))
	E := r.Denom()
	v.DivMod(r.Num(), E, &X)
	u = v.String()
	if X.Cmp(&W) > 0 {
		for X.Cmp(&W) > 0 {
			g = X.String()
			k, Y := S[g]
			if Y {
				p = k
				break
			}
			S[g] = len(D)
			X.Mul(&X, F)
			T := N(0)
			T.DivMod(&X, E, &X)
			D = append(D, byte(48+T.Int64()))
		}
		o = string(D)
		if p < 0 {
			u += "." + o
		} else {
			u += "." + o[:p] + "(" + o[p:] + ")"
		}
	}

	Print(u)
}