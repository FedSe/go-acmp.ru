package main
import (
	. "fmt"
	. "strings"
)
var n, m, v, h, o, i, j, u, k int
func F(q[]int, g [][]byte, z *int) [][]byte {
	for len(q) > 0 {
		h = q[0]
		o = q[1]
		j = q[2]
		q = q[3:]
		if h == n-2 && o == m-2 {
			*z = j
		}
		j++
		if g[h-1][o] == 46 {
			g[h-1][o] = 42
			q = append(q, h-1, o, j)
		}
		if g[h+1][o] == 46 {
			g[h+1][o] = 42
			q = append(q, h+1, o, j)
		}
		if g[h][o-1] == 46 {
			g[h][o-1] = 42
			q = append(q, h, o-1, j)
		}
		o++
		if g[h][o] == 46 {
			g[h][o] = 42
			q = append(q, h, o, j)
		}
	}
	return g
}
func main() {
	s := ""
	Scan(&n, &m)
	g := make([][]byte, n)
	for i < n {
		Scan(&s)
		g[i] = []byte(s)
		j = Index(s, "T")
		if j > 0 {
			h = i
			o = j
		}
	i++
	}

	q := []int{h, o, 0}
	g[h][o] = 42

	g = F(q, g, &v)

	if g[n-2][m-2] == 46 {
		k = 1
	}

	i = 1
	for i < n-1 {
		j = 1
		for j < m-1 {
			if g[i][j] != 35 {
				g[i][j] = 46
			}
		j++
		}
	i++
	}

	q = []int{1, 1, 0}
	g[1][1] = 42

	g = F(q, g, &u)

	s = "No"
	Println(u)
	if u < v || k > 0 {
		s = "Yes"
	}
	Print(s)
}