package main
import . "fmt"
var (
	s [6][2][3]byte
	e []int
	u = []int{0}
	R = 12
	i = 0
	w = ""
	P = Print
)

func D(a, b, c, d *byte) {
	t := *a
	*a = *b
	*b = *c
	*c = *d
	*d = t
}

func A() {
	D(&s[0][0][0], &s[0][1][0], &s[0][1][1], &s[0][0][1])
	D(&s[1][0][1], &s[5][0][0], &s[3][1][0], &s[4][1][1])
	D(&s[1][1][1], &s[5][0][1], &s[3][0][0], &s[4][1][0])
}

func B() {
	D(&s[3][0][0], &s[3][1][0], &s[3][1][1], &s[3][0][1])
	D(&s[0][0][1], &s[5][0][1], &s[2][1][0], &s[4][0][1])
	D(&s[0][1][1], &s[5][1][1], &s[2][0][0], &s[4][1][1])
}

func C() {
	D(&s[4][0][0], &s[4][1][0], &s[4][1][1], &s[4][0][1])
	D(&s[0][0][0], &s[3][0][0], &s[2][0][0], &s[1][0][0])
	D(&s[0][0][1], &s[3][0][1], &s[2][0][1], &s[1][0][1])
}

func H(a int) {
	if a < 0 {
		a = -a
		if a < 2 {
			A()
			A()
		}
		if a == 2 {
			B()
			B()
		}
		if a > 2 {
			C()
			C()
		}
	}
	if a < 2 {
		A()
	}
	if a == 2 {
		B()
	}
	if a > 2 {
		C()
	}
}

func F() {
	i := 0
	for i < 6 {
		j := 0
		for j < 2 {
			k := 0
			for k < 2 {
				if s[i][j][k] != s[i][0][0] {
					if len(u)+1 < R {
						l := -3
						for l < 4 {
							if l != 0 {
								H(l)
								u = append(u, l)
								F()
								u = u[:len(u)-1]
								H(-l)
							}
							l++
						}
					}
					return
				}
				k++
			}
			j++
		}
		i++
	}
	i = len(u)
	if i < R {
		R = i
		e = make([]int, R)
		copy(e, u)
	}
}

func main() {
	for i < 2 {
		j := 0
		for j < 6 {
			Scan(&w)
			s[j][i][0] = w[0]
			s[j][i][1] = w[1]
			j++
		}
		i++
	}

	F()
	if R == 1 {
		P("Solved")
		return
	}
	i = 1
	for i < len(e) {
		b := e[i]
		if b < 0 {
			b = -b
		}
		P(string(" FRU"[b]))
		if e[i] < 0 {
			P("'")
		}
		i++
	}
}