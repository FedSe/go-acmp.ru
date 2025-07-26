package main
import . "fmt"
func main() {
	var (
		a[100]int
		n, m, i int
	)
	Scan(&n, &m)
	for i < n {
		j := 0
		for j < m {
			Scan(&a[j])
		j++
		}
		j = 0
		for j < m {
			Println(a[m-1-j])
		j++
		}
	i++
	}
}