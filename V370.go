package main
import . "fmt"
func main() {
	var (
		x, y    [5e4]float64
		a       = 0.
		n, i, l int
	)

	Scan(&n)
	for i < n {
		Scan(&x[i], &y[i])
		i++
	}

	for l < n {
		i = (l + 1) % n
		a += x[l]*y[i] - x[i]*y[l]
		l++
	}

	if a < 0 {
		a = -a
	}

	Printf("%.1f", a/2)
}