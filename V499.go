package main
import . "fmt"
var (a, z, x, b, y, u, k, w int
s = "NO"
)
func p(n, m int) {
	if m >= k && n <= w {
		s = "YES"
	}
}

func main() {
	Scan(&k, &w, &a, &b, &z, &y, &x, &u)

	p(a, b)
	p(z, y)
	p(x, u)
	p(a+z, b+y)
	p(a+x, b+u)
	p(x+z, u+y)
	p(a+z+x, b+y+u) 

	Print(s)
}