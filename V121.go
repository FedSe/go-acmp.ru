package main
import (
	. "fmt"
	. "sort"
)
func main() {
	var (
		x, d [100]int
		n, l int
		i    = 2
	)

	Scan(&n)
	for l < n {
		Scan(&x[l])
		l++
	}

	Ints(x[:n])
	d[0] = 1e9
	d[1] = x[1] - x[0]
	for i < n {
		a := d[i-1]
		l = x[i-1]
		if i == 2 {
			d[i] = x[i] - l + a
		} else {
			if d[i-2] < a {
				a = d[i-2]
			}
			d[i] = a + x[i] - l
		}
		i++
	}
	Print(d[n-1])
}