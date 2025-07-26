package main
import . "fmt"
func main() {
	var (
		n, m, i, j int
		s = ""
		a [200]string
	)

	Scan(&m, &n)
	for j < n*2 {
		Scan(&a[j])
	j++
	}

	Scan(&s)
	for i < n {
		for
		j = 0
		j < m
		{
			Print(s[ a[i][j]*2 + a[n+i][j] - 144 ] - 48)
		j++
		}
		Println()
	i++
	}
}