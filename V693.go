package main
import (
	. "fmt"
	. "sort"
	. "strings"
)
func main() {
	s:="No"
	a:=s
	b:=s
	Scan(&a, &b)

	c := []byte(ToLower(a))
	d := []byte(ToLower(b))

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