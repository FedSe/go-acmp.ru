package main
import (
	b "bufio"
	. "fmt"
	. "os"
)
func main() {
	var (
		m [200001]int
		a, r []int
		x, i int
		s = b.NewScanner(Stdin)
	)
	s.Split(b.ScanWords)

	Scan(&x)
	for s.Scan() {
		Sscan(s.Text(), &x)
		m[x]++
	}
	for j, v := range m {
		if v > 0 {
			a = append(a, j)
		}
	}
	for len(a) > 0 {
		b := []int{}
		l := a[0]
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

	Println(x/2)
	for i < x {
		Println(r[i], r[i+1])
	i += 2
	}
}