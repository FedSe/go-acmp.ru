package main
import (
	. "fmt"
	. "sort"
)
func F(a []rune ) string {
	Slice(a, func(i, j int) bool { return a[i] < a[j] })
	return string (a)
}
func main() {
	s := "YES"
	c := s
	d := s
	Scan(&c, &d)

	i := len(c)
	if F([]rune(c)) != F([]rune(d)) || i != len(d) {
		s = "NO"
	}

	for 0 < i {
	i--
		if c[i] == d[i] {
			s = "NO"
		}
	}

	Print(s)
}