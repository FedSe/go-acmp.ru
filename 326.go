package main
import . "fmt"
func main() {
	var (
		n, m, i, k, l int
		a, c          [201]int
		P = Println
	)

	Scan(&n)
	for i < n {
		Scan(&a[i])
		c[a[i]+100]++
		i++
	}

	for j, v := range c {
		if v > c[m] {
			m = j
		}
	}

	i = m - 100
	for k < n {
		if a[k] != i {
			P(a[k])
		}
		k++
	}

	for l < c[m] {
		P(i)
		l++
	}
}