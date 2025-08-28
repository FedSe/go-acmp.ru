package main
import . "fmt"
func main() {
	var (
		n, x, m, j, i int
		s, d, c       [2e3]int
		S             = Scan
	)

	S(&n)
	for 0 < n {
		S(&x)
		c[n] = x
		n--
	}
	S(&n)

	for j < n {
		S(&s[j])
		if s[j] > m {
			m = s[j]
		}
		j++
	}

	d[0] = 1
	for _, v := range c {
		j = v
		for j <= m {
			if d[j-v] > 0 {
				d[j] = 1
			}
			j++
		}
	}
	for i < n {
		Println(d[s[i]])
		i++
	}
}