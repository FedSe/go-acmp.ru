package main

import (
	. "fmt"
	. "sort"
)

type T []int

var (
	P = T{2, 3, 5, 7, 11, 13, 17, 19}
	D = 0
	q = 1 << 60
	o = q
)

func F(n int, f T) {
	if n == 1 {
		w := make(T, len(f))
		for i, d := range f {
			w[i] = d - 1
		}
		Sort(Reverse(IntSlice(w)))
		n := 1
		for i, e := range w {
			if i > 7 {
				goto A
			}
			p := P[i]
			j := 0
			for j < e {
				if n > 1e15/p {
					goto A
				}
				n *= p
				j++
			}
		}
		if n < q {
			q = n
		}
	A:
		return
	}

	v := 2
	x := len(f)
	if x > 0 {
		v = f[x-1]
	}
	for v <= n {
		if n%v < 1 {
			F(n/v, append(f, v))
		}
		v++
	}
}

func main() {
	Scan(&D)

	F(D, nil)
	if q == o {
		q = 0
	}

	Print(q)
}