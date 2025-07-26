package main
import (
	. "fmt"
	. "strings"
	. "os"
	b "bufio"
)
func main() {
	var (
		p [1000]string
		n, m, i, k int
		a = 1001
		v = -1
		d = v
		c = 1000
		u = "CIRCLE"
		o = u
		s = b.NewScanner(Stdin)
	)

	Scan(&n, &m, &o)
	for s.Scan() {
		p[i] = s.Text()
		k = Index(p[i], "*")
		if k >= 0 {
			if k < a {
				a = k
			}
			k = LastIndex(p[i], "*")
			if k > v {
				v = k
			}
			if i < c {
				c = i
			}
			if i > d {
				d = i
			}
		}
	i++
	}

	if a != 1001 {
		for
		k = a - 1
		k <= a+1
		{
			l := c - 1
			for l <= c+1 {
				y := v - 1
				for y <= v+1 {
					z := d - 1
					for z <= d+1 {
						if y-k == z-l && y-k >= 2 && k >= 0 && y <= m && l >= 0 && z <= n {
							a += 2
							for a < v-1 {
								x := c + 2
								for x < d-1 {
									if p[x][a] != 42 {
										goto A
									}
								x++
								}
							a++
							}
							u = "SQUARE"
							goto A
						}
					z++
					}
				y++
				}
			l++
			}
		k++
		}
	}
A:
	Print(u)
}