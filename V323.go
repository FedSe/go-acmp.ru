package main
import . "fmt"
func main() {
	var (
		a [1e3]int
		m = 0
		i = 2
	)

	Scan(&m)
	for i*i < m {
		j := i + i
		for j < m {
			a[j] = 1
			j += i
		}
		i++
	}

	i = 2
	for a[i]+a[m-i] > 0 {
		i++
	}

	Print(i, m-i)
}