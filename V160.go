package main
import . "fmt"
func main() {
	var (
		d, p [101]int
		n, l int
		r    = 1
		i    = 1
	)

	Scan(&n)
	for l < n {
		l++
		Scan(&p[l])
	}

	for i <= n {
		if d[i] < 1 {
			l = i
			h := 0
			for d[l] < 1 {
				d[l] = 1
				l = p[l]
				h++
			}
			if h > 0 {
				l = r
				m := h
				for m > 0 {
					l, m = m, l%m
				}
				r = r / l * h
			}
		}
		i++
	}

	Print(r)
}