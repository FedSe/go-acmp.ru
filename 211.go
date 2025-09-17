package main
import . "fmt"
func main() {
	var (
		g                   [75]string
		w, h, i, a, b, c, d int
		S                   = Scan
	)

	S(&w, &h)
	for i < h {
		S(&g[i])
		i++
	}

	for {
		S(&a, &b, &c, &d)
		if a+b+c+d < 1 {
			break
		}
		var p, s [100][100]int
		a += 9
		b += 9
		c += 9
		d += 9
		q := [][]int{{a, b}}
		s[b][a] = 1
		for len(q) > 0 {
			x := q[0][0]
			y := q[0][1]
			q = q[1:]
			if x == c && y == d {
				goto A
			}
			i = 0
			for i < 4 {
				X := x + (1-i&1*2)*(1-i/2)
				Y := y + i - i&1*i - i/2
				k := 1 < 0
				if X > 9 && X < 10+w && Y > 9 && Y < 10+h {
					k = g[Y-10][X-10] == 88
				}
				if X >= 0 && X < 100 && Y >= 0 && Y < 100 &&
					s[Y][X] < 1 && (X == c && Y == d || !k) {
					s[Y][X] = 1
					p[Y][X] = p[y][x] + 1
					q = append(q, []int{X, Y})
				}
				i++
			}
		}
	A:
		Println(p[d][c])
	}
}