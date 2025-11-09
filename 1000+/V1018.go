package main
import (
	b "bufio"
	. "fmt"
	. "os"
)
type T struct{ y, x int }
func main() {
	var (
		w             [1e3][1e3]int
		n, m, y, x, i int
		s             = ""
		r             = b.NewReader(Stdin)
		P             = Print
	)

	Scan(&n, &m, &y, &x)
	for i < n {
		Fscan(r, &s)
		j := 0
		for j < m {
			w[i][j] = int(s[j]-48) - 1
			j++
		}
		i++
	}

	q := []T{{y - 1, x - 1}}
	for len(q) > 0 {
		y = q[0].y
		x = q[0].x
		q = q[1:]
		if x*y < 1 || x == m-1 || y == n-1 {
			P(w[y][x] + 1)
			return
		}
		i = 0
		for i < 6 {
			z := y + []int{-1, -1, 0, 0, 1, 1}[i]
			c := x + [][]int{
				{-1, 0, -1, 1, -1, 0},
				{0, 1, -1, 1, 0, 1}}[y&1][i]
			if w[z][c] == 0 {
				w[z][c] = w[y][x] + 1
				q = append(q, T{z, c})
			}
			i++
		}
	}

	P("No solution")
}