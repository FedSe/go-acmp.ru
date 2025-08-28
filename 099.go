package main
import . "fmt"
type A struct{ a, b, c int }
func main() {
	var (
		g          [50][50]string
		d          [50][50][50]int
		x, o       A
		h, m, n, i int
		S          = Scan
	)

	S(&h, &m, &n)
	h--
	for i < h+1 {
		l := 0
		for l < m {
			S(&g[i][l])
			j := 0
			for j < n {
				k := g[i][l][j]
				if i < 1 && k == 49 {
					o = A{0, l, j}
				}
				if i == h && k == 50 {
					x = A{h, l, j}
				}
				d[i][l][j]--
				j++
			}
			l++
		}
		if i < h {
			S()
		}
		i++
	}

	q := []A{o}
	d[o.a][o.b][o.c] = 0

	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		i = 0
		for i < 4 {
			r := u.b + i%2 - i/3*2
			v := u.c + 1 - i + i/3*2
			if r >= 0 && r < m && v >= 0 && v < n &&
				g[u.a][r][v] != 'o' && d[u.a][r][v] < 0 {
				d[u.a][r][v] = d[u.a][u.b][u.c] + 5
				q = append(q, A{u.a, r, v})
			}
			i++
		}
		if u.a < h {
			f := u.a + 1
			r := u.b
			v := u.c
			if g[f][r][v] != 'o' && d[f][r][v] < 0 {
				d[f][r][v] = d[u.a][u.b][u.c] + 5
				q = append(q, A{f, r, v})
			}
		}
	}

	Print(d[x.a][x.b][x.c])
}