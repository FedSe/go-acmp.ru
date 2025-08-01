package main

import . "fmt"

type R struct {
	x, y, z, c int
}

type U struct {
	p, r [100]int
}

func H(x int, u *U) int {
	if u.p[x] != x {
		u.p[x] = H(u.p[x], u)
	}
	return u.p[x]
}

func main() {
	var (
		s             [100]R
		t             [100]int
		n, i, l, j, k int
		c             = map[int]int{}
	)

	Scan(&n)
	for l < n {
		var x, y, z, c, r int
		Scan(&x, &y, &z, &c, &r)

		if x > z {
			x, z = z, x
		}
		if y > c {
			y, c = c, y
		}

		s[l] = R{x - r, y - r, z + r, c + r}
		l++
	}

	for k < n {
		t[k] = k
		k++
	}
	u := &U{t, [100]int{}}

	for i < n {
		l = i + 1
		for l < n {
			w := s[i]
			v := s[l]
			if w.x <= v.z && v.x <= w.z && w.y <= v.c && v.y <= w.c {
				p := H(i, u)
				k = H(l, u)
				if p != k {
					if u.r[p] < u.r[k] {
						u.p[p] = k
					} else {
						u.p[k] = p
						if u.r[p] > u.r[k] {
							u.r[p]++
						}
					}
				}
			}
			l++
		}
		i++
	}

	for j < n {
		c[H(j, u)] = 1
		j++
	}

	Print(len(c))
}