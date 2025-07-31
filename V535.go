package main
import (
	. "fmt"
	. "sort"
)

func H(x, y int) int {
	v := 0
	p := 0
	for x+y > 0 {
		q := x%10 + y%10
		x /= 10
		y /= 10
		m := 1
		i := 0
		for i < p {
			m *= 10
			i++
		}
		v += m * q
		p++
		if q > 9 {
			p++
		}
	}
	return v
}

func main() {
	var (
		a, b, c int
		m       = map[int]int{}
		s       []int
		P       = Println
		w       = "YES"
	)

	Scan(&a, &b, &c)
	for _, p := range [][3]int{
		{a, b, c}, {a, c, b},
		{b, a, c}, {b, c, a},
		{c, a, b}, {c, b, a}} {
		m[H(H(p[0], p[1]), p[2])] = 1
	}

	for k := range m {
		s = append(s, k)
	}
	Ints(s)

	if len(m) < 2 {
		w = "NO"
	}

	P(w)
	for _, v := range s {
		P(v)
	}
}