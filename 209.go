package main
import . "fmt"
func main() {
	var (
		x, y          [1e3]int
		n, q, l, j, i int
	)

	Scan(&n)
	for l < n {
		Scan(&x[l], &y[l])
		l++
	}

	for j < n {
		l = (j + 1) % n
		q += x[j]*y[l] - x[l]*y[j]
		j++
	}

	if q < 0 {
		q = -q
	}

	for i < n {
		j = (i + 1) % n
		l = x[j] - x[i]
		if l < 0 {
			l = -l
		}
		d := y[j] - y[i]
		if d < 0 {
			d = -d
		}
		for d > 0 {
			l, d = d, l%d
		}
		q -= l
		i++
	}

	Print((q + 2) / 2)
}