package main
import . "fmt"
func main() {
	var (
		n, m, k, y, c, s, t, i, l int
		a                         [202][202]int
	)
    
	Scan(&n, &m, &k)
	for s < n+2 {
		t = 0
		for t < m+2 {
			a[s][t] = 46
			t++
		}
		s++
	}
    
	for c < k {
		Scan(&t, &y)
		a[t][y] = 42
		c++
	}
    
	for l < n {
		l++
		x := 1
		for x <= m {
			if a[l][x] > 42 {
				c = 0
				y = l - 1
				s = x - 1
				t = x + 1
				for _, p := range [][2]int{
					{y, s}, {y, x},
					{y, t}, {l, s},
					{l, t}, {y + 2, s},
					{y + 2, x}, {y + 2, t}} {
					if a[p[0]][p[1]] < 43 {
						c++
					}
				}
				if c > 0 {
					a[l][x] = c + 48
				}
			}
			x++
		}
	}

	for i < n {
		i++
		l = 1
		for l <= m {
			Printf("%c", a[i][l])
			l++
		}
		Println()
	}
}