package main
import . "fmt"
func main() {
	var (
		a          [1001]int
		n, l, r, i int
		S = Scan
	)

	S(&n)
	for i < n {
		i++
		S(&a[i])
	}
	S(&l, &r)

	i = l
	for l < r {
		l++
		if a[l] > a[i] {
			i = l
		}
	}

	Print(a[i], i)
}