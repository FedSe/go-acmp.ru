package main
import . "fmt"

type P struct {
	r, b, e int
}

func main() {
	var (
		n, e, t, h, m, v, l int
		p []*P
		q []int
		u = "%02d:%02d:%02d "
	)

	Scanln(&n, &e)
	for 0 < n {
		q = append(q, len(p))
		var o P
		o.b--
		Scanf(u + "%d\n", &h, &m, &v, &l)
		v += h*3600 + m*60
		o.r = l
		p = append(p, &o)

		if len(p) == e {
			if t < v { t = v }
			for len(q) > 0 {
				v = q[0]
				q = q[1:]
				c := p[v]

				if c.b < 0 {
					c.b = t
				}

				if c.r < 11 {
					t += c.r
					c.e = t
				} else {
					c.r -= 10
					t += 10
					q = append(q, v)
				}
			}
			for _, c := range p {
				Printf(u + u, c.b/3600%24, c.b/60%60, c.b%60,
					c.e/3600%24, c.e/60%60, c.e%60)
			}
			p = []*P{}
		}
	n--
	}
}