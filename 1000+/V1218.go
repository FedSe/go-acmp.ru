package main
import . "fmt"
func main() {
	var (
		a [1000]int
		n, j, i int
		m = 1
	)

	Scan(&n)
	n++

	for i < n {
	i++
		Scan(&a[i])
	}

	for j < n {
		if a[j] >= a[n] {
			m++
		}
	j++
	}

	Print(m)
}