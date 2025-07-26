package main
import (
	. "fmt"
	. "os"
	b "bufio"
)
func main() {
	var (
		x, y, p, j [32]int
		n, m, a, k, c, d, e, f, g, h, i int
		r = b.NewReader(Stdin)
	)

	Scanln(&n)
	for i < n {
		Scanf("%d.%d.%d.%d\n", &x[i], &y[i], &p[i], &j[i])
	i++
	}

	Scanln(&m)
	for 0 < m {
		u := 0
		o, _ := r.ReadString('\n') 
		Sscanf(o, "%d.%d.%d.%d %d.%d.%d.%d", &a, &c, &e, &g, &k, &d, &f, &h)
		for
		i = 0
		i < n
		{
			if a & x[i] == k & x[i] && c & y[i] == d & y[i] && e & p[i] == f & p[i] && g & j[i] == h & j[i] {
				u++
			}
		i++
		}
		Println(u)
	m--
	}
}