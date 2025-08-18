package main
import . "fmt"
func main() {
	var (
		a          [100][100]int
		n, m, i, l int
	)

	Scan(&n, &m)
	for i < n {
		j := 0
		for j < m {
			Scan(&a[i][j])
			j++
		}
		i++
	}

	q := a[0][0]
	for l < n {
		var p [100]int
		i = l
		for i < n {
			j := 0
			for j < m {
				p[j] += a[i][j]
				j++
			}

			w := p[0]
			o := p[0]
			j = 1
			for j < m {
				if w < 0 {
					w = p[j]
				} else {
					w += p[j]
				}
				if w > o {
					o = w
				}
				j++
			}
			if o > q {
				q = o
			}
			i++
		}
		l++
	}

	Print(q)
}