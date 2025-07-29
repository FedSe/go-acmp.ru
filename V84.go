package main
import . "fmt"
func main() {
	var (
		n, m, i, j, k int
		a             [100]string
		S             = Scan
		r             = -1
		d             = r
	)

	S(&n, &m)
	for j < n {
		S(&a[j])
		j++
	}

	l, u := m, n
	for i < n {
		j = 0
		for j < m {
			if a[i][j] == 42 {
				if i < u {
					u = i
				}
				if i > d {
					d = i
				}
				if j < l {
					l = j
				}
				if j > r {
					r = j
				}
			}
			j++
		}
		i++
	}

	for k < n {
		j = 0
		for j < m {
			s := "."
			if k >= u && k <= d && j >= l && j <= r {
				s = "*"
			}
			Print(s)
			j++
		}
		Println()
		k++
	}
}