package main
import (
	. "fmt"
	. "sort"
)
func F(t int, o string) int {
	var (
		k, i, j int
		q = ""
		w = q
	)
	if o != "0" {
		v := o[0] == 45
		if v {
			o = o[1:len(o)]
		}
		l := len(o)
		p := make([]byte, l)
		for j < l {
			p[j] = o[j]
		j++
		}
		Slice(p, func(i, j int) bool {
			return p[i] < p[j]
		})
		for p[k] == 48 {
			k++
		}
		q = string(p[k])
		for i < k {
			q += "0"
		i++
		}
		for k < l-1 {
		k++
			q += string(p[k])
		}
		for l > 0 {
		l--
			w += string(p[l])
		}
		if t > 0 {
			w, q = q, w
		}
		if v {
			q = "-" + w
		}
		Sscan(q, &k)
	}
	return k
}

func main() {
	a := ""
	b := a
	Scan(&a, &b)
	Print(F(1, a) - F(0, b))
}