package main
import . "fmt"
func main() {
	var (
		y, i, l, n int
		a          [10][10]int
	)

	Scan(&n)
	for 0 < n {
		Scan(&y, &i)
		a[y][i] = 1
		n--
	}

	for l < 10 {
		i = 0
		for i < 10 {
			if a[l][i] > 0 {
				n += 4 - a[l-1][i] - a[l+1][i] - a[l][i-1] - a[l][i+1]
			}
			i++
		}
		l++
	}

	Print(n)
}