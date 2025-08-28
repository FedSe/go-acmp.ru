package main
import . "fmt"
func main() {
	var n, m, e int

	Scan(&n)
	for 0 < n {
		Scan(&m)
		t := 1 << 30
		i := 2
		for i < 37 {
			var (
				d [36]int
				x = m
				c = 0
			)

			for x > 0 {
				d[x%i] = 1
				x /= i
				c++
			}

			for _, d := range d {
				c += d
			}

			if c < t {
				t = c
				e = i
			}
			i++
		}

		s := ""
		for m > 0 {
			t = m % e
			if t > 9 {
				t += 7
			}
			s += string(t + 48)
			m /= e
		}

		r := []rune(s)
		i = 0
		t = len(r)
		for i < t {
			t--
			r[i], r[t] = r[t], r[i]
			i++
		}

		Println(e, string(r))
		n--
	}
}