package main
import . "fmt"
func main() {
	var (
		y, i, l, n int
		a [10][10]int
	)

	for i < 10 {
		for
		n = 0
		n < 10
		{
			a[i][n] = 1
		n++
		}
	i++
	}

	Scan(&n)
	for 0 < n {
		Scan(&y, &i)
		a[y][i] = 0
	n--
	}

	for l < 10 {
		for
		i = 0
		i < 10
		{
			if a[l][i] < 1 {
				n += a[l-1][i] + a[l+1][i] + a[l][i-1] + a[l][i+1]
			}
		i++
		}
	l++
	}

	Print(n)
}