package main
import . "fmt"
func main() {
	var (
		n, m, k, x, y, c, s, t int
		a [202][202]int
	)
	Scan(&n, &m, &k)

	for s < n+2 {
		for
		t = 0
		t < m+2
		{
			a[s][t] = 46
		t++
		}
	s++
	}

	for c < k {
		Scan(&x, &y)
		a[x][y] = 42
	c++
	}

	for
	k = 1
	k <= n
	{
		for
		x = 1
		x <= m
		{
			if a[k][x] > 42 {
				c = 0
				y = k-1
				s = x-1
				t = x+1
				if a[y][s] < 43 { c++ }
				if a[y][x] < 43 { c++ }
				if a[y][t] < 43 { c++ }
				if a[k][s] < 43 { c++ }
				if a[k][t] < 43 { c++ }
				y += 2
				if a[y][s] < 43 { c++ }
				if a[y][x] < 43 { c++ }
				if a[y][t] < 43 { c++ }
				if c > 0 {
					a[k][x] = c+48
				}
			}
		x++
		}
	k++
	}

	for
	k = 1
	k <= n
	{
		for
		x = 1
		x <= m
		{
			Printf("%c", a[k][x])
		x++
		}
		Println()
	k++
	}

}