package main
import . "fmt"
func main() {
	var (
		n, i, a, b, c, e int
		d                [441][441]int
		x                = []int{-2, -2, -1, -1, 1, 1, 2, 2, -1, 1, -2, 2, -2, 2, -1, 1}
	)

	Scan(&n, &a, &b, &c, &e)
	for i <= n {
		v := 0
		for v <= n {
			d[i][v] = 999
			v++
		}
		i++
	}

	d[a][b] = 0
	q := [][2]int{{a, b}}

	for len(q) > 0 {
		a = q[0][0]
		b = q[0][1]
		q = q[1:]
		i = 0
		for i < 8 {
			u := a + x[i]
			v := b + x[i+8]
			if u*v > 0 && u <= n && v <= n && d[u][v] > d[a][b]+1 {
				d[u][v] = d[a][b] + 1
				q = append(q, [2]int{u, v})
			}
			i++
		}
	}

	Print(d[c][e])
}