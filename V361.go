package main
import (
	. "fmt"
	. "sort"
)
func main() {
	s := ""
	q := 0

	Scan(&s)
	n := len(s)
	L := n - 1
	for L > 0 {
		e := map[any]int{}
		i := 0
		for i <= n-L {
			r := []rune(s[i : i+L])
			Slice(r, func(i, j int) bool {
				return r[i] < r[j]
			})
			h := string(r)
			if e[h] > 0 {
				q = L
				goto A
			}
			e[h] = 1
			i++
		}
		L--
	}
A:
	Print(q)
}