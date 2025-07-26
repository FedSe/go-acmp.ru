package main
import . "fmt"
func main() {
	var n, r, d, a int
	Scan(&n)
	for 0 < n {
		Scan(&a)

		k := 0
		f := a
		i := 2
		for f > 1 {
			if f % i < 1 {
				k++
			}
			for f % i < 1 {
				f /= i
			}
		i++
		}

		if d < k || (d == k && a < r) {
			r = a
			d = k
		}
	n--
	}

	Print(r)
}