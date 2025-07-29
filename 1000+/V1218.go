package main
import . "fmt"
func main() {
	var (
		a       [1e3]int
		n, j, i int
	)

	Scan(&n)
    
	n++
	for i < n {
		i++
		Scan(&a[i])
	}

	i = 1
	for j < n {
		if a[j] >= a[n] {
			i++
		}
		j++
	}

	Print(i)
}