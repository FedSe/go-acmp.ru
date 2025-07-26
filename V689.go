package main
import . "fmt"
func main() {
	var n, m, e int
	Scan(&n)
	for 0 < n {
		Scan(&m)
		t := 999999999
		for
		b := 2
		b <= 36
		{
			var (
				d [36]bool
				x = m
				c = 0
			)

			for x > 0 {
				d[x%b] = 1>0
				x /= b
				c++
			}

			for _, d := range d {
				if d {
					c++
				}
			}

			if c < t {
				t = c
				e = b
			}
		b++
		}

		s := ""
		for m > 0 {
			r := m % e
			if r > 9 {
				r += 7
			}
			s += string(r + 48)
		m /= e
		}

		r := []rune(s)
		for
		i, j := 0, len(r)
		i < j
		{
		j--
			r[i], r[j] = r[j], r[i]
		i++
		}

		Println(e, string(r))
	n--
	}
}