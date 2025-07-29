package main
import . "fmt"
func main() {
	var (
		n, s, c, i, j int
		a             [1e4]int
		t             = 1 > 0
	)

	Scan(&n)
	for j < n {
		Scan(&a[j])
		j++
	}

	n--
	for i <= n {
		j = a[n]
		if a[i] >= j {
			j = a[i]
			i++
			n++
		}

		c += j

		if t {
			c -= j
			s += j
		}

		t = !t
		n--
	}

	Print(s, ":", c)
}