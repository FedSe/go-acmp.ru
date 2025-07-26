package main
import . "fmt"
func main() {
	var a, b, c, d, t int
	s := "YES"

	Scan(&t)
	for 0 < t {
		Scan(&a, &b, &c, &d)
		g:="NO"

		if a == c && b == d {
			g=s
		}
		for b > 0 {
			if b > a {
				a, b = b, a
				if a == c && b == d {
					g=s
				}
			}
			if b > 0 {
				O := a-c
				a -= a / b * b
				if b == d && a <= c && 0 <= O && O%b == 0 {
					g=s
				}
			}
		}

	Println(g)
	t--
	}

}