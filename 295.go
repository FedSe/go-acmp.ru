package main
import (
	. "fmt"
	. "strings"
)
func main() {
	var (
		v    [40]string
		s    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		o    = "IMPOSSIBLE"
		n    = s
		m    = s
		i, a int
		I    = Index
	)

	Scan(&n, &m)
	for i < 27 {
		w := ""
		for _, x := range n {
			w += string(s[(I(s, string(x))+i)%26])
		}
		if I(w, m) > -1 {
			v[a] = w
			a++
		}
		i++
	}

	if a > 0 {
		o = v[a-1]
	}
	Print(o)
}