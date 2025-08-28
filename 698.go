package main

import (
	. "fmt"
	. "strings"
)

var (
	N, M, c, l, v, i int
	h, k             [35]string
	z, y             [35]int
	w                [][]int
	o                = "6789TJQKA"
	s                = "NO"
	R                = o
	B                = IndexByte
	S                = Scan
)

func G(a, b byte) bool {
	return B(o, a) > B(o, b)
}

func F(i int) int {
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

func main() {
	S(&N, &M, &R)
	r := R[0]
	for l < N {
		S(&h[l])
		l++
	}

	for v < M {
		S(&k[v])
		v++
	}

	w = make([][]int, M)
	for i, a := range k[:M] {
		for j, d := range h[:N] {
			v = 0
			if d[1] == a[1] && G(d[0], a[0]) {
				v = 1
			}
			if d[1] == r {
				if a[1] != r || a[1] == r && G(d[0], a[0]) {
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