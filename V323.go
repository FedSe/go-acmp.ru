package main
import . "fmt"
func main() {
	var (
		a [999]int
		m = 0
		i = 2
	)
	Scan(&m)

	for i*i < m {
		for
		j := i + i
		j < m
		j += i {
			a[j] = 1
		}
	i++
	}
	i = 2
	for !(a[i]<1 && a[m-i]<1) {
		i++
	}
	Println(i, m-i)
}