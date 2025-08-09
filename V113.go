package main
import . "fmt"
func main() {
	var (
		d          [1e3][1e3]int
		x          [1e3]string
		n, m, j, i int
	)

	Scan(&n)
	for j < n {
		Scan(&x[j])
		j++
	}

	for i < n {
		j = 0
		for j < n {
			if x[i][j] > 48 {
				y := &d[i][j]
				*y = 1
				if i*j > 0 {
					v := d[i-1]
					q := v[j]
					g := d[i][j-1]
					if q > g {
						q = g
					}
					g = v[j-1]
					if q > g {
						q = g
					}
					*y += q
				}
				if *y > m {
					m = *y
				}
			}
			j++
		}
		i++
	}

	Print(m * m)
}