package main
import . "fmt"
func main() {
	var (
		a, b, c [1000]int
		i, n, t, v, o int
	)

	Scan(&n)
	for i < n {
		Scan(&a[i], &b[i], &c[i])
	i++
	}

	Scan(&t)
	for o < t {
		i = 0
		h := 0
		for i < n {
			if a[i] <= o && b[i] > o {
				h += c[i]
			}
		i++
		}
		v += h
		if v < 0 {
			v = 0
		}
	o++
	}

	Print(v)
}