package main
import . "fmt"
func main() {
	var n, m, y, x, i, j, d, c, t, l int

	Scan(&n, &m, &y, &x)
A:
	for i != y-1 || j != x-1 {
		q := i + d%2 - d/3*2
		w := j + 1 - d + d/3*2
		if q >= t && q < n && w >= l && w < m {
			i = q
			j = w
			c++
			goto A
		}
		t += 1 >> d
		m -= d >> 1 & d
		n -= d / 2 & d
		l += d / 3

		d++
		d %= 4
	}

	Print(c + 1)
}