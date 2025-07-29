package main
import . "fmt"
func main() {
	var (
		a       [1e3]int
		n, x, i int
		S = Scan
	)

	S(&n)
	for i < n {
		S(&a[i])
		i++
	}

	S(&x)

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