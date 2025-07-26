package main
import . "fmt"
func main() {
	var (
		c [101]int
		r [6001][3]int
		n, m, i, k, y, d, p, o int
		j = 1
	)

	Scan(&n, &m)

	for k < m {
		for
		o = 0
		o < 3
		{
			Scan(&r[k][o])
		o++
		}
	k++
	}

	for j <= n {
		c[j] = j
	j++
	}

	for i < n-1 {
		d = 30001

		for
		k = 0
		k < m
		{
			y = r[k][2]
			if c[r[k][0]] != c[r[k][1]] && y < d {
				d = y
				j = k
			}
		k++
		}
		p += d
		o = c[r[j][0]]
		y = c[r[j][1]]

		for
		j = 0
		j < n
		{
		j++
			if c[j] == y {
				c[j] = o
			}
		}
	i++
	}

	Print(p)
}