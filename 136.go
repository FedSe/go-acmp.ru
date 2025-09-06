package main
import . "fmt"
func main() {
	var (
		a                [101][101]int
		n, i, j, k, l, o int
		F                = 9999
	)

	Scan(&n)
	for i < n {
		j = 0
		for j < n {
			Scan(&a[i][j])
			if a[i][j] < 0 {
				a[i][j] = F
			}
			j++
		}
		i++
	}

	for k < n {
		i = 0
		for i < n {
			j = 0
			for j < n {
				m := a[i][k] + a[k][j]
				if a[i][j] > m {
					a[i][j] = m
				}
				j++
			}
			i++
		}
		k++
	}

	for l < n {
		j = 0
		for j < n {
			m := a[l][j]
			if m < F && m > o {
				o = m
			}
			j++
		}
		l++
	}

	Print(o)
}