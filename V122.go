package main
import . "fmt"
func main() {
	var (
		a, d       [1e3]int
		N, q, x, i int
	)

	Scan(&N)
	for i < N {
		Scan(&x)
		a[i] = x
		d[i] = 1
		i++
	}

	i = 1
	for i < N {
		j := 0
		for j < i {
			x = d[j] + 1
			if a[j] < a[i] && x > d[i] {
				d[i] = x
			}
			j++
		}
		i++
	}

	for 0 < N {
		N--
		if d[N] > q {
			q = d[N]
		}
	}

	Print(q)
}