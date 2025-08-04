package main

import (
	. "fmt"
	. "sort"
)

func main() {
	var (
		a          [100]int
		L, N, i, c int
	)

	Scan(&L, &N)
	for i < N {
		Scan(&a[i])
		i++
	}

	Ints(a[:N])
	t := a[0] + L
	i = 1
	for i < N {
		if a[i]-L > t {
			c++
			t = a[i] + L
		}
		i++
	}
	Print(c + 1)
}
