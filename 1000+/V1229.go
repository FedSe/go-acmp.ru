package main

import . "fmt"

type P struct {
	x, y int
}

func F(a, b P) int {
	d := a.x - b.x
	t := a.y - b.y
	return d*d + t*t
}
func main() {
	var (
		p [3]P
		b = 0
		s = "No"
	)

	for b < 3 {
		Scan(&p[b].x, &p[b].y)
		b++
	}

	b = F(p[1], p[2])
	c := F(p[0], p[2])

	switch F(p[0], p[1]) {
	case c - b, b - c, b + c:
		s = "Yes"
	}

	Print(s)
}