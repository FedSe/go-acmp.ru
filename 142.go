package main
import . "fmt"
func main() {
	var (
		c                      [101]int
		r                      [6e3][3]int
		n, m, i, k, y, d, p, o,j int
	)

	Scan(&n, &m)
	for k < m {
		o = 0
		for o < 3 {
			Scan(&r[k][o])
			o++
		}
		k++
	}

	for j < n {
		j++
		c[j] = j
	}

	for i < n-1 {
		d = 30001
		k = 0
		for k < m {
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
		j = 0
		for j < n {
			j++
			if c[j] == y {
				c[j] = o
			}
		}
		i++
	}

	Print(p)
}