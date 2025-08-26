package main
import . "fmt"
type T string
type R struct {
	s T
	i int
}
func main() {
	var (
		a, b, c, d T
		k          = 0
		D          = []int{0, 0, -1, 1}
	)

	Scan(&a, &b, &c, &d)
	a += b
	c += d
	if a != c {
		k--
		v := map[any]int{a: 1}
		q := []R{{a, 0}}
		for len(q) > 0 {
			var d, p, i int
			t := q[0]
			q = q[1:]
			for i < 8 {
				if t.s[i] == 35 {
					p = i
				}
				i++
			}
			i = p / 4
			j := p % 4
			for d < 4 {
				x := i - D[3-d]
				y := j + D[d]
				if x >= 0 && x < 2 && y >= 0 && y < 4 {
					n := x*4 + y
					w := []byte(t.s)
					w[p], w[n] = w[n], w[p]
					b = T(w)
					if b == c {
						k = t.i + 1
						goto A
					}
					if v[b] < 1 {
						v[b] = 1
						q = append(q, R{b, t.i + 1})
					}
				}
				d++
			}
		}
	}
A:
	Print(k)
}