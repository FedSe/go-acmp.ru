package main
import (
	b "bufio"
	. "fmt"
	. "os"
	. "sort"
)
func main() {
	var (
		v       [100][100]int
		m, n, k int
		s       = b.NewReader(Stdin)
		S       = Fscan
		P       = Println
	)

	S(s, &m, &n, &k)
	o := make([][]int, m+1)
	for 0 < k {
		y := ""
		e := 0
		d := 0
		S(s, &y, &d)
		switch len(y) {
		case 3:
			S(s, &e)
			e--
			if v[e][d] < 1 {
				v[e][d] = 1
				o[d] = append(o[d], e+1)
			}
		case 7:
			y = ""
			for e < m {
				e++
				if v[d-1][e] > 0 {
					y += Sprintln(e)
				}
			}
			if len(y) < 1 {
				y = "-1"
			}
			P(y)
		case 10:
			g := o[d]
			if len(g) < 1 {
				P(-1)
			} else {
				Ints(g)
				for _, f := range g {
					P(f)
				}
			}
		}
		k--
	}
}