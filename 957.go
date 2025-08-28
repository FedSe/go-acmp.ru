package main
import . "fmt"
var (
	d             [64][64]int
	g, r, v, i, t int
	P             = Print
)

func F(n, a, b, x, y int) {
	h := 3
	if n < 1 {
		return
	}
	n--
	e := 1 << n

	if a < x+e && b < y+e {
		h = 0
	}
	if a >= x+e && b < y+e {
		h = 1
	}
	if a < x+e && b >= y+e {
		h = 2
	}

	o := [][]int{
		{x + e - 1, y + e - 1},
		{x + e, y + e - 1},
		{x + e - 1, y + e},
		{x + e, y + e}}

	t++

	i := 0
	for i < 4 {
		if i != h {
			d[o[i][1]][o[i][0]] = t
		}
		i++
	}

	for I, V := range [][]int{{0, 0}, {e, 0}, {0, e}, {e, e}} {
		if h == I {
			F(n, a, b, x+V[0], y+V[1])
		} else {
			F(n, o[h][0], o[h][1], x+V[0], y+V[1])
		}
	}
}

func main() {
	Scan(&g, &r, &v)

	r--
	v--
	z := 1 << g
	d[v][r] = 0
	F(g, r, v, 0, 0)

	for i < z {
		g = 0
		for g < z {
			P(d[i][g], " ")
			g++
		}
		P(`
`)
		i++
	}
}