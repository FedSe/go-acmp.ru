package main
import (
	b "bufio"
	. "fmt"
	. "os"
	. "strings"
)
type T [4]int
func main() {
	var (
		I                         = 9 << 18
		R, c, v, h, m, u, t, j, k int
		x                         = "RGBYX"
		g                         [50]string
		q                         T
		o                         = I
		r                         = b.NewReader(Stdin)
		P                         = Print
		S                         = Fscan
	)

	S(r, &R, &c)
	for j < 4 {
		S(r, &q[j])
		j++
	}

	for k < R {
		S(r, &g[k])
		j = 0
		for j < c {
			l := g[k][j]
			if l == 83 {
				v = k
				h = j
			}
			if l == 69 {
				m = k
				u = j
			}
			j++
		}
		k++
	}

	for t < 16 {
		var (
			y [50][50]int
			p = []T{{v, h}}
			z = map[any]int{}
			U = 0
			i = 0
		)
		k = 0
		for i < 4 {
			if 1<<i&t > 0 {
				k += q[i]
				z[x[i]] = 1
			}
			i++
		}
		for len(p) > 0 {
			r := p[0][0]
			j = p[0][1]
			p = p[1:]
			if r == m && j == u {
				U = 1
				break
			}
			i = 0
			for i < 4 {
				a := r + (1-i&1*2)*(1-i/2)
				d := j + i - i&1*i - i/2
				if a >= 0 && a < R && d >= 0 && d < c && y[a][d] < 1 {
					e := g[a][d]
					if !Contains(x, string(e)) || z[e] > 0 {
						y[a][d] = 1
						p = append(p, T{a, d})
					}
				}
				i++
			}
		}
		if U > 0 && k < o {
			o = k
		}
		t++
	}

	if o == I {
		P("Sleep")
		return
	}

	P(o)
}