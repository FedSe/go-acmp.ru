package main
import (
	. "fmt"
	. "strings"
)
func main() {
	var (
		v []int
		n, r, l, j int
		a = ""
		b = a
	)

	Scan(&a, &b)
	for a != b {
		n++
		for a[l] == b[l] {
			l++
		}
		r = Index(a[l:], string(b[l])) + l
		v = append(v, l, r)

		i := 0
		for i <= (r-l)/2 {
			e := []byte(a)
			e[l+i], e[r-i] = e[r-i], e[l+i]
			a = string(e)
		i++
		}
	}

	Println(n)
	n = len(v)
	if n > 0 {
		for j < n-1 {
			Println(v[j]+1, v[j+1]+1)
		j += 2
		}
	}
}