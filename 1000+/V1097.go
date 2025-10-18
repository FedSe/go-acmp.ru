package main
import . "fmt"
func main() {
	var n, m, i, q int
	P := Print

	Scan(&n, &m)
	if n > m {
		n, m = m, n
		q = 1
	}

	P(m)
	for i < m {
		P(`
`, n, " ")
		j := 0
		for j < n {
			l := (j+i)%m + 1
			j++
			h := j
			if q < 1 {
				l, h = h, l
			}
			P(l, "-", h, " ")
		}
		i++
	}
}