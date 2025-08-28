package main
import . "fmt"
func main() {
	var (
		a                [50][50]int
		n, i, l, k, m, q int
		F                = 99999
	)

	Scan(&n)
	for i < n {
		j := 0
		for j < n {
			Scan(&q)
			if q < 0 {
				q = F
			}
			a[i][j] = q
			j++
		}
		i++
	}

	for k < n {
		i = 0
		for i < n {
			q = 0
			for q < n {
				f := a[i][k] + a[k][q]
				if a[i][q] > f {
					a[i][q] = f
				}
				q++
			}
			i++
		}
		k++
	}

	for l < n {
		i = 0
		for i < n {
			k = a[l][i]
			if k < F && k > m {
				m = k
			}
			i++
		}
		l++
	}

	Print(m)
}