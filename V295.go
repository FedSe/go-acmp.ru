package main
import (
	. "fmt"
	. "strings"
)
func main() {
	var (
		v [40]string
		s = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		o = "IMPOSSIBLE"
		n = s
		m = s
		i, a int
	)

	Scan(&n, &m)
	for i < 27 {
		w := ""
		for _ , x := range n {
			w += string(s[(Index(s, string(x))+i)%26])
		}
		if Index(w, m) > -1 {
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