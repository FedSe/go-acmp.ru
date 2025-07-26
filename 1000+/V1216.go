package main
import . "fmt"
func main() {
	var (
		a [1001]int
		n, l, r, i int
	)
	Scan(&n)
	for i < n {
	i++
		Scan(&a[i])
	}
	Scan(&l, &r)

	i = l
	for l < r {
	l++
		if a[l] > a[i] {
			i = l
		}
	}

	Print(a[i], i)
}