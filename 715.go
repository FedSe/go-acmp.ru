package main
import . "fmt"
func main() {
	var (
		n, m, i, k, l int
		a             [200]string
	)

	Scan(&n, &m)
	for i < n*2 {
		Scan(&a[i])
		i++
	}
	for l < n {
		i = 0
		for i < m {
			if a[l][i] == a[n+l][i] {
				k++
			}
			i++
		}
		l++
	}

	Print(k)
}