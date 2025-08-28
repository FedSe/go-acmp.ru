package main
import . "fmt"
var p, o [21][21][21]int
func F(a, b, c int) int {
	if a <= 0 || b <= 0 || c <= 0 {
		return 1
	}

	if a > 20 || b > 20 || c > 20 {
		return F(20, 20, 20)
	}

	r := p[a][b][c]
	if o[a][b][c] < 1 {
		if a < b && b < c {
			r = F(a, b, c-1) + F(a, b-1, c-1) - F(a, b-1, c)
		} else {
			a--
			r = F(a, b, c) + F(a, b-1, c) + F(a, b, c-1) - F(a, b-1, c-1)
			a++
		}

		p[a][b][c] = r
		o[a][b][c] = 1
	}
	return r
}

func main() {
	var a, b, c int
	Scan(&a, &b, &c)
	Print(F(a, b, c))
}