package main
import . "fmt"
func main() {
	var (
		g                      [102][102]int
		n, m, k, a, b, r, d, x int
	)

	Scan(&n, &m, &k)
	for 0 < k {
		Scan(&a, &b, &r, &d)
		for a <= r {
			c := b
			for c <= d {
				k := -1
				for k < 2 {
					v := -1
					for v < 2 {
						g[a+k][c+v] = 1
						v++
					}
					k++
				}
				c++
			}
			a++
		}
		k--
	}

	h := make([]int, m+1)

	for k < n {
		var s, p []int
		k++
		a = 0
		for a < m {
			a++
			h[a]++
			if g[k][a] > 0 {
				h[a] = 0
			}
		}
		p = append(h, 0)
		b = 0
		for b < len(p) {
			a = len(s)
			s = append(s, b)
			for a > 0 && p[s[a-1]] > p[b] {
				a--
				r = p[s[a]]
				s = s[:a]
				d = b
				if a > 0 {
					d -= s[a-1] + 1
				}
				if r*d > x {
					x = r * d
				}
			}
			b++
		}
	}

	Print(x)
}