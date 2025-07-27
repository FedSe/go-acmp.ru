package main
import (
	b "bufio"
	. "fmt"
	. "os"
)
func main() {
	var (
		a, p       [131071]int
		n, m, q, i int
		s          = b.NewReader(Stdin)
	)

	Fscan(s, &n, &m)
	for i < n {
		i++
		p[i] = m + i - 1
	}
	n += m

	for m > 0 {
		Fscan(s, &q)
		i = p[q]
		x := i
		u := 0
		for x > 0 {
			u += a[x]
			x &= x + 1
			x--
		}
		Println(i - m + 1 - u)
		x = i
		for x < n {
			a[x]++
			x |= x + 1
		}
		m--
		p[q] = m
	}
}