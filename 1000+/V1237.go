package main
import . "fmt"
func main() {
	var (
		a [10000]int
		n, m, i, l int
	)

	Scan(&n, &m)
	n *= m
	for i < n {
		Scan(&a[i])
	i++
	}
	for l < n {
		Scan(&i)
		Println(a[l] + i)
	l++
	}
}