package main
import (
	. "fmt"
	. "os"
	b "bufio"
)
func main() {
	var (
		a, p [131071]int
		n, m, q, i int
		s = b.NewScanner(Stdin)
	)

	Scan(&n, &m)
	s.Split(b.ScanWords)

	for i < n {
	i++
		p[i] = m + i - 1
	}
	n += m

	for s.Scan() {
		Sscan(s.Text(), &q)
		i = p[q]
		x := i
		u := 0
		for x > 0 {
			u += a[x]
		x &= x + 1
		x--
		}
		Print(i-m+1-u, " ")
		x = i
		for x < n {
			a[x]++
		x |= x + 1
		}
		m--
		p[q] = m
	}
}