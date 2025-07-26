package main
import (
	b "bufio"
	. "fmt"
	. "os"
)
func main() {
	var (
		m       [200001]int
		a, r    []int
		x, l, i int
		s       = b.NewReader(Stdin)
	)

	Fscan(s, &l)
	for l > 0 {
		Fscan(s, &x)
		m[x]++
		l--
	}
	for j, v := range m {
		if v > 0 {
			a = append(a, j)
		}
	}
	for len(a) > 0 {
		b := []int{}
		l = a[0]
		x = a[0]
		for _, c := range a {
			if c != a[0] {
				if x+1 != c {
					r = append(r, l, x)
					l = c
				}
			}
			x = c
			m[c]--
			if m[c] > 0 {
				b = append(b, c)
			}
		}
		r = append(r, l, x)
		a = b
	}
	x = len(r)

	Println(x / 2)
	for i < x {
		Println(r[i], r[i+1])
		i += 2
	}
}