package main
import . "fmt"
func main() {
	var a, b, c, d, e, f int

	Scanf(`%2d %2d
%2d %2d %4d`, &a, &b, &c, &d, &e)
	for c != a || d != b {
		h := 28
		if e%4 < 1 && e%100 > 0 || e%400 < 1 {
			h = 29
		}
		if d != 2 {
			h = []int{0, 31, 0, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}[d]
		}
		c++
		if c == 32 && d == 12 {
			c = 1
			d = 1
			e++
		} else if c == h+1 {
			c = 1
			d++
		}
		f++
	}

	Print(f)
}