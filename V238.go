package main
import . "fmt"
type A struct{ x, y int }
func main() {
	var (
		d, g                   [100][100]int
		D                      = []int{-1, 0, 1, 0, -1}
		t                      = map[A]A{}
		n, m, k, h, x, y, j, i int
		S                      = Scan
		r                      = "Impossible"
	)

	S(&n, &m, &k, &h)
	z := A{k - 1, h - 1}
	for i < n {
		j := 0
		for j < m {
			S(&g[i][j])
			d[i][j] = -1
			j++
		}
		i++
	}
	q := []A{z}
	d[z.x][z.y] = 1

	S(&i)
	for 0 < i {
		S(&k, &h, &x, &y)
		t[A{k - 1, h - 1}] = A{x - 1, y - 1}
		i--
	}

	S(&i)
	e := make([]A, i)
	for j < i {
		S(&k, &h)
		e[j] = A{k - 1, h - 1}
		j++
	}

	for len(q) > 0 {
		z = q[0]
		q = q[1:]
		v, o := t[z]
		k = d[z.x][z.y] + 1
		if o && d[v.x][v.y] < 0 {
			d[v.x][v.y] = k
			q = append(q, v)
		}
		i = 0
		for i < 4 {
			v := A{z.x + D[i], z.y + D[i+1]}
			if v.x >= 0 && v.x < n && v.y >= 0 && v.y < m && g[v.x][v.y] < 1 && d[v.x][v.y] < 0 {
				d[v.x][v.y] = k
				q = append(q, v)
			}
			i++
		}
	}

	x = -1
	for _, v := range e {
		i = d[v.x][v.y]
		if i > 0 && (x < 0 || i < x) {
			x = i
			r = Sprint(x)
		}
	}

	Print(r)
}