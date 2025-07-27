package main
import . "fmt"
func main() {
	var (
		p, e                   [10001]int
		n, a, s, g, i, x, y, j int
	)

	Scan(&n)
	for i < n {
		i++
		Scan(&a)
		if a > 0 {
			p[x] = i
			s += a
			x++
		}
		if a < 0 {
			e[y] = i
			g -= a
			y++
		}
	}

	if s >= g {
		e = p
		y = x
	}

	Println(y)
	for j < y {
		Println(e[j])
		j++
	}
}