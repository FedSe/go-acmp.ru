package main
import . "fmt"
func main() {
	var (
		n, m, t, i, j, k int
		a                [100]int
		S = Scan
	)

	S(&n)
	for i < n {
		S(&a[i])
		i++
	}

	S(&m)
	for j < m {
		S(&t)
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