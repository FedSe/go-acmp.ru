package main
import . "fmt"
type S struct {
	a, b, c, d int
}
func main() {
	var (
		D       = []int{-2, -2, -1, -1, 1, 1, 2, 2, -1, 1, -2, 2, -2, 2, -1, 1}
		d       [8][8][8][8]int
		c, w, o int
		H       = ' '
		G       = H
		P       = Print
	)

	Scanf(`%c%d %c%d
`, &H, &w, &G, &o)

	h := int(H - 97)
	g := int(G - 97)
	w--
	o--
	q := []S{{w, h, o, g}}
	d[w][h][o][g] = 1
	for len(q) > 0 {
		c++
		z := len(q)
		s := 0
		for s < z {
			z := q[0]
			q = q[1:]
			i := 0
			for i < 8 {
				y := z.a + D[i]
				x := z.b + D[i+8]
				if y > 0 && y < 8 && x > 0 && x < 8 {
					j := 0
					for j < 8 {
						v := z.c + D[j]
						m := z.d + D[j+8]
						if v > 0 && v < 8 && m > 0 && m < 8 {
							if y == v && x == m {
								P(c)
								return
							}
							if d[y][x][v][m] < 1 {
								d[y][x][v][m] = 1
								q = append(q, S{y, x, v, m})
							}
						}
						j++
					}
				}
				i++
			}
			s++
		}
	}

	P(-1)
}