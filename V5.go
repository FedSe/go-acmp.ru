package main
import . "fmt"
func main() {
	var (
		n, c, r, t int
		a          [100]int
		s          = "YES"
	)
	Scan(&n)

	for n > 0 {
		Scan(&c)
		if c%2 > 0 {
			Println(c)
			r++
		} else {
			a[t] = c
			t++
		}
		n--
	}

	Println()
	for n < t {
		Println(a[n])
		n++
	}

	if t < r {
		s = "NO"
	}

	Println(s)
}