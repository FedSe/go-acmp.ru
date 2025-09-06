package main
import . "fmt"
func main() {
	var (
		t                      [1440]int
		n, c, m, o, p, i, a, z int
	)

	Scan(&n)
	for i < n {
		Scan(&c, &m, &o, &p)
		m += 60 * c
		p += 60 * o
		x := 0
		y := 1440
		j := 0
		if m < p {
			x = m
			y = p
		} else if p < m {
			x = m
			j--
		}
		for j < 1 {
			for x < y {
				t[x]++
				x++
			}
			x = 0
			y = p
			j++
		}
		i++
	}

	for z < 1440 {
		if t[z] == n {
			a++
		}
		z++
	}

	Print(a)
}