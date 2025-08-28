package main
import . "fmt"
func main() {
	var (
		a       [2e3]int
		n, k, i int
	)

	Scan(&n)
	for i <= n {
		Scan(&a[i])
		i++
	}
	for 0 < i {
		i--
		if a[i] == a[n] {
			k++
		}
	}

	Print(k - 1)
}