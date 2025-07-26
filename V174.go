package main
import . "fmt"
func main() {
	var (
		f, n, i, l int
		s = 0.
		a [40]float64
	)
	Scan(&n)

	for l < n {
		Scan(&a[l])
	l++
	}

	for f < n-1 {
	f++
		j := f
		k := a[f]
		for j > 0 && k < a[j-1] {
			a[j] = a[j-1]
			j--
		}
		a[j] = k
	}

	Scan(&s)
	for i < n {
		if s < a[i] {
			s = (s + a[i]) / 2
		}
	i++
	}

	Printf("%.6f", s)
}