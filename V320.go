package main
import . "fmt"
func main() {
	var n, m, i int
	a := [51]int{1}

	Scan(&m, &n)
	for i < n {
		i++
		a[i] = a[i-1]
		if i >= m {
			a[i] += a[i-m]
		}
	}

	Print(a[n])
}