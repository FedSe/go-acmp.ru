package main
import . "fmt"
func main() {
	var a, b, z, y, x, u, k, w, i int
	s := "NO"

	Scan(&k, &w, &a, &b, &z, &y, &x, &u)
	for i < 8 {
		i++
		n := 0
		m := 0
		if i&1 > 0 {
			n += a
			m += b
		}
		if i&2 > 0 {
			n += z
			m += y
		}
		if i&4 > 0 {
			n += x
			m += u
		}
		if m >= k && n <= w {
			s = "YES"
		}
	}

	Print(s)
}