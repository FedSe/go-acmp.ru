package main

import (
	. "fmt"
	. "math/big"
)

func main() {
	n := 0
	Scan(&n)

	if n%2 > 0 {
		Print(0)
		return
	}

	n /= 2
	m := n * 9

	d := make([][]*Int, n+1)
	for i := range d {
		d[i] = make([]*Int, m+1)
		for j := range d[i] {
			d[i][j] = new(Int)
		}
	}

	d[0][0].SetInt64(1)

	i := 1
	for i <= n {
		s := 0
		for s <= m {
			j := 0
			for j <= 9 {
				if s >= j {
					d[i][s].Add(d[i][s], d[i-1][s-j])
				}
				j++
			}
			s++
		}
		i++
	}

	r := new(Int)
	i = 0
	for i <= m {
		a := d[n][i]
		t := new(Int).Mul(a, a)
		r.Add(r, t)
		i++
	}

	Print(r)
}
