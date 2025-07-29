package main
import (
	. "fmt"
	. "strings"
)
func main() {
	var (
		v       [1e3]string
		a       = ""
		n, o, i int
		S = Scan
	)

	S(&n)
	for i < n {
		S(&v[i])
		i++
	}

	S(&a)
	for 0 < n {
		n--
		b := a
		i = 0
		for x := range v[n] {
			q := IndexByte(b, v[n][x])
			if q > -1 {
				b = b[:q] + "@" + b[q+1:]
			} else {
				break
			}
			i++
		}
		if i == len(v[n]) {
			o++
		}
	}

	Print(o)
}