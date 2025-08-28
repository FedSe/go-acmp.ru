package main
import (
	. "fmt"
	. "sort"
)
func main() {
	var k, l, A int

	for l < 2 {
		a := ""
		i := 0
		j := 0
		w := a
		Scan(&a)
		A = k
		k = 0
		if a != "0" {
			v := a[0] == 45
			if v {
				a = a[1:]
			}
			z := len(a)
			p := make([]byte, z)
			for j < z {
				p[j] = a[j]
				j++
			}
			Slice(p, func(i, j int) bool {
				return p[i] < p[j]
			})
			for p[k] == 48 {
				k++
			}
			a = string(p[k])
			for i < k {
				a += "0"
				i++
			}
			for k < z-1 {
				k++
				a += string(p[k])
			}
			for z > 0 {
				z--
				w += string(p[z])
			}
			if l < 1 {
				w, a = a, w
			}
			if v {
				a = "-" + w
			}
			Sscan(a, &k)
		}
		l++
	}

	Print(A - k)
}