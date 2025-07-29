package main
import (
	b "bufio"
	. "fmt"
	. "os"
)
func main() {
	var (
		x, y, p, j                      [32]int
		n, m, a, k, c, d, e, f, g, h, i int
		r                               = b.NewReader(Stdin)
		q                               = "%d.%d.%d.%d "
		S                               = Scanln
	)

	S(&n)
	for i < n {
		Scanf(q, &x[i], &y[i], &p[i], &j[i])
		i++
	}

	S(&m)
	for 0 < m {
		u := 0
		o, _ := r.ReadString('\n')
		Sscanf(o, q+q, &a, &c, &e, &g, &k, &d, &f, &h)
		i = 0
		for i < n {
			if a&x[i] == k&x[i] && c&y[i] == d&y[i] && e&p[i] == f&p[i] && g&j[i] == h&j[i] {
				u++
			}
			i++
		}
		Println(u)
		m--
	}
}