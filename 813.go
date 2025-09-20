package main
import . "fmt"
type I int
var (
	a, b, c, d I
)

func F(w I) map[I]I {
	var (
		x, t, i I
		m       = []I{a, b, c, d}
		s       = map[I]I{}
	)
	for i < 4 {
		if 1<<i&w > 0 {
			t++
			x = i
		}
		i++
	}
	if t == 1 {
		s[m[x]] = 1
		return s
	}
	g := w
	for g > 0 {
		g--
		g &= w
		if g > 0 {
			for x := range F(g) {
				for y := range F(w ^ g) {
					s[x+y] = 1
					s[x-y] = 1
					s[x*y] = 1
				}
			}
		}
	}

	return s
}

func main() {
	Scan(&a, &b, &c, &d)
	Print([]any{"NO", "YES"}[F(15)[24]])
}