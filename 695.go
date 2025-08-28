package main
import . "fmt"
type T []int
func main() {
	var (
		g = T{-2, -2, -1, -1, 1, 1, 2, 2, -1, 1, -2, 2, -2, 2, -1, 1}
		v = map[int]int{}
		a = ""
		b = a
		d = 0
	)

	Scan(&a, &b)
	j := int(a[0] - 65)
	l := int(a[1] - 49)

	if a != b {
		d--
		q := []T{{j, l, 0}}
		v[j*10+l] = 1
		for len(q) > 0 {
			var (
				z []T
				c = q[0][0]
				r = q[0][1]
				m = q[0][2]
				i = 0
				x = -1
			)
			q = q[1:]
			if (c+r)&1 < 1 {
				for x < 3 {
					y := -1
					for y < 3 {
						i = 1
						for i < 9 {
							s := c + x*i
							k := r + y*i
							if s < 0 || s > 8 || k < 0 || k > 8 {
								break
							}
							z = append(z, T{s, k})
							i++
						}
						y += 2
					}
					x += 2
				}
			} else {
				for i < 8 {
					z = append(z, T{c + g[i+8], r + g[i]})
					i++
				}
			}

			for _, o := range z {
				i = o[0]
				r = o[1]
				c = i*10 + r
				if i == int(b[0]-65) && r == int(b[1]-49) {
					d = m + 1
					goto A
				}
				if i >= 0 && r >= 0 && v[c] < 1 {
					v[c] = 1
					q = append(q, T{i, r, m + 1})
				}
			}
		}
	}
A:
	Print(d)
}