package main
import . "fmt"
type S struct{ u, v, c, i int }
func main() {
	var (
		h                   []int
		n, m, u, v, c, l, o int
		q                   = int16(1 << 14)
		P                   = Println
	)

	Scan(&n, &m)
	if n < 2 {
		P(0, 0)
		return
	}

	y := make([]S, m)
	for l < m {
		Scan(&u, &v, &c)
		y[l] = S{u - 1, v - 1, c, l + 1}
		l++
	}

	z := 1 << n
	d := make([]int16, z)
	for i := range d {
		d[i] = q
	}
	d[1] = 0

	for o < z {
		if d[o] != q {
			for _, e := range y {
				if (o>>e.u)&1 > (o>>e.v)&1 {
					m = o | (1 << e.v)
					x := d[o] + int16(e.c)
					if x < d[m] {
						d[m] = x
					}
				}
			}
		}
		o++
	}

	z--
	c = z
	for c != 1 {
		l = 0
		m = 0
		n = 0
		for _, e := range y {
			if (c>>e.v)&1 > 0 {
				u = c ^ (1 << e.v)
				if (u>>e.u)&1 == 1 && d[u]+int16(e.c) == d[c] {
					l = e.i
					m = u
					n = 1
					break
				}
			}
		}
		if n < 1 {
			break
		}
		h = append(h, l)
		c = m
	}

	l = len(h)
	P(d[z], l)
	for l > 0 {
		l--
		P(h[l])
	}
}