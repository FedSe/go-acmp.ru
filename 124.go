package main
import . "fmt"
func main() {
	var (
		n, m, i, f int
		a          [100]int
	)

	Scan(&n, &m)
	m *= 2
	for 0 < m {
		Scan(&f)
		a[f-1]++
		m--
	}
	for i < n {
		Println(a[i])
		i++
	}
}