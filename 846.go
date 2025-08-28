package main
import . "fmt"
type R struct{ x, y int }
func main() {
	var (
		g                      [20]string
		d                      [20][20]int
		D                      = []int{0, 0, -1, 1}
		N, M, Q, w, L, s, i, j int
		S                      = Scan
	)

	S(&N, &M)
	for i < N {
		S(&g[i])
		i++
	}
	S(&Q, &w, &L)

	Q--
	w--
	q := []R{{Q, w}}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		i = 0
		for i < 4 {
			z := p.x - D[3-i]
			v := p.y + D[i]
			if g[z][v] != 49 && d[z][v] < 1 {
				d[z][v] = d[p.x][p.y] + 1
				q = append(q, R{z, v})
			}
			i++
		}
	}

	for j < 4 {
		S(&N, &M, &i)
		N--
		M--
		if d[N][M] <= L && d[N][M] > 0 || N == Q && M == w {
			s += i
		}
		j++
	}

	Print(s)
}