package main
import . "fmt"
func main() {
	var (
		w, h, p, d    [900]int
		n, t, l, i, c int
		q             = `
`
	)
	d[0] = 1

	Scan(&n)
	for l < n {
		l++
		Scan(&w[l])
		p[l]--
		t += w[l]
	}

	if t%3 < 1 {
		t /= 3
		for i < n {
			l = t
			i++
			for l >= w[i] {
				if d[l-w[i]] > d[l] {
					d[l] = 1
					p[l] = l - w[i]
					h[l] = i
				}
				l--
			}
		}
		for t > 0 {
			q += Sprintln(h[t])
			t = p[t]
			c++
		}
	}

	Print(c, q)
}