package main
import . "fmt"
type H struct{ a, b int }
func main() {
	var (
		w    = 131
		z    = 1313
		o    = 1000000007
		d    = 998244353
		k, g [2e5 + 1]int
		q, i int
		a    = ""
		b    = a
		h    = map[H]bool{}
	)

	k[0] = 1
	g[0] = 1
	for i < 2e5 {
		i++
		k[i] = k[i-1] * w % o
		g[i] = g[i-1] * z % d
	}

	Scan(&a, &b)
	A := len(a)
	B := len(b)
	if B <= A {
		var (
			x, m, y, c [2e5 + 2]int
			F          = func(l, r int) H {
				return H{(m[r] - m[l]*k[r-l]%o + o) % o,
					(x[r] - x[l]*g[r-l]%d + d) % d}
			}
			l, j, I, i int
		)
		for l < A {
			m[l+1] = (m[l]*w + int(a[l])) % o
			x[l+1] = (x[l]*z + int(a[l])) % d
			l++
		}
		s := b + b
		for j < len(s) {
			y[j+1] = (y[j]*w + int(s[j])) % o
			c[j+1] = (c[j]*z + int(s[j])) % d
			j++
		}
		for I < B {
			h[H{(y[I+B] - y[I]*k[B]%o + o) % o,
				(c[I+B] - c[I]*g[B]%d + d) % d}] = 1 > 0
			I++
		}
		for i <= A-B {
			if h[F(i, i+B)] {
				q++
			}
			i++
		}
	}

	Print(q)
}