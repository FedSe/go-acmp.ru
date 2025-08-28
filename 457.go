package main
import (
	. "fmt"
	. "sort"
)
var n, k, m int

func s() {
	var (
		t, d, j int
		a []int
		x = n
		i = 4
		e = 1000
	)

	for i > 0 {
		a = append(a, x % 10)
		x /= 10
	i--
	}
	Ints(a)

	i = 1
	for j < 4 {
		t += a[j] * e
		d += a[j] * i
		i *= 10
		e /= 10
	j++
	}

	m = d - t
}

func main() {
	Scan(&n)

	s()
	for n != m {
		n = m
		s()
		k++
	}

	Print(n, k)
}