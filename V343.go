package main
import . "fmt"
func main() {
	var (
		s             [53][53]int
		n, m, v, y, x int
	)

	Scan(&n, &m)
	for y < 53 {
		x = 0
		for x < 53 {
			o := 1
			if y <= n && x <= m && x*y > 0 {
				o = 0
			}
			s[y][x] = o
			x++
		}
		y++
	}

	Scan(&n)
	for n > 0 {
		Scan(&m, &y, &x)
		a, b, c, d := &s[y+1][x+1], &s[y][x+1], &s[y+1][x], &s[y][x]
		w := []*int{a, b, c}
		if m == 2 {
			w = []*int{a, c, d}
		}
		if m == 3 {
			w = []*int{a, b, d}
		}
		if m > 3 {
			w = []*int{b, c, d}
		}
		if *w[0]+*w[1]+*w[2] < 1 {
			*w[0] = 1
			*w[1] = 1
			*w[2] = 1
			v += 3
		}
		n--
	}

	Print(v)
}