package main
import . "fmt"
func main() {
	var (
		n, m, i, j, x int
		q = 1001
		r             = -q
		a             [100]int
	)

	Scan(&n, &m)
	for i < n {
		s := q
		l := 0
		for l < m {
			Scan(&x)
			if x < s {
				s = x
			}
			if x > a[l] || i < 1 {
				a[l] = x
			}
			l++
		}
		if s > r {
			r = s
		}
		i++
	}

	n = q
	for j < m {
		if a[j] < n {
			n = a[j]
		}
		j++
	}

	Print(r, n)
}