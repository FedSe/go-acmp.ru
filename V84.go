package main
import . "fmt"
func main() {
	var (
		n, m, i, j, k int
		a [100]string
	)
	Scan(&n, &m)

	for j < n {
		Scan(&a[j])
	j++
	}

	l, r, u, d := m, -1, n, -1
	for i < n {
		for
		j = 0
		j < m
		{
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
		for
		j = 0
		j < m
		{
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