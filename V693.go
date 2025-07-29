package main
import (
	. "fmt"
	. "sort"
	. "strings"
)
func main() {
	s := "No"
	a := s
	b := s
	L := ToLower

	Scan(&a, &b)

	c := []byte(L(a))
	d := []byte(L(b))

	Slice(c, func(i, j int) bool {
		return c[i] < c[j]
	})
	Slice(d, func(i, j int) bool {
		return d[i] < d[j]
	})

	if string(c) == string(d) {
		s = "Yes"
	}

	Print(s)
}