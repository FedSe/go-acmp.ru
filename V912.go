package main
import . "fmt"
func main() {
	var (n, t, r int
	c = 1
	i = 1
	a [100]int
)
	Scan(&n)

	for 0 < n {
		Scan(&t)
		a[t-1]++
	n--
	}

	for i < 100 {
		if a[i] > a[r] {
			r = i
			c = 1
		} else if a[i] == a[r] {
			c++
		}
	i++
	}

	if c > 1 {
		r = -1
	}

	Print(r + 1)
}