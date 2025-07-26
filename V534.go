package main
import . "fmt"
func main() {
	var (
		n, m, t, i, j, k int
		a [100]int
	)

	Scan(&n)
	for i < n {
		Scan(&a[i])
	i++
	}

	Scan(&m)
	for j < m {
		Scan(&t)
		a[t-1]--
	j++
	}

	for k < n {
		s := "no"
		if a[k] < 0 {
			s = "yes"
		}
		Println(s)
	k++
	}
}