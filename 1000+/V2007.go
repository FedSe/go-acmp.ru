package main
import (
	b "bufio"
	. "fmt"
	. "os"
)
func main() {

	var (
		a          [2e5]int
		i, c, v, q int
		r          = b.NewScanner(Stdin)
	)

	Scan(&v)
	for r.Scan() {
		Sscan(r.Text(), &c)
		a[i] += c
		a[i/2] += 2*c - i%2*c
		if i%2 > 0 && i > 1 {
			a[i/2+1] += c
		}
		i++
	}

	c = 1
	for c <= v {
		if a[c] > q {
			q = a[c]
			i = c
		}
		c++
	}

	Print(q, i)
}