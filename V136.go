package main
import . "fmt"
func main() {
	var (
		a [101][101]int
		n, i, j, k, l, o int
		F = 9999
	)

	Scan(&n)
	
	for i < n {
		for
		j = 0
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
			j = 0
			j < n
			{
				m := a[i][k]+a[k][j]
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
		for
		j = 0
		j < n
		{
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