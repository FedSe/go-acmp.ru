package main
import (
	. "fmt"
	. "sort"
)
func main() {
	var n [5]int
	a := 0
	s := "Nothing"

	for a < 5 {
		Scan(&n[a])
	a++
	}
	Ints(n[:])

	a = n[0]

	var (
		b = n[1]
		c = n[2]
		d = n[3]
		e = n[4]

		h = a == b
		f = h && a == c
		k = d == e
		g = e == c && k
		w = b == c
		v = c == d
	)

	if h || w || v || k {
		s = "One Pair"
	}
	if h && v || (w && k) || (h && k) {
		s = "Two Pairs"
	}
	if f || (w && b == d) || (v && c == e) {
		s = "Three of a Kind"
	}
	if a+1 == b && b+1 == c && c+1 == d && d+1 == e {
		s = "Straight"
	}
	if f && k || (h && g) {
		s = "Full House"
	}
	if f && a == d || (e == b && g) {
		s = "Four of a Kind"
		if a == e {
			s = "Impossible"
		}
	}

	Print(s)
}