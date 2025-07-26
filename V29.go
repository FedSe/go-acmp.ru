package main
import . "fmt"
func main() {
	var (
		h, m [30000]int
		n, i, a int
	)
	Scan(&n)

	for i < n {
		Scan(&h[i])
	i++
	}

	if n > 1 {
		m[1] = h[1] - h[0]
		if m[1] < 0 {
			m[1] = -m[1]
		}
		for
		i = 2
		i < n
		{
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
		a = m[n-1]
	}

	Print(a)
}