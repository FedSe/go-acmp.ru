package main
import . "fmt"
func main() {
	var (
		a [50][50]int
		n, i, l, k, m int
		F = 99999
	)

	Scan(&n)

	for i < n {
		for
		j := 0
		j < n
		{
			Scan(&a[i][j])
			if a[i][j] < 0 {
				a[i][j] = F
			}
		j++
		}
	i++
	}

	for k < n {
		for
		i = 0
		i < n
		{
			for
			j := 0
			j < n
			{
				f := a[i][k] + a[k][j]
				if a[i][j] > f {
					a[i][j] = f
				}
			j++
			}
		i++
		}
	k++
	}

	for l < n {
		for
		i = 0
		i < n
		{
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