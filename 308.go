package main
import . "fmt"
func main() {
	var (
		a, b, c, t int
		v          = map[int]int{}
		w          = v
		n          = 1000
		g          = n * n
		m          = g
		P          = Print
	)
	Scan(&a, &b, &c, &t)

	q := []int{a * g}

	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		u := v[x]

		y := (x / n) % n
		z := x % n
		x /= g

		if m > u {
			F := func(f, o *int, l int) {
				if *f > 0 && *o < l {
					d := l - *o
					if d > *f {
						d = *f
					}
					h, j := *f, *o
					*f -= d
					*o += d
					s := x*g + y*n + z
					if w[s] < 1 {
						w[s] = 1
						v[s] = u + 1
						if x == t {
							if u+1 < m {
								m = u + 1
							}
						} else {
							q = append(q, s)
						}
					}
					*f, *o = h, j
				}
			}

			F(&x, &y, b)
			F(&x, &z, c)
			F(&y, &z, c)
			F(&y, &x, a)
			F(&z, &x, a)
			F(&z, &y, b)
		}
	}

	if m < g {
		P(m)
	} else {
		P("IMPOSSIBLE")
	}
}