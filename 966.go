package main
import . "fmt"
func main() {
	var (
		a, b, c       [1e3]int
		i, n, t, v, o int
		S             = Scan
	)

	S(&n)
	for i < n {
		S(&a[i], &b[i], &c[i])
		i++
	}

	S(&t)
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