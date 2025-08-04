package main
import (
	. "fmt"
	. "strings"
)
func main() {
	var (
		N, M, c, l, v, i int
		h, k             [35]string
		z, y             [35]int
		o                = "6789TJQKA"
		s                = "NO"
		R                = o
		F                func(int) int
	)

	Scan(&N, &M, &R)
	r := R[0]
	for l < N {
		Scan(&h[l])
		l++
	}

	for v < M {
		Scan(&k[v])
		v++
	}

	g := func(a, b byte) bool {
		return IndexByte(o, a) > IndexByte(o, b)
	}

	w := make([][]int, M)
	for i, a := range k[:M] {
		for j, d := range h[:N] {
			v = 0
			if d[1] == a[1] && g(d[0], a[0]) {
				v = 1
			}
			if d[1] == r {
				if a[1] != r || a[1] == r && g(d[0], a[0]) {
					v = 1
				}
			}
			if v > 0 {
				w[i] = append(w[i], j)
			}
		}
	}

	for i := range z {
		z[i] = -1
	}

	F = func(i int) int {
		for _, j := range w[i] {
			if y[j] < 1 {
				y[j] = 1
				if z[j] < 0 || F(z[j]) > 0 {
					z[j] = i
					return 1
				}
			}
		}
		return 0
	}

	for i < M {
		for j := range y {
			y[j] = 0
		}
		if F(i) > 0 {
			c++
		}
		i++
	}

	if c == M {
		s = "YES"
	}
	Print(s)
}