package main
import . "fmt"
func main() {
	var (
		s = "NO"
		q = "YES"
		a, b, c, p, w, f int
		x = 1
		y = 1
	)
	Scanf("%c%d %c%d %c%d", &a, &w, &b, &p, &c, &f)
	
	if (c == b && (a != b || (w-p)*(w-f) > 0)) ||
	   (f == p && (w != p || (a-b)*(a-c) > 0)) {
		s = q
	}

	c = b-c
	f = p-f

	if c > 0 { x = -x }
	if f > 0 { y = -y }

	for c * f != 0 {
		c += x
		f += y
		b += x
		p += y

		if a == b && w == p { c = 9 }
	}
	if c+f == c*f { s = q }

	Print(s)
}