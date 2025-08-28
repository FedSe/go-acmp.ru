package main
import . "fmt"
var (
	k = 0
	a = ""
	b = a
	P = Print
)
func F(n, x, y int) {
	n++
	if n < 3 {
		i := x - int(b[0])
		l := y - int(b[1])
		if i < 0 {
			i = -i
		}
		if l < 0 {
			l = -l
		}
		if (i == 1 && l == 2) || (i == 2 && l == 1) {
			k = n
			return
		}

		for _, v := range [][2]int{
			{1, 2}, {1, -2}, {-1, 2}, {-1, -2},
			{2, 1}, {2, -1}, {-2, 1}, {-2, -1}} {
			q := x + v[0]
			g := y + v[1]
			if q > 96 && q < 105 && g > 48 && g < 57 {
				F(n, q, g)
			}
		}
	}
}

func main() {
	Scan(&a, &b)

	F(0, int(a[0]), int(a[1]))

	if k < 1 {
		P("NO")
	} else {
		P(k)
	}
}