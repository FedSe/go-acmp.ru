package main
import . "fmt"
func main() {
	var (
		a       [1e3]int
		n, x, i int
	)

	Scan(&n)
	for i < n {
		Scan(&a[i])
		i++
	}

	Scan(&x)

	m := 2001
	c := x + m

	for 0 < n {
		n--
		i = a[n]
		d := x - i
		if d < 0 {
			d = -d
		}
		if d <= m {
			if i < c {
				c = i
			}
			if d != m {
				c = i
				m = d
			}
		}
	}

	Print(c)
}