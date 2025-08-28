package main
import . "fmt"
func main() {
	a := ""
	b := a
	Scan(&a, &b)
	n := len(a)
	m := len(b)
	e := n + m
	d := -m + 1
	for d < n {
		l := d
		if l > 0 {
			l = 0
		}
		h := d + m
		if h < n {
			h = n
		}
		if h-l < e {
			c := l
			for c < h {
				s := 0
				if c >= 0 && c < n {
					s += int(a[c] - 48)
				}
				j := c - d
				if j >= 0 && j < m {
					s += int(b[j] - 48)
				}
				if s > 3 {
					goto A
				}
				c++
			}
			e = h - l
		A:
		}
		d++
	}
	Print(e)
}