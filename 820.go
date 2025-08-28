package main
import . "fmt"
type P struct {
	a [4]string
}
func main() {
	var (
		s P
		i = 0
		u = map[P]int{}
		W = Print
	)

	for i < 4 {
		Scan(&s.a[i])
		i++
	}

	q := []P{s}
	u[s] = 0
	for len(q) > 0 {
		p := q[0]
		e := 1
		i = 0
		for i < 4 {
			j := 0
			for j < 4 {
				if p.a[i][j] != p.a[0][0] {
					e = 0
				}
				j++
			}
			i++
		}
		if e > 0 {
			W(u[p])
			return
		}
		q = q[1:]
		i = 0
		for i < 4 {
			j := 0
			for j < 4 {
				y := p
				z := -1
				for z < 2 {
					l := -1
					for l < 2 {
						w := l
						if z*z+w*w < 2 {
							h := i + z
							g := j + w
							if 0 <= h && h < 4 && 0 <= g && g < 4 {
								y.a[h] = y.a[h][:g] + string(217-y.a[h][g]) + y.a[h][g+1:]
							}
						}
						l++
					}
					z++
				}
				_, o := u[y]
				if !o {
					u[y] = u[p] + 1
					q = append(q, y)
				}
				j++
			}
			i++
		}
	}

	W("Impossible")
}