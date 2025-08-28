package main
import . "fmt"
type B string
type A struct {
	a B
	b int
}
func main() {
	var (
		D                = []int{-2, -2, -1, -1, 1, 1, 2, 2, -1, 1, -2, 2, -2, 2, -1, 1}
		a, b, z, w, h, g B
		P                = Print
	)

	Scan(&a, &w, &b, &h, &z, &g)

	a += b + z
	d := []A{{a, 0}}
	q := map[any]int{a: 1}
	w += h + g

	for len(d) > 0 {
		z = d[0].a
		m := d[0].b
		d = d[1:]
		i := 0
		for i < 3 {
			j := 0
			for j < 3 {
				k := 0
				for k < 8 {
					p := i + D[k]
					c := j + D[k+8]
					if p >= 0 && p < 3 && c >= 0 && c < 3 {
						v := []byte(z)
						v[i*3+j] = 46
						v[p*3+c] = z[i*3+j]
						b = B(v)
						if b == w {
							P(m + 1)
							return
						}
						if q[b] < 1 {
							q[b] = 1
							d = append(d, A{b, m + 1})
						}
					}
					k++
				}
				j++
			}
			i++
		}
	}

	P(-1)
}