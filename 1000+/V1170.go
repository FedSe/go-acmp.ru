package main
import . "fmt"
type I = int
type P struct{ x, y I }
type B []byte
type T [1e3][1e3]I
var (
	n, m, r I
	M       I = 9e9
	s       []B
	w       T
)

func F(x, y I) {
	w = T{}
	q := []P{{x, y}}
	w[x][y] = 1
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		x = p.x
		y = p.y
		i := 0
		for i < 4 {
			X := x - 1 + i - i/3*2
			Y := y + i - i/2*2 - i/3*2
			if X >= 0 && X < n && Y >= 0 && Y < m &&
				s[X][Y] == 46 && w[X][Y] == 0 {
				w[X][Y] = w[x][y] + 1
				q = append(q, P{X, Y})
			}
			i++
		}
	}
	r = M
	i := 0
	for i < n {
		j := 0
		for j < m {
			h := w[i][j]
			if (i*j < 1 || i == n-1 || j == m-1) && h > 0 && h < r {
				r = h
			}
			j++
		}
		i++
	}
}

func main() {
	var (
		S                = "No way "
		h                = S
		i, l, e, g, d, v I
		P                = Print
	)

	Scan(&n, &m)
	s = make([]B, n)
	for i < n {
		Scan(&h)
		s[i] = B(h)
		j := 0
		for j < m {
			if s[i][j] == 75 {
				e = i
				g = j
				s[i][j] = 46
			}
			if s[i][j] == 84 {
				d = i
				v = j
				s[i][j] = 46
			}
			j++
		}
		i++
	}

	F(d, v)
	if r < M {
		P(r)
		return
	}
	P(S)
	if w[e][g] < 1 {
		P("No key")
		return
	}
	P(w[e][g]-1, " ")
	for l < n {
		i = 0
		for i < m {
			if s[l][i] == 45 {
				s[l][i]++
			}
			i++
		}
		l++
	}
	F(e, g)
	if r < M {
		P(r)
		return
	}
	P(S)
}