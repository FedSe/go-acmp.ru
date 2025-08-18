package main
import (
	. "fmt"
	. "math/big"
)
func H(A, o [2][2]*Int) [2][2]*Int {
	var C [2][2]*Int
	i := 0
	for i < 2 {
		j := 0
		for j < 2 {
			C[i][j] = new(Int).Add(
				new(Int).Mul(A[i][0], o[0][j]),
				new(Int).Mul(A[i][1], o[1][j]))
			j++
		}
		i++
	}
	return C
}

func main() {
	var (
		N = NewInt
		n = 0
		q = N(1)
		g = N(0)
		m = [2][2]*Int{
			{q, q},
			{q, g}}
		r = [2][2]*Int{
			{q, g},
			{g, q}}
	)

	Scan(&n)
	if n > 0 {
		if n == 1 {
			r = m
		}
		o := m
		for n > 0 {
			if n&1 != 0 {
				r = H(r, o)
			}
			o = H(o, o)
			n >>= 1
		}
	}

	Print(q.Set(r[0][1]))
}