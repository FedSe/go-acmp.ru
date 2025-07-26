package main
import . "fmt"
func main() {
	var (
		n, s, i, j int
		a [1002]int
	)
	Scan(&n)

	for i < n {
		Scan(&a[i])
	i++
	}
	a[n] = a[0]
	a[n+1] = a[1]

	for j < n {
	j++
		i = a[j-1] + a[j] + a[j+1]
		if i > s {
			s = i
		}
	}

	Print(s)
}