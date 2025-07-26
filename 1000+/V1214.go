package main
import . "fmt"
func main() {
	var (
		a [1001]int
		n, k, i, j int
	)
	Scan(&n)
	for i < n+1 {
		Scan(&a[i])
	i++
	}

	for j < n {
		if a[j] == a[n] {
			k++
		}
	j++
	}
	Print(k)
}