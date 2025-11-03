package main
import (
	. "fmt"
	. "math/big"
)
type I = int
var m = map[I]int64{}

func O(a, b, c I) I {
	s := 0
	a %= c
	for b > 0 {
		if b&1 > 0 {
			s += a
			s %= c
		}
		a = a * 2 % c
		b /= 2
	}
	return s
}

func F(n I) {
	if n != 1 {
		d := n - 1
		k := 0
		l := 0
		a := 2 % n
		for _, p := range []I{2, 3, 5, 7, 11, 13} {
			if n == p {
				goto A
			}
			if n%p < 1 {
				goto B
			}
		}
		for d&1 < 1 {
			d /= 2
			k++
		}
		for l < 2 {
			r := 0
			x := 1
			t := 1
			w := d
			for w > 0 {
				if w&1 > 0 {
					x = O(x, a, n)
				}
				a = O(a, a, n)
				w /= 2
			}
			if x != 1 && x != n-1 {
				for r < k-1 {
					x = O(x, x, n)
					if x == n-1 {
						t = 0
					}
					r++
				}
				if t > 0 {
					goto B
				}
			}
			a = 325 % n
			l++
		}
	A:
		m[n]++
		return
	B:
		k = 2
		if n&1 > 0 {
			d = 1
			for {
				f := func(x I) I {
					return (O(x, x, n) + d) % n
				}
				x := 2
				a = 2
				k = 1
				for k == 1 {
					x = f(x)
					a = f(f(a))
					k = a - x
					if x > a {
						k = x - a
					}
					l := n
					for l > 0 {
						k, l = l, k%l
					}
					if k == n {
						break
					}
				}
				if k != n {
					break
				}
				d++
			}
		}
		F(k)
		F(n / k)
	}
}

func main() {
	var (
		n I
		N = NewInt
		r = N(1)
		q Int
	)

	Scan(&n)
	F(n)
	for p, a := range m {
		g := N(int64(p))
		o := N(1)
		if a != 1 {
			o = o.Exp(g, N(a-1), nil)
		}
		r.Mul(r, o.Mul(o, q.Add(g, q.Mul(N(a), q.Sub(g, N(1))))))
	}

	Print(r)
}