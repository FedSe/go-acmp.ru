package main
import . "fmt"
func main() {
	var (
		a [52]int
		n, m, i int
	)
	Scan(&m, &n)

	for i < n {
		a[i]++
	i++
	}

	for
	i = m
	i <= n
	{
		a[i] = a[i-1] + a[i-m]
	i++
	}

	Print(a[n])
}