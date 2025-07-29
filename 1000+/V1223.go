package main
import (
	. "fmt"
	. "math"
)

type P struct {
	x, y float64
}

func F(p1 P, p2 P) float64 {
	return Sqrt((p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y))
}

func main() {
	var (
		a             [1e3]P
		n, p, r, i, j int
		q             = 1
		s             = 1
	)
	Scan(&n)
	for j < n {
		Scan(&a[j].x, &a[j].y)
		j++
	}

	z := F(a[0], a[1])
	x := z

	for i < n-1 {
		j = i + 1
		for j < n {
			c := F(a[i], a[j])
			if c < z {
				z = c
				p = i
				q = j
			}
			if c > x {
				x = c
				r = i
				s = j
			}
			j++
		}
		i++
	}

	Print(p+1, q+1, r+1, s+1)
}