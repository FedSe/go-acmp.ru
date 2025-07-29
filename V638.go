package main
import . "fmt"
func main() {
	var (
		u, o                      [2e5]int
		a, b, d, f, k, c, t, i, j int
	)

	Scan(&a, &b, &d, &j, &f)
	for 0 < j {
		Scan(&t)
		u[t] = 1
		j--
	}

	for i < a {
		i++
		o[i] = u[f]
		if f == d {
			f = 0
		}
		f++
	}

	Scan(&j)
	for 0 < j {
		Scan(&t)
		o[t] = 1
		j--
	}

	for j < b {
		j++
		if o[j] > 0 {
			k++
		}
	}
	if k < 1 {
		c++
	}

	i = b
	for i < a {
		i++
		if o[i] > 0 {
			k++
		}
		if o[i-b] > 0 {
			k--
		}
		if k == 0 {
			c++
		}
	}

	Print(c)
}