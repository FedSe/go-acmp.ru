package main
import . "fmt"
func main() {
	var (
		n = 0
		k = 1
		i = 2
		a [100001]int
	)
	Scan(&n)

	a[1] = 1
	for i <= n {

		a[i] = a[i-1] + 2

		if i == a[k] {
			a[i]++
			k++
		} else if i > a[k] {
			k++
		}

	i++
	}

	Print(a[n])
}