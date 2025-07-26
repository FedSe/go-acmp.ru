package main
import . "fmt"
func main() {
	var (
		n, p, s, v, i, m int
		a [100]int
	)
	Scan(&n)

	for i < n {
		Scan(&m)
		if m > p {
			p = m
		}
		a[i] = m
		s += m
	i++
	}

	Scan(&v)

	p--
	for p+1 < s {
		m = (s + p) / 2
		c := 1
		d := 0
		for
		j := 0
		j < n
		{
			i = a[j]
			d += i
			if d > m {
				c++
				d = i
			}
		j++
		}

		if c > v {
			p = m
		} else {
			s = m
		}
	}

	Print(s)
}