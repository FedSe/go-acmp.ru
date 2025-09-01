package main
import (
	. "fmt"
	. "strings"
)

var (
	q                [6e4][]int
	e, d             [6e4][3]int
	n, x, y, m, i, j int
	s                = ""
	a                = -1
	S                = Scan
	P                = Print
)

func F(w, z int) {
	for _, v := range q[w] {
		if v != z {
			F(v, w)
		}
	}
	i = 0
	for i < 3 {
		if e[w][i] < 1 {
			d[w][i] = -1e6
		} else {
			j = 1 - i&1 - i/2
			x = 1
			for _, v := range q[w] {
				n = -1e6
				y = 0
				for y < 3 {
					if y != i && d[v][y] > n {
						n = d[v][y]
					}
					y++
				}
				if n < 0 {
					x = 0
				}
				j += n
			}
			d[w][i] = -1e6
			if x > 0 {
				d[w][i] = j
			}
		}
		i++
	}
}

func main() {
	S(&n)
	for i < n-1 {
		S(&x, &y)
		q[x] = append(q[x], y)
		q[y] = append(q[y], x)
		i++
	}

	for j < n {
		j++
		S(&s)
		i = 0
		for k, v := range "IBV" {
			x = Index(s, string(v)) + 1
			e[j][k] = x
			i += x
		}
		if i < 1 {
			P(a)
			return
		}
	}

	F(1, 0)
	for m < 3 {
		x = d[1][m]
		if x > a {
			a = x
		}
		m++
	}

	P(a)
}