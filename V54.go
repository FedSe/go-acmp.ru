package main
import . "fmt"
func main() {
	var (
		n, m, x, i, j, l int
		w, c [100]int
	)
	Scan(&n, &m)

	for i < n {
		w[i] = 1001
		for
		j := 0
		j < m
		{
			Scan(&x)
			if i == 0 {
				c[j] = -1001
			}
			if w[i] > x { w[i] = x }
			if c[j] < x { c[j] = x }
		j++
		}
	i++
	}

	a := c[0]
	for j < m {
		x = c[j]
		if x < a {
			a = x
		}
	j++
	}

	k := w[0]
	for l < n {
		x = w[l]
		if x > k {
			k = x
		}
	l++
	}

	Print(k, a)
}