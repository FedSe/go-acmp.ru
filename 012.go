package main
import . "fmt"

type P struct {
	x, y float64
}

func F(a, b, c P) float64 {
	A := (b.x-a.x)*(c.y-a.y) - (c.x-a.x)*(b.y-a.y)
	if A < 0 {
		A = -A
	}
	return A / 2
}
func main() {
	var (
		a, b, c, d, h P
		n, k          int
	)
	Scan(&n)
	for 0 < n {
		Scan(&h.x, &h.y, &a.x, &a.y, &b.x, &b.y, &c.x, &c.y, &d.x, &d.y)
		if F(a, b, h)+F(b, c, h)+F(c, d, h)+F(d, a, h)-F(a, b, c)-F(b, c, d) < 1e-7 {
			k++
		}
		n--
	}
	Print(k)
}