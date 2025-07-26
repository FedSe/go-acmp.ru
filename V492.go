package main
import . "fmt"
func main() {
	var x, y, a, b, v, t, d int
	s := "NO"

	Scan(&x, &y, &a, &b, &v, &t, &d)

	x += a * t
	y += b * t
	a = x*x + y*y
	v *= t
	t = d-v
	v += d
	if d*d > a {
		if t < 0 || a >= t*t {
			s = "YES"
		}
	} else if a <= v*v {
		s = "YES"
	}
	Print(s)
}