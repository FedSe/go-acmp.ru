package main
import . "fmt"
func main() {
	var (
		n, k, l, i int
		d, f, o    [13]int
		a          = 1
	)
	f[0] = 1

	Scan(&n, &k)
	for l < 12 {
		if l < k {
			Scan(&d[l])
		}
		l++
		f[l] = f[l-1] * l
	}

	for i < k {
		l = 1
		w := 0
		h := 1
		for h < d[i] {
			if o[h] < 1 {
				w++
			}
			h++
		}

		h = n - i - 1
		y := k - i - 1

		if y > 0 {
			l = f[h] / f[h-y]
		}

		a += w * l
		o[d[i]] = 1
		i++
	}

	Print(a)
}