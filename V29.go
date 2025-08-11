package main
import . "fmt"
func main() {
	var (
		h, m [3e4]int
		n, i int
	)

	Scan(&n)
	for i < n {
		Scan(&h[i])
		i++
	}

	m[1] = h[1] - h[0]
	if m[1] < 0 {
		m[1] = -m[1]
	}
	i = 2
	for i < n {
		e := h[i] - h[i-1]
		if e < 0 {
			e = -e
		}
		m[i] = e + m[i-1]

		e = h[i] - h[i-2]
		if e < 0 {
			e = -e
		}

		e = e*3 + m[i-2]
		if e < m[i] {
			m[i] = e
		}

		i++
	}

	Print(m[n-1])
}