package main
import (
	b "bufio"
	. "fmt"
	. "math/rand"
	. "os"
	. "sort"
)
type I = int
func main() {
	var (
		X, i, l, j I
		z, g       [2e5 + 1]I
		q, h       [2e5 + 1]I
		D          = 1000000123
		s          = b.NewScanner(Stdin)
	)

	s.Buffer(make([]byte, 4096), 2e5)
	s.Scan()
	w := s.Text()
	n := len(w)
	w += w

	e := Intn(D-256) + 257
	e -= 1 - e&1

	q[0] = 1
	z[0] = 1
	for i < 2*n {
		i++
		q[i] = q[i-1] * e % D
		z[i] = z[i-1] * e
	}

	for j < 2*n {
		h[j+1] = (h[j] + I(w[j])*q[j]) % D
		g[j+1] = g[j] + I(w[j])*z[j]
		j++
	}

	F := func(o, x I) (I, I) {
		A := h[o+x] - h[o]
		B := g[o+x] - g[o]
		if A < 0 {
			A += D
		}
		t := 2*n - (o + x)
		A = A * q[t] % D
		B *= z[t]
		return A, B
	}

	o := make([]I, n)
	for l < n {
		o[l] = l
		l++
	}

	i = n - 1
	for i > 0 {
		j := Intn(i + 1)
		o[i], o[j] = o[j], o[i]
		i--
	}

	Slice(o, func(i, j I) bool {
		d, c := o[i], o[j]
		u := 0
		y := n + 1
		for y-u > 1 {
			k := (u + y) / 2
			a, m := F(d, k)
			A, B := F(c, k)
			if a == A && m == B {
				u = k
			} else {
				y = k
			}
		}
		if u < n {
			return w[d+u] < w[c+u]
		}
		return d < c
	})

	for i, v := range o {
		if v == 0 {
			X = i
			break
		}
	}

	Println(X + 1)
	W := make([]byte, n)
	for i, v := range o {
		W[i] = w[v+n-1]
	}

	Print(string(W))
}