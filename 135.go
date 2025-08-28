package main
import . "fmt"
func main() {
	var (
		a [101][101]int
		n, i, k, l, j int
	)

	Scan(&n)

	for i < n {
		for
		j = 0
		j < n
		{
			Scan(&a[i][j])
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
				v := a[i][k]+a[k][j]
				if a[i][j] > v {
					a[i][j] = v
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
			Println(a[l][j])
		j++
		}
	l++
	}
}