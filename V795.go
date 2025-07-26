package main
import . "fmt"
func main() {
	var (
		s [100]int
		n, x, y, t, p, i int
		a = -1
	)

	Scan(&n)
	for i < n {
		Scan(&s[i])
	i++
	}
	Scan(&y, &t, &p)

	t--
	y--
	y /= p
	t = t/p-y+1
	if t > n { t = n }

	for a < 2 {
		p = a * y
		p %= n
		p += n
		p %= n
		i = 0
		for i < t {
			if x < s[p] { x = s[p] }
			p += a
			p %= n
			p += n
			p %= n
		i++
		}
	a += 2
	}

	Print(x)
}