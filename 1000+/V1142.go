package main
import . "fmt"
func main() {
	var (
		a, b int
		k = 1
		m = 1
		x = 1
		y = 1
	)
	Scan(&b)
	for b > 0 {
		a = b
		Scan(&b)
		if b > a {
			k++
		} else {
			if k > m {
				m = k
			}
			k = 1
		}
		if b < a {
			x++
		} else {
			if x > y {
				y = x
			}
			x = 1
		}
	}

	if m > y {
		y = m
	}

	Print(y)
}