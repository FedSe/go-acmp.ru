package main
import (
	b "bufio"
	. "fmt"
	. "os"
)

var e, s [4e5]int

func F(v, f, z int) {
	if f == z {
		e[v] = s[f]
		return
	}
	d := f + (z-f)/2
	F(2*v+1, f, d)
	F(2*v+2, d+1, z)
	e[v] = e[2*v+1] + e[2*v+2]
}

func Q(v, f, z, l, r int) int {
	if r < f || l > z {
		return 0
	}
	if l <= f && z <= r {
		return e[v]
	}
	d := f + (z-f)/2
	return Q(2*v+1, f, d, l, r) + Q(2*v+2, d+1, z, l, r)
}

func P(v, f, z, p, a int) {
	if f == z {
		e[v] = a
		return
	}
	d := f + (z-f)/2
	if p > d {
		P(2*v+2, d+1, z, p, a)
	} else {
		P(2*v+1, f, d, p, a)
	}
	e[v] = e[2*v+1] + e[2*v+2]
}

func main() {
	var (
		n, m, i, l, r int
		q             = ""
		R             = b.NewReader(Stdin)
		S             = Fscan
	)

	S(R, &n)
	for i < n {
		S(R, &s[i])
		i++
	}

	F(0, 0, n-1)
	S(R, &m)
	for m > 0 {
		m--
		S(R, &q, &l, &r)
		if q == "s" {
			Println(Q(0, 0, n-1, l-1, r-1))
		} else {
			P(0, 0, n-1, l-1, r)
		}
	}
}